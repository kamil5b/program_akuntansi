package controllers

import (
	"program_akuntansi/accountancy_service/models"
	"program_akuntansi/accountancy_service/repositories"
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

func GetItemByBarcode(barcode uint) (models.Item, error) {
	return repositories.GetItem("barcode = ?", barcode)
}

func GetAllItems() ([]models.Item, error) {
	return repositories.GetAllItems()
}

func GetItemFamilyByID(id uint) ([]models.Item, error) {
	items := []models.Item{}
	item, err := GetItemByID(id)
	if err != nil {
		return nil, err
	}
	tmp_item := item
	for tmp_item.ID != 0 {
		tmp_item, _ := GetItemByID(item.SubitemID)
		if tmp_item.ID != 0 {
			item = tmp_item
		}
	}

	items = append(items, item)
	for item.SubitemID != 0 && err != nil {
		item, err = GetItemByID(item.SubitemID)
		items = append(items, item)
	}
	if err != nil {
		return items, err
	}
	return items, nil
}
