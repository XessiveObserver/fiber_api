package main

import (
	"github.com/XessiveObserver/fiber_api/database"
	"github.com/XessiveObserver/fiber_api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetUpRoutes(app)

	// handle unavailable routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 Not Found
	})

	app.Listen(":8000")
}
