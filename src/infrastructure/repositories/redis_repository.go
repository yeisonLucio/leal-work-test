package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const expiration = time.Minute * 5

var ctx = context.Background()

type RedisRepository struct {
	RedisClient *redis.Client
	Logger      *logrus.Entry
}

func (r *RedisRepository) SetByKey(key string, object string) error {
	log := r.Logger.WithFields(logrus.Fields{
		"file":   "redis_repository",
		"method": "SetByKey",
		"key":    key,
		"object": object,
	})

	if err := r.RedisClient.Set(ctx, key, object, expiration); err.Err() != nil {
		log.WithError(err.Err()).Error("error set key")
		return err.Err()
	}

	return nil
}

func (r *RedisRepository) GetByKey(key string) (string, error) {
	log := r.Logger.WithFields(logrus.Fields{
		"file":   "redis_repository",
		"method": "GetByKey",
		"key":    key,
	})

	result, err := r.RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.WithError(err).Error("error get key")
		return "", err
	}

	return result, nil
}
