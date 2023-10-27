package main

import (
	"fmt"

	"program_akuntansi/database"
	"program_akuntansi/models"
	"program_akuntansi/routes"
	"program_akuntansi/utilities"

	"github.com/gofiber/fiber/v2"
)

func SetupTemplate(server_url, db_url, user, password, protocol, db string) {

	database.Connect(
		db_url, user, password, protocol, db,
		&models.User{},
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

func main() {
	var (
		server_url = utilities.GoDotEnvVariable("SERVER_URL")
		db_url     = utilities.GoDotEnvVariable("DATABASE_URL")
		user       = utilities.GoDotEnvVariable("DATABASE_USER")
		password   = utilities.GoDotEnvVariable("DATABASE_PASSWORD")
		protocol   = utilities.GoDotEnvVariable("DSN_PROTOCOL")
		db         = utilities.GoDotEnvVariable("DATABASE_NAME")
	)
	SetupTemplate(server_url, db_url, user, password, protocol, db)
}
