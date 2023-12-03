package api

import (
	db "github.com/MhunterDev/Web/db"

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

// Returns homepage html
func handleHome(c *fiber.Ctx) error {
	var cookie string
	c.CookieParser(&cookie)
	if len(cookie) != 0 {
		return c.Render("/workspaces/Web/html/home.html", fiber.Map{}, "html")
	} else {
		return c.Render("/workspaces/Web/html/welcome.html", fiber.Map{}, "html")
	}
}

// Handles button authentication
func handleAuth(c *fiber.Ctx) error {
	var u struct {
		Username string `json:"username" form:"username" xml:"username" binding:"required"`
		Password string `json:"password" form:"password" xml:"password" binding:"required"`
	}
	parseErr := c.BodyParser(&u)
	if parseErr != nil {
		return parseErr
	}

	err := db.AuthPass(u.Username, u.Password)

	if err != nil {
		return c.Render("/workspaces/Web/html/welcome.html", fiber.Map{}, "html")
	}
	cookie := new(fiber.Cookie)
	cookie.Name = u.Username
	c.Cookie(cookie)
	return c.Redirect("/home", 200)
}

// Router handles all routes and listens tls
func Router() {
	app := fiber.New()

	// Serves home page
	app.Get("/", handelWelcome)
	app.Put("/home", handleHome)
	app.Route("/home",func (fiber.)
	// Handles button submission
	app.Post("/", handleAuth)

	//Static CSS file
	app.Static("/css/style.css", "./css/style.css", fiber.Static{ModifyResponse: handleCSS})
	app.Listen(":8080")
}
