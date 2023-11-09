package requests

import (
	"errors"
	"program_akuntansi/accountancy_service/controllers"
	"program_akuntansi/accountancy_service/models"
	"program_akuntansi/accountancy_service/services"
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
			"message": err.Error(),
		})
	}

	acc_id, err := GetAccountID(c)
	if err != nil {
		c.Status(401)

		return c.JSON(fiber.Map{
			"status":  401,
			"message": err.Error(),
		})
	}
	if err := controllers.RegisterAuthUser(acc_id, data["name"], data["role"]); err != nil {
		c.Status(401)

		return c.JSON(fiber.Map{
			"status":  401,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  201,
		"message": "success",
	})
}

func LoginUser(c *fiber.Ctx) error { //GET
	user, err := GetUserByAuth(c)

	if err != nil {
		if err.Error() == "record not found" {
			c.Status(401)

			return c.JSON(fiber.Map{
				"status":  401,
				"message": "the account haven't registered yet",
			})
		}
		c.Status(401)

		return c.JSON(fiber.Map{
			"status":  401,
			"message": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    user,
	})
}

func GetAccountID(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()
	if _, ok := headers["Authorization"]; !ok {
		return 0, errors.New("authorization header not presented")

	}
	if len(headers["Authorization"]) == 0 {
		return 0, errors.New("authorization header is null")
	}
	return services.AuthUser(headers["Authorization"][0])
}

func GetUserByAuth(c *fiber.Ctx) (models.User, error) {
	acc_id, err := GetAccountID(c)
	if err != nil {
		return models.User{}, err
	}
	return controllers.GetUserByAccID(uint(acc_id))
}

func AuthUser(c *fiber.Ctx, auth_role_env string) error {
	user, err := GetUserByAuth(c)
	if err != nil {
		return err
	}
	auth_roles := utilities.GoDotEnvVariable(auth_role_env)
	arr_roles := strings.Split(auth_roles, ",")

	_, found := utilities.Find(arr_roles, func(x string) bool {
		return x == user.Role
	})

	if !found && user.Role != "admin" {
		return errors.New("forbidden")
	}
	return nil
}
