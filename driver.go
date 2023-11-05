package main

import (
	"program_akuntansi/accountancy_service"
	"program_akuntansi/auth_service"
	"program_akuntansi/utilities"
)

func main() {
	var (
		acc_url  = utilities.GoDotEnvVariable("ACCOUNTANCY_URL")
		auth_url = utilities.GoDotEnvVariable("AUTH_URL")
		db_url   = utilities.GoDotEnvVariable("DATABASE_URL")
		user     = utilities.GoDotEnvVariable("DATABASE_USER")
		password = utilities.GoDotEnvVariable("DATABASE_PASSWORD")
		protocol = utilities.GoDotEnvVariable("DSN_PROTOCOL")
		db       = utilities.GoDotEnvVariable("DATABASE_NAME")
	)
	go auth_service.SetupTemplate(auth_url, db_url, user, password, protocol, db)
	accountancy_service.SetupTemplate(acc_url, db_url, user, password, protocol, db)
}
