package core

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redis map[string]*redis.Client
var ctx = context.Background()

func initializeRedis(name string, config *redis.Options) {
	Redis[name] = redis.NewClient(config)
	rds := Redis[name]
	_, err := rds.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败...")
	} else {
		fmt.Println("redis连接成功...")
	}
}
