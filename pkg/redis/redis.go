package redis

import (
	"context"
	"log"
	"social-media-app/pkg/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Connect(conf config.Redis) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client = rdb
	log.Println("redis connect successfully ...")
}

func Client() *redis.Client {
	return client
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Delete(ctx context.Context, key string) error {
	if err := client.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

func Exists(ctx context.Context, key string) (bool, error) {
	val, err := client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

func Ping(ctx context.Context) error {
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return nil
}
