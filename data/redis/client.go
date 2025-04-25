package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis(addr string) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr, // e.g., "localhost:6379"
		Password: "",   // set if you use one
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Client.Ping(ctx).Result()
	return err
}