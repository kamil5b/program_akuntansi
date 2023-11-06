package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
		})
	})

	api := app.Group("/api")

	// ==Auth==
	auth := api.Group("/auth")
	// auth.Post("/register", requests.RegisterUserAuth)
	// auth.Get("/user", requests.LoginUser)

	//==DUMMY==
	auth.Get("/dummy", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": map[string]string{
				"sub":  "28361023832",
				"name": "test_dummy",
			},
		})
	})

}
