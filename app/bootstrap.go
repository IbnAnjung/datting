package app

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/datting/driver"
	"github.com/IbnAnjung/datting/repository/mysqlgorm/user_repository"
	authUC "github.com/IbnAnjung/datting/usecase/auth"
)

func Start(ctx context.Context) (func(), error) {

	conf, err := LoadConfig()
	if err != nil {
		return func() {}, fmt.Errorf("load config failed: %s", err.Error())
	}

	mysqlConf := conf.Mysql

	sqlConfig := driver.LoadMySqlConfig(
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Schema,
		mysqlConf.MaxLifeConnection,
		mysqlConf.MaxConnection,
		mysqlConf.MaxIddleConnection,
	)

	dbconn, mysqlCleanup, err := driver.NewMysqlConnection(ctx, sqlConfig)
	if err != nil {
		return func() {}, err
	}

	redisConf := conf.Redis
	redisConfig := driver.LoadRedisConfig(redisConf.Host, redisConf.Port, redisConf.User, redisConf.Password, redisConf.Db)
	_, redisCleanup, err := driver.NewRedisConnection(ctx, redisConfig)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	orm, err := driver.NewGormOrm("mysql", dbconn)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}
	// repository
	userRepo := user_repository.New(orm)

	// validator
	validator, err := driver.NewValidator()
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	crypt := driver.NewBycrypt()
	jwt := driver.NewJwt(conf.App.Name, conf.Jwt.SecretKey, conf.Jwt.ExpireDuration)

	authUC := authUC.New(
		userRepo,
		validator,
		crypt,
	)

	router := LoadGinRouter(
		authUC, jwt,
	)

	httpCleanup, err := driver.RunGinHttpServer(ctx, router, driver.LoadHttpConfig(conf.Http.Port))
	if err != nil {
		return func() {
			redisCleanup()
			mysqlCleanup()
		}, err
	}

	return func() {
		mysqlCleanup()
		redisCleanup()
		httpCleanup()
	}, nil
}
