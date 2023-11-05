package models

import (
	"gorm.io/gorm"
)

type Invoice interface {
	GetTransactions() ([]Transaction, error)
	GetTotalTransaction() uint //GET TOTAL TRANSACTION
	PayTransaction(
		User, //PIC
		string, //PAYMENT TYPE
		uint, //PAYMENT ID
		uint, //NOMINAL
	) (InvoiceHistory, error)
}

type InvoiceForm struct {
	ID           string            `json:"id"`
	InvoiceType  string            `json:"invoice_type"`
	ClientID     uint              `json:"client_id"`
	Transactions []TransactionForm `json:"transactions"`
}

type InvoiceHistory struct {
	gorm.Model
	PersonInChargeID uint   `json:"person_in_charge_id"`
	PersonInCharge   User   `json:"person_in_charge"`
	InvoiceID        uint   `json:"invoice_id"`   //BISA DEBIT OR CREDIT
	InvoiceType      string `json:"invoice_type"` //DEBIT/CREDIT
	PaymentType      string `json:"payment_type"` //CASH(KWITANSI), GIRO, QRIS, TRF
	PaymentID        uint   `json:"payment_number"`
	Payment          uint   `json:"payment"`
}
