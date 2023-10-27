package controllers

import (
	"program_akuntansi/models"
	"program_akuntansi/repositories"
)

// ===== CREATE INVOICE =====

// CREATE DEBIT INVOICE
func DebitInvoiceCreate(debit_invoice models.DebitInvoice) (uint, error) {
	return repositories.CreateDebitInvoice(debit_invoice)
}

// CREATE CREDIT INVOICE
func CreditInvoiceCreate(credit_invoice models.CreditInvoice) (uint, error) {
	return repositories.CreateCreditInvoice(credit_invoice)
}

// ===== PAY INVOICE ======

// PAY TRANSACTION
func PayTransaction(invoice models.Invoice, PIC models.User, payment_type string, payment_id uint, nominal uint) (uint, error) {
	invoice_history, err := invoice.PayTransaction(PIC, payment_type, payment_id, nominal)
	if err != nil {
		return 0, err
	}
	return repositories.CreateInvoiceHistory(invoice_history)
}

// PAY DEBIT TRANSACTION
func PayDebitTransaction(invoice_id uint, PIC models.User, payment_type string, payment_id uint, nominal uint) (uint, error) {
	invoice, err := GetDebitInvoiceByID(invoice_id)
	if err != nil {
		return 0, err
	}
	return PayTransaction(&invoice, PIC, payment_type, payment_id, nominal)
}

// PAY CREDIT TRANSACTION
func PayCreditTransaction(invoice_id uint, PIC models.User, payment_type string, payment_id uint, nominal uint) (uint, error) {
	invoice, err := GetCreditInvoiceByID(invoice_id)
	if err != nil {
		return 0, err
	}
	return PayTransaction(&invoice, PIC, payment_type, payment_id, nominal)
}

// ==== GET ======

// == GET BY ID ==

// GET INVOICE HISTORY BY HISTORY ID
func GetInvoiceHistoryByID(id uint) (models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistory("ID = ?", id)
}

// GET INVOICE HISTORY BY PAYMENT ID
func GetInvoiceHistoryByPaymentID(id uint) (models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistory("payment_id = ?", id)
}

// GET DEBIT INVOICE BY ID
func GetDebitInvoiceByID(id uint) (models.DebitInvoice, error) {
	return repositories.GetDebitInvoice("ID = ?", id)
}

// GET CREDIT INVOICE BY ID
func GetCreditInvoiceByID(id uint) (models.CreditInvoice, error) {
	return repositories.GetCreditInvoice("ID = ? OR invoice_credit_id = ?", id, id)
}

// == GET SOME (ARRAY) ==

// = INVOICE HISTORY =

// GET INVOICE HISTORIES BY PIC ID
func GetInvoiceHistoriesByPICID(id uint) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("person_in_charge_id = ?", id)
}

// GET INVOICE HISTORIES BY INVOICE ID
func GetInvoiceHistoriesByInvID(id uint) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("invoice_id = ?", id)
}

// GET INVOICE HISTORIES BY INVOICE ID & TYPE
func GetInvoiceHistoriesByInvType(invType string) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("invoice_type = ?", invType)
}

// GET INVOICE HISTORIES BY INVOICE ID & TYPE
func GetInvoiceHistoriesByInvIDType(id uint, invType string) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("invoice_id = ? and invoice_type = ?", id, invType)
}

// GET INVOICE HISTORIES BY INVOICE & PAYMENT TYPE
func GetInvoiceHistoriesByPayType(payType string) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("payment_type = ?", payType)
}

// GET INVOICE HISTORIES BY INVOICE & PAYMENT TYPE
func GetInvoiceHistoriesByInvPayType(invType, payType string) ([]models.InvoiceHistory, error) {
	return repositories.GetInvoiceHistories("invoice_type = ? and payment_type = ?", invType, payType)
}

// = DEBIT INVOICE =

// GET DEBIT INVOICES BY PIC ID
func GetDebitInvoicesByClientID(id uint) ([]models.DebitInvoice, error) {
	return repositories.GetDebitInvoices("client_id = ?", id)
}

// = CREDIT INVOICE =

// GET CREDIT INVOICES BY PIC ID
func GetCreditInvoicesByClientID(id uint) ([]models.CreditInvoice, error) {
	return repositories.GetCreditInvoices("store_id = ?", id)
}

// == GET ALL ==

// GET ALL INVOICE HISTORY
func GetAllInvoiceHistories() ([]models.InvoiceHistory, error) {
	return repositories.GetAllInvoiceHistories()
}

// GET ALL DEBIT INVOICES
func GetAllDebitInvoices() ([]models.DebitInvoice, error) {
	return repositories.GetAllDebitInvoices()
}

// GET ALL CREDIT INVOICES
func GetAllCreditInvoicess() ([]models.CreditInvoice, error) {
	return repositories.GetAllCreditInvoices()
}
