package services

import (
	"net/http"
	"program_akuntansi/utilities"
	"strconv"
)

// PROCESS

func AuthUser(auth string) (int, error) {
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
		return 0, err
	}
	return strconv.Atoi(response["data"]["sub"])
}
