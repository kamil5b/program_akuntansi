package models

import (
	"gorm.io/gorm"
)

type DebitInvoice struct { //PENJUALAN
	gorm.Model
	ClientID uint  `json:"client_id"`
	Client   Store `json:"client"`
	Debt     uint  `json:"debt"` //REMAINING PAYMENT
}

func (d DebitInvoice) GetTransactions() ([]Transaction, error) {
	return getTransactions(d.ID, "debit_invoices")
}

func (d DebitInvoice) GetTotalTransaction() uint {
	return getTotalTransaction(d.ID, "debit_invoices")
}

func (d DebitInvoice) PayTransaction(PIC User, payment_type string, payment_id uint, nominal uint) (InvoiceHistory, error) {
	total_price := d.GetTotalTransaction()
	return payTransaction(PIC, payment_type, payment_id, nominal, d.ID, "DEBIT", total_price)
}
