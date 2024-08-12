package services

import (
	"context"

	"github.com/edaywalid/chat-app/configs"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService(config *configs.Config) *RedisService {
	opt, err := redis.ParseURL(config.RedisUrl)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	return &RedisService{client: client}
}

func (s *RedisService) Publish(channel string, message []byte) error {
	return s.client.Publish(context.Background(), channel, message).Err()
}

func (s *RedisService) Subscribe(channel string) *redis.PubSub {
	return s.client.Subscribe(context.Background(), channel)
}
