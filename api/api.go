package api

import (
	"github.com/gofiber/fiber/v2"
)

// Sets content type for CSS file
func handleCSS(c *fiber.Ctx) error {
	c.Set("Content-type", "text/css")
	return nil
}

// Returns the html home page
func handelWelcome(c *fiber.Ctx) error {
	return c.Render("/workspaces/Web/html/welcome.html", fiber.Map{}, "html")
}

// Handles button authentication
func handleAuth(c *fiber.Ctx) error {
	return nil
}

// Router handles all routes and listens tls
func Router() {
	app := fiber.New()

	// Serves home page
	app.Get("/", handelWelcome)

	// Handles button submission
	app.Post("/", handleAuth)

	//Static CSS file
	app.Static("/css/style.css", "./css/style.css", fiber.Static{ModifyResponse: handleCSS})
	app.Listen(":8080")
}
