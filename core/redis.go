package core

import (
	"context"
	"fiberWeb/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var Redis *redis.Client
var ctx = context.Background()

func initializeRedis() {
	db, _ := strconv.Atoi(config.Config("REDIS_DB"))
	poolSize, _ := strconv.Atoi(config.Config("REDIS_POOL_SIZE"))
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Config("REDIS_ADDR"),
		Password: config.Config("REDIS_PASSWORD"),
		DB:       db,
		PoolSize: poolSize,
	})
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败...")
	} else {
		fmt.Println("redis连接成功...")
	}
}
