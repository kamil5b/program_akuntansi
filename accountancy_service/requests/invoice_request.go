package requests

import (
	"program_akuntansi/accountancy_service/controllers"
	"program_akuntansi/accountancy_service/models"
	"strings"

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
			"message": err.Error(),
		})
	}
	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_INVOICE_CREATE"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
			})
		}
	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_INVOICE_CREATE"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
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
			"message": err.Error(),
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
		InvoiceType  string                   `json:"invoice_type"`
		Transactions []models.TransactionForm `json:"transactions"`
	}
	/*
		Authorization Header
		{
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
			"message": err.Error(),
		})
	}

	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_INVOICE_CREATE"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
			})
		}
	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_INVOICE_CREATE"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
			})
		}
	} else {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Invalid invoice type",
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}
	err = controllers.InputTransactionToInvoice(uint(id), data.InvoiceType, data.Transactions)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
			"message": err.Error(),
		})
	}
	if data.InvoiceType == "DEBIT" {
		if err := AuthUser(c, "AUTH_DEBIT_PAY"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
			})
		}

	} else if data.InvoiceType == "CREDIT" {
		if err := AuthUser(c, "AUTH_CREDIT_PAY"); err != nil {
			c.Status(403)

			return c.JSON(fiber.Map{
				"status":  403,
				"message": err.Error(),
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
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":         201,
		"message":        "success",
		"transaction_id": id,
	})
}

// Invoice History

func GetAllInvoiceHistory(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	invoice, err := controllers.GetAllInvoiceHistories()
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetInvoiceHistoryByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetInvoiceHistoryByID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetInvoiceHistoryByPaymentID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVOICE_PAYMENT_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		invoice, err := controllers.GetInvoiceHistoriesByCash()
		if err != nil {
			c.Status(400)

			return c.JSON(fiber.Map{
				"status":  400,
				"message": err.Error(),
			})
		}

		c.Status(201)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
			"data":    invoice,
		})
	}

	invoice, err := controllers.GetInvoiceHistoryByPaymentID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetInvoiceHistoriesByPICID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVOICE_PIC_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetInvoiceHistoriesByPICID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetInvoiceHistoriesByInvoiceID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_HISTORY_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetInvoiceHistoriesByInvoiceID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetInvoiceHistoriesByInvoiceType(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	inv_type := strings.ToUpper(c.Params("inv_type"))
	if inv_type != "DEBIT" && inv_type != "CREDIT" {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "invalid invoice type",
		})
	}

	if err := AuthUser(c, "AUTH_GET_HISTORY_TYPE"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	invoice, err := controllers.GetInvoiceHistoriesByInvoiceType(inv_type)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})

}

func GetInvoiceHistoriesByInvoiceTypeID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	inv_type := strings.ToUpper(c.Params("inv_type"))
	if inv_type != "DEBIT" && inv_type != "CREDIT" {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "invalid invoice type",
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	if err := AuthUser(c, "AUTH_GET_INVOICE_PIC_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	invoice, err := controllers.GetInvoiceHistoriesByInvIDType(uint(id), inv_type)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

// Invoice Debit

func GetDebitInvoiceByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_DEBIT_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetDebitInvoiceByID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetDebitInvoiceByClientID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_DEBIT_CLIENT_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetDebitInvoicesByClientID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetAllDebitInvoice(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_DEBIT_INVOICE"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	invoice, err := controllers.GetAllDebitInvoices()
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

// Invoice Credit

func GetCreditInvoiceByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_CREDIT_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id := c.Params("id")

	invoice, err := controllers.GetCreditInvoiceByID(id)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetCreditInvoiceByClientID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_CREDIT_INVOICE_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	invoice, err := controllers.GetCreditInvoicesByClientID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}

func GetAllCreditInvoice(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_CREDIT_INVOICE"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	invoice, err := controllers.GetAllCreditInvoicess()
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    invoice,
	})
}
