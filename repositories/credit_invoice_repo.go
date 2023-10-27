package repositories

import (
	"program_akuntansi/database"
	"program_akuntansi/models"
)

//======GET======

// TO KNOW IF THE CREDIT INVOICE EXIST OR NOT
func IsCreditInvoiceExist(query string, val ...interface{}) bool {
	var credit_invoice models.CreditInvoice
	database.DB.Where(query, val...).Last(&credit_invoice)
	return credit_invoice.ID != 0
}

// TO GET A CREDIT INVOICE
func GetCreditInvoice(query string, val ...interface{}) (models.CreditInvoice, error) {
	var credit_invoice models.CreditInvoice
	db := database.DB.Where(query, val...).Last(&credit_invoice)
	return credit_invoice, db.Error
}

// TO GET AN ARRAY OF CREDIT INVOICES (NOT ALL BUT CAN ALL)
func GetCreditInvoices(query string, val ...interface{}) ([]models.CreditInvoice, error) {
	var credit_invoices []models.CreditInvoice
	db := database.DB.Where(query, val...).Find(&credit_invoices)
	return credit_invoices, db.Error
}

// TO GET ALL CREDIT INVOICES
func GetAllCreditInvoices() ([]models.CreditInvoice, error) {
	var credit_invoices []models.CreditInvoice
	db := database.DB.Find(&credit_invoices)
	return credit_invoices, db.Error
}

// CREATE CREDIT INVOICE
func CreateCreditInvoice(credit_invoice models.CreditInvoice) (uint, error) {
	db := database.DB.Create(&credit_invoice)
	return credit_invoice.ID, db.Error
}

// UPDATE CREDIT INVOICE
func UpdateCreditInvoice(updated models.CreditInvoice, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.CreditInvoice{}).Where(query, val...).Updates(&updated)
	return db.Error
}

// DELETE CREDIT INVOICE
func DeleteCreditInvoice(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.CreditInvoice{})
	return db.Error
}
