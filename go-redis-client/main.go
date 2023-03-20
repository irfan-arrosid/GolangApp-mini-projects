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
}

// Go to video on 7:30
