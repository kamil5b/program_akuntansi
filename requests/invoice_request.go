package requests

import (
	"program_akuntansi/controllers"
	"program_akuntansi/models"

	"github.com/gofiber/fiber/v2"
)

func CreateInvoice(c *fiber.Ctx) error { //POST
	var data models.InvoiceForm
	/*
		Authorization Header
		id 				: string //nomor invoice
		invoice_type	: string // debit/credit
		client_id		: uint //harus di regis dulu
		transactions	: [
			{
				item_id 	: uint
				unit		: uint
				total_price	: uint
				discount	: uint
			},
			{
				item_id 	: uint
				unit		: uint
				total_price	: uint
				discount	: uint
			},
		]

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}
	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_INVOICE_CREATE"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}
	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_INVOICE_CREATE"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}
	} else {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Invalid invoice type",
		})
	}

	id, err := controllers.CreateInvoice(data)
	if err != nil {
		if id == 0 {
			c.Status(201)

			return c.JSON(fiber.Map{
				"status":     201,
				"message":    "success transaction failed",
				"invoice_id": id,
			})
		}
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":     201,
		"message":    "success",
		"invoice_id": id,
	})
}

func InputTransaction(c *fiber.Ctx) error {
	var data struct {
		InvoiceID    uint                     `json:"invoice_id"`
		InvoiceType  string                   `json:"invoice_type"`
		Transactions []models.TransactionForm `json:"transactions"`
	}
	/*
		Authorization Header
		{
			invoice_id: 	uint   //BISA DEBIT OR CREDIT
			invoice_type: 	string //DEBIT/CREDIT
			transactions: [
				{
					item_id :		uint
					unit :			uint
					total_price : 	uint
					discount : 		uint
				},
				{
					item_id :		uint
					unit :			uint
					total_price : 	uint
					discount : 		uint
				}
			]
		}
	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_INVOICE_CREATE"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}
	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_INVOICE_CREATE"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}
	} else {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Invalid invoice type",
		})
	}

	err := controllers.InputTransactionToInvoice(data.InvoiceID, data.InvoiceType, data.Transactions)
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
	})
}

func PayTransaction(c *fiber.Ctx) error {
	var data models.InvoiceHistory
	/*
		Authorization Header
		{
			person_in_charge_id: uint //user id
			invoice_id: 		 uint   //BISA DEBIT OR CREDIT
			invoice_type: 		 string //DEBIT/CREDIT
			payment_type: 		 string //CASH(KWITANSI), GIRO, QRIS, TRF
			payment_number: 	 uint //PAYMENT ID (kalo cash 0)
			payment: 			 uint //Nominal
		}

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}
	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_PAY"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}

	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_PAY"); err != nil {
			c.Status(403)
			return c.JSON(fiber.Map{
				"status":  403,
				"message": err,
			})
		}
	} else {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Invalid invoice type",
		})
	}

	id, err := controllers.PayTransactionFromHistory(data)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":         201,
		"message":        "success",
		"transaction_id": id,
	})
}
