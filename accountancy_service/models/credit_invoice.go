package models

import (
	"gorm.io/gorm"
)

type CreditInvoice struct { //PEMBELIAN
	gorm.Model
	InvoiceCreditID string `json:"invoice_credit_id"`
	StoreID         uint   `json:"store_id"`
	Store           Store  `json:"store"`
	Debt            uint   `json:"debt"` //REMAINING PAYMENT
}

func (c CreditInvoice) GetTransactions() ([]Transaction, error) {
	return getTransactions(c.ID, "credit_invoices")
}

func (credit_invoice CreditInvoice) GetTotalTransaction() uint {
	return getTotalTransaction(credit_invoice.ID, "credit_invoices")
}

func (ci CreditInvoice) PayTransaction(PIC User, payment_type string, payment_id uint, nominal uint) (InvoiceHistory, error) {
	total_price := ci.GetTotalTransaction()
	return payTransaction(PIC, payment_type, payment_id, nominal, ci.ID, "CREDIT", total_price)
}
