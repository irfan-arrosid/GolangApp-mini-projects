package main

import (
	"GolangApp-mini-projects/go-redis-client/cache"
	"context"
	"log"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0
)

func main() {
	ctx := context.Background()

	c := cache.NewCache(RedisAddr, RedisPassword, RedisDB)
	if err := c.Ping(ctx); err != nil {
		log.Panic("Failed to connect Redis")
	}

	log.Println("Redis connected....")

	if err := c.Set(ctx, "user:name", "Atira", 0); err != nil {
		log.Println("Error: could not store a value to Redis")
	}

	log.Println("Value stored to Redis")

	res, err := c.Get(ctx, "user:name")
	if err != nil {
		log.Println("Error: could not get a value from Redis")
	}

	log.Println("Result:", res)
}
