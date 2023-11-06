package routes

import (
	"math/rand"
	"strconv"

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

	//==DUMMY== TEMPORARY
	auth.Get("/user", func(c *fiber.Ctx) error {
		tmp := rand.Intn(1000000) + 1
		c.Status(200)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": map[string]string{
				"sub":  strconv.Itoa(tmp),
				"name": "test_dummy",
			},
		})
	})

}
