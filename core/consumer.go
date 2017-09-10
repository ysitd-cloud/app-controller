package core

import "github.com/go-redis/redis"

type Consumer struct {
	Redis *redis.Client
}

func NewConsumer(client *redis.Client) *Consumer {
	return &Consumer{
		Redis: client,
	}
}
