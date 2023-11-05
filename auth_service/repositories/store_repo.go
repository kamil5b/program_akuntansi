package repositories

import (
	"program_akuntansi/auth_service/database"
	"program_akuntansi/auth_service/models"
)

//======GET======

// TO KNOW IF THE STORE EXIST OR NOT
func IsStoreExist(query string, val ...interface{}) bool {
	var store models.Store
	database.DB.Where(query, val...).Last(&store)
	return store.ID != 0
}

// TO GET A STORE
func GetStore(query string, val ...interface{}) (models.Store, error) {
	var store models.Store
	db := database.DB.Where(query, val...).Last(&store)
	return store, db.Error
}

// TO GET AN ARRAY OF STORES (NOT ALL BUT CAN ALL)
func GetStores(query string, val ...interface{}) ([]models.Store, error) {
	var stores []models.Store
	db := database.DB.Where(query, val...).Find(&stores)
	return stores, db.Error
}

// TO GET ALL STORES
func GetAllStores() ([]models.Store, error) {
	var stores []models.Store
	db := database.DB.Find(&stores)
	return stores, db.Error
}

// CREATE STORE
func CreateStore(store models.Store) (uint, error) {
	db := database.DB.Create(&store)
	return store.ID, db.Error
}

// UPDATE STORE
func UpdateStore(updated models.Store, query string, val ...interface{}) error { //UPDATE
	db := database.DB.Model(models.Store{}).Where(query, val...).Updates(&updated)
	return db.Error
}

// DELETE STORE
func DeleteStore(query string, val ...interface{}) error { //DELETE
	db := database.DB.Where(query, val...).Delete(&models.Store{})
	return db.Error
}
