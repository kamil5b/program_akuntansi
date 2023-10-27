package controllers

import (
	"errors"
	"net/http"
	"program_akuntansi/models"
	"program_akuntansi/repositories"
	"program_akuntansi/utilities"
	"strconv"
)

// CREATE
func UserCreate(user models.User) (uint, error) {
	return repositories.CreateUser(user)
}

func RegisterUser(name, role string) (uint, error) {
	user := models.User{
		Name: name,
		Role: role,
	}
	return UserCreate(user)
}

func RegisterExistingUserAcc(acc_id, id uint) error {
	acc := models.Account{
		AuthID: acc_id,
		UserID: id,
	}
	err := repositories.CreateAccount(acc)
	return err
}

func RegisterAuthUser(auth, name, role string, body any) error {
	auth_url := utilities.GoDotEnvVariable("AUTH_URL") //URL TO AUTH SERVICE
	var response map[string]string
	err := utilities.HTTPRequest(
		"POST",
		auth_url,
		http.Header{
			"Authorization": {auth},
			"Content-Type":  {"application/json"},
		},
		body,
		&response,
	)
	if err != nil {
		return err
	}
	acc_id, err := strconv.Atoi(response["sub"])

	if err != nil {
		return err
	}
	acc := models.Account{
		AuthID: uint(acc_id),
		User: models.User{
			Name: name,
			Role: role,
		},
	}
	return repositories.CreateAccount(acc)
}

func RegisterExistingUser(id, acc_id uint) error {
	if !IsUserExistByID(id) {
		return errors.New("id user is not found")
	}
	if IsUserExistByAccID(id, acc_id) {
		return errors.New("account id registered already")
	}
	if IsAccExistByID(acc_id) {
		return errors.New("account has been registered in other user")
	}
	return repositories.CreateAccount(models.Account{
		AuthID: acc_id,
		UserID: id,
	})
}

// UPDATE
func UserIDUpdate(id uint, user models.User) error {
	return repositories.UpdateUser(user, "ID = ?", id)
}

// DELETE
func UserIDDelete(id uint) error {
	return repositories.DeleteUser("ID = ?", id)
}

// GET
func IsUserExistByID(id uint) bool {
	return repositories.IsUserExist("ID = ?", id)
}

func GetUserByID(id uint) (models.User, error) {
	return repositories.GetUser("ID = ?", id)
}

func IsUserExistByAccID(id, acc_id uint) bool {
	return repositories.IsAccountExist("user_id = ? and auth_id = ?", id, acc_id)
}

func GetUserByAccID(id uint) (models.User, error) {
	acc, err := repositories.GetAccount("auth_id = ?", id)
	return acc.User, err
}

func IsAccExistByID(id uint) bool {
	return repositories.IsAccountExist("auth_id = ?", id)
}

func GetAccByUserID(id uint) ([]models.Account, error) {
	return repositories.GetAccounts("user_id = ?", id)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// PROCESS

func AuthUser(auth string, body any) (models.User, error) {
	auth_url := utilities.GoDotEnvVariable("AUTH_URL") //URL TO AUTH SERVICE
	var response map[string]string
	err := utilities.HTTPRequest(
		"POST",
		auth_url,
		http.Header{
			"Authorization": {auth},
			"Content-Type":  {"application/json"},
		},
		body,
		&response,
	)
	if err != nil {
		return models.User{}, err
	}
	acc_id, err := strconv.Atoi(response["sub"])

	if err != nil {
		return models.User{}, err
	}
	return GetUserByAccID(uint(acc_id))
}
