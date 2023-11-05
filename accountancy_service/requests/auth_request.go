package requests

import (
	"program_akuntansi/accountancy_service/controllers"
	services "program_akuntansi/accountancy_service/service"
	"program_akuntansi/utilities"
	"strings"

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

func LoginUser(c *fiber.Ctx) error { //GET
	headers := c.GetReqHeaders()

	user, err := services.AuthUser(headers["Authorization"][0])
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"content": user,
	})
}

func AuthUser(c *fiber.Ctx, auth_role_env string) error {
	headers := c.GetReqHeaders()

	user, err := services.AuthUser(headers["Authorization"][0])
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	auth_roles := utilities.GoDotEnvVariable(auth_role_env)
	arr_roles := strings.Split(auth_roles, ",")

	_, found := utilities.Find(arr_roles, func(x string) bool {
		return x == user.Role
	})

	if !found && user.Role != "admin" {
		c.Status(403)
		return c.JSON(fiber.Map{
			"status":  403,
			"message": "Forbidden",
		})
	}
	return nil
}
