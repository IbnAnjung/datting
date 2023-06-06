package app

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/datting/driver"
)

func Start(ctx context.Context) (func(), error) {

	conf, err := LoadConfig()
	if err != nil {
		return func() {}, fmt.Errorf("Load config failed: %s", err.Error())
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

	_, mysqlCleanup, err := driver.NewMysqlConnection(ctx, sqlConfig)
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

	router := LoadGinRouter()
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
