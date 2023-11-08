package routes

import (
	"fmt"
	"log/slog"

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
		headers := c.GetReqHeaders()
		if _, ok := headers["Authorization"]; !ok {
			c.Status(401)
			slog.Error("authorization header not presented")
			return c.JSON(fiber.Map{
				"status":  401,
				"message": "authorization header not presented",
			})
		}
		if len(headers["Authorization"]) == 0 {
			c.Status(401)
			slog.Error("authorization header not presented")
			return c.JSON(fiber.Map{
				"status":  401,
				"message": "authorization header not presented",
			})
		}
		fmt.Println(headers["Authorization"][0])
		tmp := headers["Authorization"][0][9:]
		c.Status(200)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data": map[string]string{
				"sub":  tmp,
				"name": "test_dummy",
			},
		})
	})

}
