package requests

import (
	"errors"
	"program_akuntansi/accountancy_service/controllers"
	"program_akuntansi/accountancy_service/models"

	"github.com/gofiber/fiber/v2"
)

//ADD

func StoreCreate(c *fiber.Ctx) error { //POST
	var data models.Store
	/*
		Authorization Header
		name 			: string
		nomor_telepon	: string
		address			: string

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	if err := AuthUser(c, "AUTH_STORE_CREATE"); err != nil {
		return err
	}

	id, err := controllers.StoreCreate(data)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":   201,
		"message":  "success",
		"store_id": id,
	})
}

//UPDATE

func StoreUpdate(c *fiber.Ctx) error { //POST
	var data models.Store
	/*
		Authorization Header
		id				: uint
		name 			: string
		nomor_telepon	: string
		address			: string

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	if err := AuthUser(c, "AUTH_STORE_UPDATE"); err != nil {
		return err
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}
	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": errors.New("id not valid"),
		})
	}
	err = controllers.StoreIDUpdate(data.ID, data)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
	})
}

//GET

func GetStoreByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header
		id : uint

	*/

	if err := AuthUser(c, "AUTH_GET_STORE_ID"); err != nil {
		return err
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}
	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": errors.New("id not valid"),
		})
	}

	store, err := controllers.GetStoreByID(uint(id))
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    store,
	})
}

func GetAllStore(c *fiber.Ctx) error { //GET
	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_STORE"); err != nil {
		return err
	}

	stores, err := controllers.GetAllStores()
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    stores,
	})
}
