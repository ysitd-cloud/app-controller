package core

import "github.com/go-redis/redis"

type Publisher struct {
	Redis *redis.Client
}

func NewPublisher(client *redis.Client) *Publisher {
	return &Publisher{
		Redis: client,
	}
}
