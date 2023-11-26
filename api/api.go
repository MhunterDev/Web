package api

import (
	"github.com/gofiber/fiber/v2"
)

func handelWelcome(c *fiber.Ctx) error {
	return c.Render("/workspaces/Web/api/html/welcome.html", fiber.Map{}, "html")
}

// Router handles all routes and listens tls
func Router() {
	app := fiber.New()

	app.Static("/static", "./css")
	app.Get("/", handelWelcome)

	app.Listen(":8080")
}
