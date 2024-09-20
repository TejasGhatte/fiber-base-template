package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"fiber-app/initializers"
	"fiber-app/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.DBMigrate()
	initializers.AddLogger()
	initializers.ConnectToCache()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routers.Config(app)

	app.Listen(":" + initializers.CONFIG.PORT)
}