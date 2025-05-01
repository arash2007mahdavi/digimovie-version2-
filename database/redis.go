package database

import (
	"context"
	"digimovie/src/config"
	"digimovie/src/logging"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
		DialTimeout: 5 * time.Second,
		DB: 0,
		PoolSize: 5,
		PoolTimeout: 100 * time.Millisecond,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Error(logging.Redis, logging.Startup, "failed to connect", nil)
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return redisClient
}

func CLoseRedis() {
	redisClient.Close()
}

func Get[T any](key string) (T, error) {
	var dest T
	result, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(result), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}

func Set[T any](key string, value T, duration time.Duration) error {
	json_value, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = redisClient.Set(context.Background(), key, json_value, duration*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}