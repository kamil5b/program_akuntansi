package controllers

import (
	"program_akuntansi/models"
	"program_akuntansi/repositories"
)

// CREATE
func ItemCreate(item models.Item) (uint, error) {
	return repositories.CreateItem(item)
}

// UPDATE
func ItemIDUpdate(id uint, item models.Item) error {
	return repositories.UpdateItem(item, "ID = ?", id)
}

// // DELETE
// func ItemIDDelete(id uint) error {
// 	return repositories.DeleteItem("ID = ?", id)
// }

// GET
func GetItemByID(id uint) (models.Item, error) {
	return repositories.GetItem("ID = ?", id)
}

func GetAllItems() ([]models.Item, error) {
	return repositories.GetAllItems()
}
