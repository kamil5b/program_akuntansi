package accountancy_service

import (
	"fmt"

	"program_akuntansi/accountancy_service/database"
	"program_akuntansi/accountancy_service/models"
	"program_akuntansi/accountancy_service/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupServer(db_type, server_url, db_url, user, password, protocol, db string) {

	database.ConnectManual(
		db_type, db_url, user, password, protocol, db,
		&models.User{},
		&models.Account{},
		&models.Store{},
		&models.Item{},
		&models.InvoiceHistory{},
		&models.Inventory{},
		&models.CreditInvoice{},
		&models.DebitInvoice{},
		&models.Transaction{},
	)
	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}

func SetupServerDSN(db_type, server_url, dsn string) {

	database.ConnectDSN(
		db_type, dsn,
		&models.User{},
		&models.Account{},
		&models.Store{},
		&models.Item{},
		&models.InvoiceHistory{},
		&models.Inventory{},
		&models.CreditInvoice{},
		&models.DebitInvoice{},
		&models.Transaction{},
	)
	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}
