package repositories

import (
	"program_akuntansi/database"
	"program_akuntansi/models"

	"gorm.io/gorm/clause"
)

//======GET======

// TO KNOW IF THE ITEM EXIST OR NOT
func IsInventoryExist(query string, val ...interface{}) bool {
	var inventory models.Inventory
	database.DB.Where(query, val...).Preload(clause.Associations).Last(&inventory)
	return inventory.ID != 0
}

// TO GET A ITEM
func GetInventory(query string, val ...interface{}) (models.Inventory, error) {
	var inventory models.Inventory
	db := database.DB.Where(query, val...).Preload(clause.Associations).Last(&inventory)
	return inventory, db.Error
}

// TO GET AN ARRAY OF ITEMS (NOT ALL BUT CAN ALL)
func GetInventories(query string, val ...interface{}) ([]models.Inventory, error) {
	var inventories []models.Inventory
	db := database.DB.Where(query, val...).Preload(clause.Associations).Find(&inventories)
	return inventories, db.Error
}

// TO GET ALL ITEMS
func GetAllInventories() ([]models.Inventory, error) {
	var inventories []models.Inventory
	db := database.DB.Preload(clause.Associations).Find(&inventories)
	return inventories, db.Error
}

func GetCurrentUnitInventory(id uint) (uint, error) {
	inventory, err := GetInventory("ID = ?", id)
	if err != nil {
		return 0, err
	}
	total_unit := inventory.Unit
	var total uint
	db := database.DB.Table("inventories").Select("sum(unit) as total").Where("prev_inventory_id = ? AND metric like ?", id, inventory.Item.Metric).Preload(clause.Associations).Find(&total)
	if db.Error != nil {
		return 0, db.Error
	}
	return total_unit - total, nil
}

// CREATE ITEM
func CreateInventory(inventory models.Inventory) (uint, error) {
	db := database.DB.Create(&inventory)
	return inventory.ID, db.Error
}
