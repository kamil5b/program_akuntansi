package repositories

import (
	"program_akuntansi/auth_service/database"
	"program_akuntansi/auth_service/models"

	"gorm.io/gorm/clause"
)

// TO GET A TRANSACTION
func GetTransaction(query string, val ...interface{}) (models.Transaction, error) {
	var transaction models.Transaction
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&transaction)
	return transaction, db.Error
}

// TO GET AN ARRAY OF TRANSACTIONS (NOT ALL BUT CAN ALL)
func GetTransactions(query string, val ...interface{}) ([]models.Transaction, error) {
	var transactions []models.Transaction
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&transactions)
	return transactions, db.Error
}

// TO GET ALL TRANSACTIONS
func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	db := database.DB.Preload(clause.Associations).Find(&transactions)
	return transactions, db.Error
}

// CREATE TRANSACTION
func CreateTransaction(transaction models.Transaction) (uint, error) {
	db := database.DB.Create(&transaction)
	return transaction.ID, db.Error
}
