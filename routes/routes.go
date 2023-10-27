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

	// //----AUTH----
	// app.Post("/api/register/:ref_code?", controllers.RegisterUser)
	// app.Get("/api/register/:ref_code", controllers.CheckReferralCode)
	// app.Post("/api/login", controllers.LoginUser)
	// app.Get("/api/:auth", controllers.AuthUser)

	// //----REFERRAL----

	// app.Get("/api/ref/new/:auth", controllers.NewReferralCode)
	// app.Get("/api/ref/:auth", controllers.GetReferralCode)
}
