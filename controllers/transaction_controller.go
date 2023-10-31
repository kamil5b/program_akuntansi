package controllers

import (
	"errors"
	"program_akuntansi/models"
	"program_akuntansi/repositories"
	"strconv"
)

// CREATE
func TransactionCreate(transaction models.Transaction) (uint, error) {
	return repositories.CreateTransaction(transaction)
}

// GET

func GetTransactionByID(id uint) (models.Transaction, error) {
	return repositories.GetTransaction("ID = ?", id)
}

func GetAllTransactions() ([]models.Transaction, error) {
	return repositories.GetAllTransactions()
}

// GET DEBIT

func GetTransactionsByDebitInvoiceID(id uint) ([]models.Transaction, error) {
	debit_invoice, err := GetDebitInvoiceByID(id)
	if err != nil {
		return nil, err
	}
	return debit_invoice.GetTransactions()
}

func GetAllDebitInvoiceTransactions() ([]models.Transaction, error) {
	return repositories.GetTransactions("invoice_type = ?", "debit_invoices")
}

// GET CREDIT

func GetTransactionsByCreditInvoiceID(id uint) ([]models.Transaction, error) {
	credit_invoice, err := GetCreditInvoiceByID(strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}
	return credit_invoice.GetTransactions()
}

func GetAllCreditInvoiceTransactions() ([]models.Transaction, error) {
	return repositories.GetTransactions("invoice_type = ?", "credit_invoices")
}

// GET INVOICE

func GetTransactionByInvoiceID(id uint, invoice_type string) ([]models.Transaction, error) {
	if invoice_type == "DEBIT" {
		return GetTransactionsByDebitInvoiceID(id)
	} else if invoice_type == "CREDIT" {
		return GetTransactionsByCreditInvoiceID(id)
	}
	return nil, errors.New("invoice type invalid")
}
