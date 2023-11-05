package models

import (
	"errors"
	"program_akuntansi/accountancy_service/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreditInvoice struct { //PEMBELIAN
	gorm.Model
	InvoiceCreditID string `json:"invoice_credit_id"`
	StoreID         uint   `json:"store_id"`
	Store           Store  `json:"store"`
	Debt            uint   `json:"debt"` //REMAINING PAYMENT
}

func (c CreditInvoice) GetTransactions() ([]Transaction, error) {
	var transactions []Transaction
	db := database.DB.Where("invoice_id = ? AND invoice_type = ?", c.ID, "credit_invoices").Preload(clause.Associations).Find(&transactions)
	if db.Error != nil {
		return transactions, db.Error
	}
	return transactions, nil
}

func (credit_invoice CreditInvoice) GetTotalTransaction() uint {
	var Total, total_discount uint
	database.DB.Table("transactions").Select("sum(total_price) as total").Where("invoice_id = ? AND invoice_type = ?", credit_invoice.ID, "credit_invoices").Preload(clause.Associations).Find(&Total)
	database.DB.Table("transactions").Select("sum(discount) as total_discount").Where("invoice_id = ? AND invoice_type = ?", credit_invoice.ID, "credit_invoices").Preload(clause.Associations).Find(&total_discount)
	return Total - total_discount
}

func (ci CreditInvoice) PayTransaction(PIC User, payment_type string, payment_id uint, nominal uint) (InvoiceHistory, error) {
	var Payed uint = 0
	database.DB.Table("credit_invoices").Select("sum(payment) as payed").Where("invoice_id = ? AND invoice_type = ?", ci.ID, "CREDIT").Preload(clause.Associations).Find(&Payed)
	total_price := ci.GetTotalTransaction()
	margin := total_price - Payed
	if int(margin)-int(nominal) < 0 {
		return InvoiceHistory{}, errors.New("nominal is too much")
	}
	ci.Debt = margin - nominal

	credit_invoice := CreditInvoice{Debt: ci.Debt}
	db := database.DB.Model(CreditInvoice{}).Where("ID = ?", ci.ID).Updates(&credit_invoice)
	if db.Error != nil {
		return InvoiceHistory{}, db.Error
	}

	return InvoiceHistory{
		PersonInChargeID: PIC.ID,
		InvoiceID:        ci.ID,
		InvoiceType:      "CREDIT",
		PaymentType:      payment_type,
		PaymentID:        payment_id,
		Payment:          nominal,
	}, nil
}
