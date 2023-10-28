package models

import (
	"errors"
	"program_akuntansi/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DebitInvoice struct { //PENJUALAN
	gorm.Model
	ClientID uint  `json:"client_id"`
	Client   Store `json:"client"`
	Debt     uint  `json:"debt"` //REMAINING PAYMENT
}

func (d DebitInvoice) GetTransactions() ([]Transaction, error) {
	var transactions []Transaction
	db := database.DB.Where("invoice_id = ? AND invoice_type = ?", d.ID, "debit_invoices").Preload(clause.Associations).Find(&transactions)
	if db.Error != nil {
		return transactions, db.Error
	}
	return transactions, nil
}

func (d DebitInvoice) GetTotalTransaction() uint {
	var Total, total_discount uint
	database.DB.Table("transactions").Select("sum(total_price) as total").Where("invoice_id = ? AND invoice_type = ?", d.ID, "debit_invoices").Preload(clause.Associations).Find(&Total)
	database.DB.Table("transactions").Select("sum(discount) as total_discount").Where("invoice_id = ? AND invoice_type = ?", d.ID, "debit_invoices").Preload(clause.Associations).Find(&total_discount)
	return Total - total_discount
}

func (d DebitInvoice) PayTransaction(PIC User, payment_type string, payment_id uint, nominal uint) (InvoiceHistory, error) {
	var Payed uint = 0
	database.DB.Table("debit_invoices").Select("sum(payment) as payed").Where("invoice_id = ? AND invoice_type = ?", d.ID, "DEBIT").Preload(clause.Associations).Find(&Payed)
	total_price := d.GetTotalTransaction()
	margin := total_price - Payed
	if int(margin)-int(nominal) < 0 {
		return InvoiceHistory{}, errors.New("nominal is too much")
	}
	d.Debt = margin - nominal

	debit_invoice := DebitInvoice{Debt: d.Debt}
	db := database.DB.Model(DebitInvoice{}).Where("ID = ?", d.ID).Updates(&debit_invoice)
	if db.Error != nil {
		return InvoiceHistory{}, db.Error
	}

	return InvoiceHistory{
		PersonInChargeID: PIC.ID,
		InvoiceID:        d.ID,
		InvoiceType:      "DEBIT",
		PaymentType:      payment_type,
		PaymentID:        payment_id,
		Payment:          nominal,
	}, nil
}
