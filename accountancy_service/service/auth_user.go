package services

import (
	"net/http"
	"program_akuntansi/accountancy_service/controllers"
	"program_akuntansi/accountancy_service/models"
	"program_akuntansi/utilities"
	"strconv"
)

// PROCESS

func AuthUser(auth string) (models.User, error) {
	auth_url := utilities.GoDotEnvVariable("AUTH_URL") //URL TO AUTH SERVICE
	var response map[string]map[string]string
	err := utilities.HTTPRequest(
		"GET",
		auth_url,
		http.Header{
			"Authorization": {auth},
			"Content-Type":  {"application/json"},
		},
		nil,
		&response,
	)
	if err != nil {
		return models.User{}, err
	}
	acc_id, err := strconv.Atoi(response["data"]["sub"])

	if err != nil {
		return models.User{}, err
	}
	return controllers.GetUserByAccID(uint(acc_id))
}
