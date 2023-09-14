package redis

import "github.com/go-redis/redis"

func ConnectRedis() {
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
