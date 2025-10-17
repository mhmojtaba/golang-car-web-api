package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:           cfg.Redis.Password,
		DB:                 0,
		DialTimeout:        cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:        cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout:       cfg.Redis.WriteTimeout * time.Second,
		PoolSize:           cfg.Redis.PoolSize,
		PoolTimeout:        cfg.Redis.PoolTimeout * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: cfg.Redis.IdlCheckFrequency * time.Millisecond,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

func Set[T any](c *redis.Client, key string, value T, expiration time.Duration) error {
	v, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return c.Set(key, v, expiration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var result T = *new(T)
	val, err := c.Get(key).Result()
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return result, err
	}
	return result, err
}
