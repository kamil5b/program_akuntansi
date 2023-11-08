package models

import (
	"errors"
	"program_akuntansi/accountancy_service/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func getTransactions(invoice_id uint, table_name string) ([]Transaction, error) {
	var transactions []Transaction
	db := database.DB.Where("invoice_id = ? AND invoice_type = ?", invoice_id, table_name).Preload(clause.Associations).Find(&transactions)
	if db.Error != nil {
		return transactions, db.Error
	}
	return transactions, nil
}

func getTotalTransaction(invoice_id uint, table_name string) uint {
	var Total, total_discount uint
	database.DB.Table("transactions").Select("sum(total_price) as total").Where("invoice_id = ? AND invoice_type = ?", invoice_id, table_name).Find(&Total)
	database.DB.Table("transactions").Select("sum(discount) as total_discount").Where("invoice_id = ? AND invoice_type = ?", invoice_id, table_name).Find(&total_discount)
	return Total - total_discount
}

func payTransaction(PIC User, payment_type string, payment_id uint, nominal, invoice_id uint, invoice_type string, total_price uint) (InvoiceHistory, error) {
	var Payed uint = 0
	database.DB.Table("invoice_histories").Select("sum(payment) as payed").Where("invoice_id = ? AND invoice_type = ?", invoice_id, invoice_type).Find(&Payed)
	margin := total_price - Payed
	if margin == 0 {
		return InvoiceHistory{}, errors.New("the invoice is payed fully")
	}
	if int(margin)-int(nominal) < 0 {
		return InvoiceHistory{}, errors.New("nominal is too much")
	}
	debt := margin - nominal
	table := ""
	if invoice_type == "DEBIT" {
		table = "debit_invoices"
	} else if invoice_type == "CREDIT" {
		table = "credit_invoices"
	}
	db := database.DB.Table(table).Where("ID = ?", invoice_id).Update("debt", debt)
	if db.Error != nil {
		return InvoiceHistory{}, db.Error
	}

	return InvoiceHistory{
		PersonInChargeID: PIC.ID,
		InvoiceID:        invoice_id,
		InvoiceType:      invoice_type,
		PaymentType:      payment_type,
		PaymentID:        payment_id,
		Payment:          nominal,
	}, nil
}
