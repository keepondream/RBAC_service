package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

type Cache struct {
	client *redis.Client
	prefix string
}

func NewCache() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       cast.ToInt(os.Getenv("REDIS_DB")),
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(`new cache failed `, err)
	}

	return &Cache{
		client: rdb,
		prefix: os.Getenv("REDIS_PREFIX"),
	}
}

func (c *Cache) GetKey(key string) string {
	return fmt.Sprintf("%s:%s", c.prefix, key)
}
