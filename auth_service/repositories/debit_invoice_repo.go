package repositories

import (
	"program_akuntansi/auth_service/database"
	"program_akuntansi/auth_service/models"
)

//======GET======

// TO KNOW IF THE DEBIT INVOICE EXIST OR NOT
func IsDebitInvoiceExist(query string, val ...interface{}) bool {
	var debit_invoice models.DebitInvoice
	database.DB.Where(query, val...).Last(&debit_invoice)
	return debit_invoice.ID != 0
}

// TO GET A DEBIT INVOICE
func GetDebitInvoice(query string, val ...interface{}) (models.DebitInvoice, error) {
	var debit_invoice models.DebitInvoice
	db := database.DB.Where(query, val...).Last(&debit_invoice)
	if db.Error != nil {
		return debit_invoice, db.Error
	}
	return debit_invoice, nil
}

// TO GET AN ARRAY OF DEBIT INVOICES (NOT ALL BUT CAN ALL)
func GetDebitInvoices(query string, val ...interface{}) ([]models.DebitInvoice, error) {
	var debit_invoices []models.DebitInvoice
	db := database.DB.Where(query, val...).Find(&debit_invoices)
	if db.Error != nil {
		return debit_invoices, db.Error
	}
	return debit_invoices, nil
}

// TO GET ALL DEBIT INVOICES
func GetAllDebitInvoices() ([]models.DebitInvoice, error) {
	var debit_invoices []models.DebitInvoice
	db := database.DB.Find(&debit_invoices)
	if db.Error != nil {
		return debit_invoices, db.Error
	}
	return debit_invoices, nil
}

// CREATE DEBIT INVOICE
func CreateDebitInvoice(debit_invoice models.DebitInvoice) (uint, error) {
	db := database.DB.Create(&debit_invoice)
	return debit_invoice.ID, db.Error
}

// UPDATE DEBIT INVOICE
func UpdateDebitInvoice(updated models.DebitInvoice, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.DebitInvoice{}).Where(query, val...).Updates(&updated)
	return db.Error
}

// DELETE DEBIT INVOICE
func DeleteDebitInvoice(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.DebitInvoice{})
	return db.Error
}
