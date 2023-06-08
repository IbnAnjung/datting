package app

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/datting/driver"
	"github.com/IbnAnjung/datting/repository/mysqlgorm"
	"github.com/IbnAnjung/datting/repository/redis_repository"
	accountUC "github.com/IbnAnjung/datting/usecase/account"
	authUC "github.com/IbnAnjung/datting/usecase/auth"
	userUC "github.com/IbnAnjung/datting/usecase/user"
	userSwapUC "github.com/IbnAnjung/datting/usecase/user_swap"
	"github.com/IbnAnjung/datting/utils"
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
	redisConn, redisCleanup, err := driver.NewRedisConnection(ctx, redisConfig)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	cache := utils.NewRedisCaching(redisConn)

	orm, err := driver.NewGormOrm("mysql", dbconn)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}
	// repository
	userRepo := mysqlgorm.NewUserRepository(orm)
	userCacheRepo := redis_repository.NewUserCacheRepository(cache)
	userSwapRepo := mysqlgorm.NewUserSwapRepository(orm)

	// validator
	validator, err := utils.NewValidator()
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	crypt := utils.NewBycrypt()
	jwt := utils.NewJwt(conf.App.Name, conf.Jwt.SecretKey, conf.Jwt.ExpireDuration)

	// usecase
	authUC := authUC.New(userRepo, validator, crypt, jwt)
	userUC := userUC.New(validator, userRepo, userCacheRepo)
	userSwapUC := userSwapUC.New(validator, userRepo, userCacheRepo, userSwapRepo)
	accountUC := accountUC.New(userRepo)

	router := LoadGinRouter(
		authUC, userUC, userSwapUC, accountUC, jwt,
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
