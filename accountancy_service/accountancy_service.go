package accountancy_service

import (
	"fmt"

	"program_akuntansi/accountancy_service/database"
	"program_akuntansi/accountancy_service/models"
	"program_akuntansi/accountancy_service/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupTemplate(server_url, db_url, user, password, protocol, db string) {

	database.Connect(
		db_url, user, password, protocol, db,
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
	/*
		origin := utilities.GoDotEnvVariable("VIEW_URL") //ganti view url ini di .env
		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     []string{origin},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		}))
	*/
	routes.Setup(app)

	err := app.Listen(server_url)
	if err != nil {
		fmt.Println(err)
		fmt.Scan(&err)
	}
}
