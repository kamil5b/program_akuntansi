package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	InventoryID uint      `json:"inventory_id"`
	Inventory   Inventory `json:"inventory"`
	InvoiceID   uint      `json:"invoice_id"`
	InvoiceType string    `json:"invoice_type"` //Credit/Debit (Out/In)
	TotalPrice  uint      `json:"total_price"`
	Discount    uint      `json:"discount"`
}

type TransactionForm struct {
	ItemID     uint `json:"item_id"`
	Unit       uint `json:"unit"`
	TotalPrice uint `json:"total_price"`
	Discount   uint `json:"discount"`
}
