package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func NewCache(addr string, password string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &Cache{client}
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}
