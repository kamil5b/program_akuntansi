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
	Invoice     Invoice   `json:"invoice"`
	TotalPrice  uint      `json:"total_price"`
	Discount    uint      `json:"discount"`
}
