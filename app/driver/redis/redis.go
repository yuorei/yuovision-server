package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if redisDB == nil {
		log.Fatalln("redis connection failed")
	}
	if err := redisDB.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("redis ping failed: ", err)
	}

	return redisDB
}
