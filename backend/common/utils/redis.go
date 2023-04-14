package utils

import (
	"ai-backend/common/global"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisUtil struct {
}

var Redis = new(redisUtil)

func (r *redisUtil) Set(key string, value interface{}, expiration time.Duration) error {
	return global.GRedis.Set(global.GRedis.Context(), key, value, expiration).Err()
}

func (r *redisUtil) GetObject(key string, value interface{}) error {
	return global.GRedis.Get(global.GRedis.Context(), key).Scan(value)
}

func (r *redisUtil) GetString(key string) (string, error) {
	get := global.GRedis.Get(global.GRedis.Context(), key)
	err := get.Err()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return get.Val(), nil
}

func (r *redisUtil) GetInt(key string) (int, error) {
	i, err := global.GRedis.Get(global.GRedis.Context(), key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return i, err
}
