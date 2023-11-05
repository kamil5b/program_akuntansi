package requests

import (
	"errors"
	"program_akuntansi/auth_service/controllers"

	"github.com/gofiber/fiber/v2"
)

//GET

func GetTransactionByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header
		id : uint

	*/

	if err := AuthUser(c, "AUTH_GET_TRANSACTION_ID"); err != nil {
		c.Status(403)
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err,
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}
	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": errors.New("id not valid"),
		})
	}

	transaction, err := controllers.GetTransactionByID(uint(id))
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    transaction,
	})
}

func GetTransactionByInvoiceID(c *fiber.Ctx) error { //GET
	var data struct {
		InvoiceID   uint   `json:"invoice_id"`
		InvoiceType string `json:"invoice_type"`
	}
	/*
		Authorization Header
		invoice_id : 	uint
		invoice_type :  string
	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}
	if err := AuthUser(c, "AUTH_GET_TRANSACTION_INVOICE"); err != nil {
		c.Status(403)
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err,
		})
	}

	transactions, err := controllers.GetTransactionByInvoiceID(data.InvoiceID, data.InvoiceType)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    transactions,
	})
}

func GetAllTransaction(c *fiber.Ctx) error { //GET
	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_TRANSACTION"); err != nil {
		c.Status(403)
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err,
		})
	}

	transactions, err := controllers.GetAllTransactions()
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    transactions,
	})
}
