package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name         string `json:"name"`
	Barcode      uint   `json:"barcode"`
	Metric       string `json:"metric"`     //METRIC SUCH AS : KOTAK, STRIP, DLL
	SubitemID    uint   `json:"subitem_id"` //SUBMETRIC ID, KALO GAADA => 0
	Multiplier   uint   `json:"multiplier"` //MULTIPLIER OF THE SUBMETRIC CONTOH : 12 batang dalam 1 bungkus -> 12, Submetric -> batang
	PricePerUnit uint   `json:"price_per_unit"`
}
