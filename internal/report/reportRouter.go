package report

import (
	"github.com/gofiber/fiber/v2"
)

var ReportGroup fiber.Router

func ReportRouter() {

	ReportGroup.Post("/create", CalculateReport)
	ReportGroup.Post("/health", func(c *fiber.Ctx) error {
		return c.SendString("Report Healthy!")
	})

}
