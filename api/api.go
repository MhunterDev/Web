package api

import (
	"github.com/gofiber/fiber/v2"
)

func handelWelcome(c *fiber.Ctx) error {
	return c.Render("/workspaces/hunterdev/Web/api/html/welcome.html", fiber.Map{}, "html")
}

// Router handles all routes and listens tls
func Router() {
	app := fiber.New()

	app.ListenTLS(":8080", "/etc/mhdev/keychain/tls/ca.crt", "/etc/mhdev/keychain/tls/secret/ca.crt")
}
