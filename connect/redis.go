package connect

import "github.com/go-redis/redis"

func NewRedisClient(address string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})
}
