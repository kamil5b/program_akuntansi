package requests

import (
	"program_akuntansi/controllers"
	"program_akuntansi/models"

	"github.com/gofiber/fiber/v2"
)

//ADD

func ItemCreate(c *fiber.Ctx) error { //POST
	var data models.Item
	/*
		Authorization Header
		name 			: string
		barcode			: uint
		metric			: string
		subitem_id		: uint
		multiplier		: uint
		price_per_unit 	: uint

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	if err := AuthUser(c, "AUTH_ITEM_CREATE"); err != nil {
		return err
	}

	id, err := controllers.ItemCreate(data)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": err,
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  201,
		"message": "success",
		"item_id": id,
	})
}

//UPDATE

func ItemUpdate(c *fiber.Ctx) error { //POST
	var data models.Item
	/*
		Authorization Header
		name 			: string
		barcode			: uint
		metric			: string
		subitem_id		: uint
		multiplier		: uint
		price_per_unit 	: uint

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	if err := AuthUser(c, "AUTH_ITEM_UPDATE"); err != nil {
		return err
	}

	err := controllers.ItemIDUpdate(data.ID, data)
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

func GetItemByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header
		id : uint

	*/

	if err := AuthUser(c, "AUTH_GET_ITEM_ID"); err != nil {
		return err
	}

	id := c.QueryInt("id", 0)

	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Item not found in query",
		})
	}

	item, err := controllers.GetItemByID(uint(id))
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
		"data":    item,
	})
}

func GetAllItem(c *fiber.Ctx) error { //GET
	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_ITEM"); err != nil {
		return err
	}

	items, err := controllers.GetAllItems()
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
		"data":    items,
	})
}

func GetItemFamilyByID(c *fiber.Ctx) error { //GET

	if err := AuthUser(c, "AUTH_GET_ITEM_FAMILY_ID"); err != nil {
		return err
	}

	id := c.QueryInt("id", 0)

	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Item not found in query",
		})
	}

	item, err := controllers.GetItemFamilyByID(uint(id))
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
		"data":    item,
	})
}
