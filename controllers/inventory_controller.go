package controllers

import (
	"errors"
	"program_akuntansi/models"
	"program_akuntansi/repositories"
)

// CREATE
// HARUS DARI TRANSACTION
func InventoryCreate(inventory models.Inventory) (uint, error) {
	return repositories.CreateInventory(inventory)
}

// UPDATE

func InventoryOpenItem(id uint, open_unit uint) (uint, error) {
	inventory, err := GetInventoryByID(id)
	if inventory.Item.SubitemID == 0 {
		return 0, errors.New("no submetrics")
	} else if int(inventory.Unit)-int(open_unit) < 0 {
		return 0, errors.New("cannot open inventory, open unit > available unit")
	} else if err != nil {
		return 0, err
	}
	new_item, err := GetItemByID(inventory.Item.SubitemID)
	if err != nil {
		return 0, err
	}
	debit_inventory := models.Inventory{
		PrevInventoryID: id,
		ItemID:          new_item.ID,
		Unit:            inventory.Unit * open_unit * new_item.Multiplier,
		Transaction:     "DEBIT",
	}
	_, err = InventoryCreate(debit_inventory)
	if err != nil {
		return 0, err
	}
	credit_inventory := models.Inventory{
		PrevInventoryID: id,
		ItemID:          inventory.ItemID,
		Unit:            inventory.Unit - open_unit,
		Transaction:     "CREDIT",
	}
	return InventoryCreate(credit_inventory)
}

// INVENTORY OUT HARUS ADA TRANSACTION
func InventoryOut(id, out_unit uint) (uint, error) {
	inventory, err := GetInventoryByID(id)
	if int(inventory.Unit)-int(out_unit) < 0 {
		return 0, errors.New("cannot out inventory, out unit > available unit")
	} else if err != nil {
		return 0, err
	}
	credit_inventory := models.Inventory{
		PrevInventoryID: id,
		ItemID:          inventory.ItemID,
		Unit:            inventory.Unit - out_unit,
		Transaction:     "CREDIT",
	}
	return InventoryCreate(credit_inventory)
}

// GET

func GetInventoryByID(id uint) (models.Inventory, error) {
	return repositories.GetInventory("ID = ?", id)
}

func GetCurrentInventoryByID(id uint) (models.Inventory, error) {
	inventory, err := GetInventoryByID(id)
	if err != nil {
		return models.Inventory{}, err
	}
	units, err := repositories.GetCurrentUnitInventory(id)
	if err != nil {
		return models.Inventory{}, err
	}
	inventory.Unit = units
	return inventory, nil
}

func GetAllInventories() ([]models.Inventory, error) {
	return repositories.GetAllInventories()
}

func GetInventoriesByItemID(id uint) ([]models.Inventory, error) {
	return repositories.GetInventories("item_id = ?", id)
}

func GetCurrentInventoriesByItemID(id uint) ([]models.Inventory, error) {
	inventories, err := GetInventoriesByItemID(id)
	if err != nil {
		return nil, err
	}
	out_inv := []models.Inventory{}
	for _, inv := range inventories {
		c_inv, err := GetCurrentInventoryByID(inv.ID)
		if err != nil {
			return nil, err
		}
		out_inv = append(out_inv, c_inv)
	}
	return out_inv, nil
}
