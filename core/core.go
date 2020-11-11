package core

import (
	"fiberWeb/config"
	"fiberWeb/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func NewEngine() {
	// 初始化数据库
	initializeDB()
	// 初始化redis
	initializeRedis()

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + config.Env("SERVER_PORT")))
}
