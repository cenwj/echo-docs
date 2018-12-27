package db

import (
	"github.com/cenwj/echo-docs/conf"
	"github.com/go-redis/redis"
)

var Rclient *redis.Client

func RedisConn(index int) *redis.Client {
	RClient := redis.NewClient(&redis.Options{
		Addr:     conf.Config().Redis.RdbHost + ":" + conf.Config().Redis.RdbPort,
		Password: conf.Config().Redis.RdbPass,
		DB:       index,
	})

	if _, err = RClient.Ping().Result(); err != nil {
		panic(err)
	}

	return RClient
}
