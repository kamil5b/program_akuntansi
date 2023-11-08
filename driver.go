package main

import (
	"program_akuntansi/accountancy_service"
	"program_akuntansi/auth_service"
	"program_akuntansi/utilities"
)

func main() {
	var (

		//server_url  = utilities.GoDotEnvVariable("SERVER_PORT")
		acc_port    = utilities.GoDotEnvVariable("PORT")
		auth_server = utilities.GoDotEnvVariable("AUTH_SERVER")
		db_url      = utilities.GoDotEnvVariable("DATABASE_URL")
		user        = utilities.GoDotEnvVariable("DATABASE_USER")
		password    = utilities.GoDotEnvVariable("DATABASE_PASSWORD")
		protocol    = utilities.GoDotEnvVariable("DSN_PROTOCOL")
		db          = utilities.GoDotEnvVariable("DATABASE_NAME")
	)
	go auth_service.SetupTemplate(auth_server, db_url, user, password, protocol, db)
	accountancy_service.SetupTemplate(acc_port, db_url, user, password, protocol, db)
}
