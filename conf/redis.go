package conf

import (
	"github.com/go-redis/redis/v8"
)

func RedisConn() *redis.Client {
	address := GetRedisConnectionURL()
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   0,
	})

	return rdb
}
