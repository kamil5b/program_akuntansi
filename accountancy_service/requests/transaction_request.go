package requests

import (
	"log/slog"
	"program_akuntansi/accountancy_service/controllers"
	"strings"

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
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)
		slog.Error("id not valid")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	transaction, err := controllers.GetTransactionByID(uint(id))
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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

	/*
		Authorization Header
		invoice_id : 	uint
		invoice_type :  string
	*/

	if err := AuthUser(c, "AUTH_GET_TRANSACTION_INVOICE"); err != nil {
		c.Status(403)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	inv_type := strings.ToUpper(c.Params("inv_type"))
	if inv_type != "DEBIT" && inv_type != "CREDIT" {
		c.Status(400)
		slog.Error("invalid invoice type")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "invalid invoice type",
		})
	}
	if id == 0 {
		c.Status(400)
		slog.Error("id not valid")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}
	transactions, err := controllers.GetTransactionByInvoiceID(uint(id), inv_type)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	transactions, err := controllers.GetAllTransactions()
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    transactions,
	})
}
