package requests

import (
	"program_akuntansi/controllers"

	"github.com/gofiber/fiber/v2"
)

// Function which processing Registeration outputing referral link
func RegisterUserAuth(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		Authorization Header
		name : string
		role : string
	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	headers := c.GetReqHeaders()

	if err := controllers.RegisterAuthUser(headers["Authorization"][0], data["name"], data["role"], data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  201,
		"message": "success",
	})
}

func LoginUser(c *fiber.Ctx) error { //POST
	var auth_id uint
	/*
		auth_id : uint
	*/
	if err := c.BodyParser(&auth_id); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := controllers.GetUserByAccID(auth_id)
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Username and Password Invalid",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"user":    user,
	})
}
