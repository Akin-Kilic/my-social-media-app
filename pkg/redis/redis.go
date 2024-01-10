package redis

import (
	"context"
	"log"
	"net"
	"social-media-app/pkg/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Connect(conf config.Redis) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	client = rdb
	log.Println("redis connect successfully ...")
}

func Client() *redis.Client {
	return client
}

func Set(key string, value string, expiration time.Duration) error {
	if err := client.Set(context.Background(), key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Delete(key string) error {
	if err := client.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}

func Exists(key string) (bool, error) {
	val, err := client.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}
