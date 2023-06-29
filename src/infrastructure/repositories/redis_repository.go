package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const expiration = time.Minute * 30

var ctx = context.Background()

type RedisRepository struct {
	RedisClient *redis.Client
}

func (r *RedisRepository) SetByKey(key string, object string) error {
	if err := r.RedisClient.Set(ctx, key, object, expiration); err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (r *RedisRepository) GetByKey(key string) (string, error) {
	result, err := r.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
