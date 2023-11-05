package routes

import (
	"program_akuntansi/accountancy_service/requests"

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
	auth.Post("/register", requests.RegisterUserAuth)
	auth.Get("/user", requests.LoginUser)

	// ==Item==
	item := api.Group("/item")
	item.Get("/", requests.GetAllItem)
	item.Get("/all", requests.GetAllItem)
	item.Post("/create", requests.ItemCreate)
	item.Patch("/:id", requests.ItemUpdate)
	item.Get("/:id", requests.GetItemByID)
	item.Get("/family/:id", requests.GetItemFamilyByID)

	// ==Store==
	store := api.Group("/store")
	store.Get("/", requests.GetAllStore)
	store.Get("/all", requests.GetAllStore)
	store.Post("/create", requests.StoreCreate)
	store.Post("/:id", requests.StoreUpdate)
	store.Get("/:id", requests.GetStoreByID)

	// ==Inventory==
	inventory := api.Group("/inventory")
	inventory.Get("/", requests.GetAllInventory)
	inventory.Get("/all", requests.GetAllInventory)
	inventory.Post("/open/:id", requests.InventoryOpenItem)
	inventory.Get("/:id", requests.GetInventoryByID)
	inventory.Get("/current/:id", requests.GetCurrentInventoryByID)

	// ==Transaction==
	transaction := api.Group("/transaction")
	transaction.Get("/", requests.GetAllTransaction)
	transaction.Get("/all", requests.GetAllTransaction)
	transaction.Get("/:id", requests.GetTransactionByID)
	transaction.Get("/invoice/:id", requests.GetTransactionByInvoiceID)

	// ==Invoice==
	invoice := api.Group("/invoice")
	invoice.Post("/create", requests.CreateInvoice)
	invoice.Post("/create/:id", requests.InputTransaction)
	invoice.Post("/pay", requests.PayTransaction)
	// Invoice History
	history := invoice.Group("/history")
	history.Get("/", requests.GetAllInvoiceHistory)
	history.Get("/all", requests.GetAllInvoiceHistory)
	history.Get("/:id", requests.GetInventoryByID)
	history.Get("/payment/:id", requests.GetInvoiceHistoryByPaymentID)
	history.Get("/user/:id", requests.GetInvoiceHistoriesByPICID)
	history.Get("/invoice/:id", requests.GetInvoiceHistoriesByInvoiceID)
	history.Get("/debit", requests.GetInvoiceHistoriesDebit)
	history.Get("/credit", requests.GetInvoiceHistoriesCredit)
	history.Get("/:inv_type/:id", requests.GetInvoiceHistoriesByInvoiceIDType)
	// Debit Invoice
	debit := invoice.Group("/debit")
	debit.Get("/", requests.GetAllDebitInvoice)
	debit.Get("/all", requests.GetAllDebitInvoice)
	debit.Get("/:id", requests.GetDebitInvoiceByID)
	debit.Get("/client/:id", requests.GetDebitInvoiceByClientID)
	// Credit Invoice
	credit := invoice.Group("/credit")
	credit.Get("/", requests.GetAllCreditInvoice)
	credit.Get("/all", requests.GetAllCreditInvoice)
	credit.Get("/:id", requests.GetCreditInvoiceByID)
	credit.Get("/client/:id", requests.GetCreditInvoiceByClientID)

}
