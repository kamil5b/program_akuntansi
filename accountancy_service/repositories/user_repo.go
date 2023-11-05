package repositories

import (
	"program_akuntansi/accountancy_service/database"
	"program_akuntansi/accountancy_service/models"

	"gorm.io/gorm/clause"
)

//======GET======

// TO KNOW IF THE USER EXIST OR NOT
func IsUserExist(query string, val ...interface{}) bool {
	var user models.User
	database.DB.Where(query, val...).Last(&user)
	return user.ID != 0
}

func IsAccountExist(query string, val ...interface{}) bool {
	var account models.Account
	database.DB.Where(query, val...).Last(&account)
	return account.AuthID != 0
}

// TO GET A USER
func GetUser(query string, val ...interface{}) (models.User, error) {
	var user models.User
	db := database.DB.Where(query, val...).Last(&user)
	return user, db.Error
}

func GetAccount(query string, val ...interface{}) (models.Account, error) {
	var account models.Account
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&account)
	return account, db.Error
}

// TO GET AN ARRAY OF USERS (NOT ALL BUT CAN ALL)
func GetUsers(query string, val ...interface{}) ([]models.User, error) {
	var users []models.User
	db := database.DB.Where(query, val...).Find(&users)
	return users, db.Error
}

func GetAccounts(query string, val ...interface{}) ([]models.Account, error) {
	var accounts []models.Account
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&accounts)
	return accounts, db.Error
}

// TO GET ALL USERS
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	db := database.DB.Find(&users)
	return users, db.Error
}

func GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	db := database.DB.Preload(clause.Associations).Find(&accounts)
	return accounts, db.Error
}

// CREATE USER
func CreateUser(user models.User) (uint, error) {
	db := database.DB.Create(&user)
	return user.ID, db.Error
}

func CreateAccount(account models.Account) error {
	db := database.DB.Create(&account)
	return db.Error
}

// UPDATE USER
func UpdateUser(updated models.User, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.User{}).Where(query, val...).Updates(&updated)
	return db.Error
}

func UpdateAccount(updated models.Account, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.Account{}).Where(query, val...).Updates(&updated)
	return db.Error
}

// DELETE USER
func DeleteUser(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.User{})
	return db.Error
}

func DeleteAccount(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.Account{})
	return db.Error
}
