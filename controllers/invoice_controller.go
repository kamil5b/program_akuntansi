package controllers

import (
	"errors"
	"program_akuntansi/models"
	"program_akuntansi/repositories"
)

// ===== CREATE INVOICE =====

func CreateInvoice(form models.InvoiceForm) (uint, error) {
	var total_transaction uint = 0
	var id uint = 0
	var err error = nil
	for _, trans := range form.Transactions {
		total_transaction = total_transaction + trans.TotalPrice - trans.Discount
	}
	if form.InvoiceType == "DEBIT" {
		di := models.DebitInvoice{
			ClientID: form.ClientID,
			Debt:     total_transaction,
		}
		id, err = repositories.CreateDebitInvoice(di)
	} else if form.InvoiceType == "CREDIT" {
		ci := models.CreditInvoice{
			InvoiceCreditID: form.ID,
			StoreID:         form.ClientID,
			Debt:            total_transaction,
		}
		id, err = repositories.CreateCreditInvoice(ci)
	} else {
		return 0, errors.New("invoice type invalid")
	}
	if err != nil {
		return id, err
	}
	for _, ft := range form.Transactions {
		inve := models.Inventory{
			ItemID:      ft.ItemID,
			Unit:        ft.Unit,
			Transaction: form.InvoiceType,
		}
		inv_type := ""
		if form.InvoiceType == "CREDIT" {
			inve.PrevInventoryID = 0
			inv_type = "credit_invoice"
		} //BUAT ELSE IF DEBIT

		trans := models.Transaction{
			Inventory:   inve,
			InvoiceID:   id,
			InvoiceType: inv_type,
			TotalPrice:  ft.TotalPrice,
			Discount:    ft.Discount,
		}
		_, err := repositories.CreateTransaction(trans)
		if err != nil {
			return id, err
		}
	}
	return id, nil
}

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

func PayTransaction2(id_inv uint, inv_type string, pic_id uint, payment_type string, payment_id uint, nominal uint) (uint, error) {
	var invoice models.Invoice
	if inv_type == "DEBIT" {
		inv, err := GetDebitInvoiceByID(id_inv)
		if err != nil {
			return 0, err
		}
		invoice = inv
	} else if inv_type == "CREDIT" {
		inv, err := GetCreditInvoiceByID(id_inv)
		if err != nil {
			return 0, err
		}
		invoice = inv
	} else {
		return 0, errors.New("invalid type")
	}
	PIC, err := GetUserByID(pic_id)
	if err != nil {
		return 0, err
	}
	invoice_history, err := invoice.PayTransaction(PIC, payment_type, payment_id, nominal)
	if err != nil {
		return 0, err
	}
	return repositories.CreateInvoiceHistory(invoice_history)
}

func PayTransactionFromHistory(iv models.InvoiceHistory) (uint, error) {
	return PayTransaction2(
		iv.InvoiceID,
		iv.InvoiceType,
		iv.PersonInChargeID,
		iv.PaymentType,
		iv.PaymentID,
		iv.Payment,
	)
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
