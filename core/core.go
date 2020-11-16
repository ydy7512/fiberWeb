package core

import (
	"fiberWeb/config"
	"fiberWeb/router"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"strconv"
)

func NewEngine() {
	// 初始化数据库
	poolSwitch, _ := strconv.Atoi(config.Env("DB_POOL_SWITCH"))
	maxIdle, _ := strconv.Atoi(config.Env("DB_POOL_SWITCH"))
	maxOpen, _ := strconv.Atoi(config.Env("DB_POOL_SWITCH"))
	defaultMysqlConfig := Config{
		Addr:       config.Env("DB_ADDR"),
		Database:   config.Env("DB_DATABASE"),
		User:       config.Env("DB_USER"),
		Password:   config.Env("DB_PASSWORD"),
		PoolSwitch: poolSwitch,
		MaxIdle:    maxIdle,
		MaxOpen:    maxOpen,
	}
	initializeDB("default", defaultMysqlConfig)
	// 初始化redis
	db, _ := strconv.Atoi(config.Env("REDIS_DB"))
	poolSize, _ := strconv.Atoi(config.Env("REDIS_POOL_SIZE"))
	initializeRedis("default", &redis.Options{
		Addr:     config.Env("REDIS_ADDR"),
		Password: config.Env("REDIS_PASSWORD"),
		DB:       db,
		PoolSize: poolSize,
	})

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + config.Env("SERVER_PORT")))
}
