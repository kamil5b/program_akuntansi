package models

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name         string `json:"name"`
	NomorTelepon string `json:"nomor_telepon"`
	Address      string `json:"address"`
}
