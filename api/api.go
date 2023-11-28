package api

import (
	"github.com/gofiber/fiber/v2"
)

func handelWelcome(c *fiber.Ctx) error {
	return c.Render("/workspaces/Web/html/welcome.html", fiber.Map{}, "html")
}

// Router handles all routes and listens tls
func Router() {
	app := fiber.New()
	app.Get("/", handelWelcome)
	app.Static("/css/style.css", "./css/style.css", fiber.Static{})
	app.Listen(":8080")
}
