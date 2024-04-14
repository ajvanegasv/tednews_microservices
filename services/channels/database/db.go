package database

import (
	"github.com/redis/go-redis/v9"
	"github.com/ajvanegasv/tednews_microservices/services/tednews/config"
	"strings"
)

var redisDb *redis.Client

func Init() {
	conf := config.GetConfig()
	redisDb = redis.NewClient(&redis.Options{
		Addr: strings.Join([]string{conf.GetString("redis.host"), conf.GetString("redis.port")}, ":"),
		Password: conf.GetString("redis.password"),
		DB: conf.GetInt("redis.db"),
	})
}

func GetRedisDb() *redis.Client {
	return redisDb
}