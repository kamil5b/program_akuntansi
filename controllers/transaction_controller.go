package controllers

import (
	"errors"
	"program_akuntansi/models"
	"program_akuntansi/repositories"
)

// CREATE
func TransactionCreate(transaction models.Transaction) (uint, error) {
	return repositories.CreateTransaction(transaction)
}

func InputTransactionToInvoice(id uint, invoice_type string, transactions []models.TransactionForm) error {
	var total_transaction uint = 0
	for _, trans := range transactions {
		total_transaction = total_transaction + trans.TotalPrice - trans.Discount
	}

	for _, ft := range transactions {

		trans := models.Transaction{
			InvoiceID:  id,
			TotalPrice: ft.TotalPrice,
			Discount:   ft.Discount,
		}

		if invoice_type == "CREDIT" {

			inve := models.Inventory{
				ItemID:          ft.ItemID,
				Unit:            ft.Unit,
				Transaction:     "DEBIT",
				PrevInventoryID: 0,
			}
			trans.InvoiceType = "credit_invoice"
			trans.Inventory = inve

		} else if invoice_type == "DEBIT" {
			trans.InvoiceType = "debit_invoice"

			inve_old, err := repositories.GetInventory("item_id = ? and transaction = ?", ft.ItemID, "DEBIT")
			if err != nil {
				return err
			}
			id_inve, err := InventoryOut(inve_old.ID, ft.Unit)
			if err != nil {
				return err
			}
			trans.InventoryID = id_inve

		}

		_, err := repositories.CreateTransaction(trans)
		if err != nil {
			return err
		}
	}

	return nil
}

// GET

func GetTransactionByID(id uint) (models.Transaction, error) {
	return repositories.GetTransaction("ID = ?", id)
}

func GetAllTransactions() ([]models.Transaction, error) {
	return repositories.GetAllTransactions()
}

// GET DEBIT

func GetDebitTransactionsByID(id uint) ([]models.Transaction, error) {
	debit_invoice, err := GetDebitInvoiceByID(id)
	if err != nil {
		return nil, err
	}
	return debit_invoice.GetTransactions()
}

func GetAllDebitTransactions() ([]models.Transaction, error) {
	return repositories.GetTransactions("invoice_type = ?", "debit_invoices")
}

// GET CREDIT

func GetCreditTransactionsByID(id uint) ([]models.Transaction, error) {
	credit_invoice, err := GetCreditInvoiceByID(id)
	if err != nil {
		return nil, err
	}
	return credit_invoice.GetTransactions()
}

func GetAllCreditTransactions() ([]models.Transaction, error) {
	return repositories.GetTransactions("invoice_type = ?", "credit_invoices")
}

// GET INVOICE

func GetTransactionByInvoice(id uint, invoice_type string) ([]models.Transaction, error) {
	if invoice_type == "DEBIT" {
		return GetDebitTransactionsByID(id)
	} else if invoice_type == "CREDIT" {
		return GetCreditTransactionsByID(id)
	}
	return nil, errors.New("invoice type invalid")
}
