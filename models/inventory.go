package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	PrevInventoryID uint   `json:"prev_inventory_id"`
	ItemID          uint   `json:"item_id"`
	Item            Item   `json:"item"`
	Unit            uint   `json:"unit"`
	Transaction     string `json:"transaction"` // CREDIT = BARANG KELUAR ; DEBIT = BARANG MASUK
}
