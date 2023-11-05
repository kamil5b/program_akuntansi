package repositories

import (
	"program_akuntansi/auth_service/database"
	"program_akuntansi/auth_service/models"
)

//======GET======

// TO KNOW IF THE ITEM EXIST OR NOT
func IsItemExist(query string, val ...interface{}) bool {
	var item models.Item
	database.DB.Where(query, val...).Last(&item)
	return item.ID != 0
}

// TO GET A ITEM
func GetItem(query string, val ...interface{}) (models.Item, error) {
	var item models.Item
	db := database.DB.Where(query, val...).Last(&item)
	return item, db.Error
}

// TO GET AN ARRAY OF ITEMS (NOT ALL BUT CAN ALL)
func GetItems(query string, val ...interface{}) ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Where(query, val...).Find(&items)
	return items, db.Error
}

// TO GET ALL ITEMS
func GetAllItems() ([]models.Item, error) {
	var items []models.Item
	db := database.DB.Find(&items)
	return items, db.Error
}

// CREATE ITEM
func CreateItem(item models.Item) (uint, error) {
	db := database.DB.Create(&item)
	return item.ID, db.Error
}

// UPDATE ITEM
func UpdateItem(updated models.Item, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.Item{}).Where(query, val...).Updates(&updated)
	return db.Error
}

// DELETE ITEM
func DeleteItem(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.Item{})
	return db.Error
}
