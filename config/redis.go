package config

import (
	"github.com/go-redis/redis"
	"log"

	"exchangeapp/global"
)

func initRedis() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Fail to connect redis ,got error: %v", err)
	}
	global.RedisDb = RedisClient
}
