package main

import (
	"fmt"
	"time"

	"learingagain/internal/report"
	"learingagain/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BuildServer() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024 * 100000,
	})
	fmt.Println("creating server...")

	app.Use(cors.New())
	app.Use(logger.New())

	StartMongo("mongodb://localhost:27017", "ejemplo", 10*time.Second)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	user.UserGroup = app.Group("/user") //Create user routes groups
	user.UserRouter()                   //Initialize all the routes for user

	report.ReportGroup = app.Group("/report")
	report.ReportRouter()

	return app
}
