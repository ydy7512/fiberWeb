package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redis = make(map[string]*redis.Client)
var Ctx = context.TODO()

// InitializeRedis 初始化 redis 配置
func InitializeRedis(name string, config *redis.Options) {
	fmt.Println("redis初始化...")
	Redis[name] = redis.NewClient(config)
	rds := Redis[name]
	_, err := rds.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败...")
	} else {
		fmt.Println("redis连接成功...")
	}
}
