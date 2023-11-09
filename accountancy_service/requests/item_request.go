package requests

import (
	"program_akuntansi/accountancy_service/controllers"
	"program_akuntansi/accountancy_service/models"

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
			"message": err.Error(),
		})
	}

	if err := AuthUser(c, "AUTH_ITEM_CREATE"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := controllers.ItemCreate(data)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  401,
			"message": err.Error(),
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
			"message": err.Error(),
		})
	}
	if err := AuthUser(c, "AUTH_ITEM_UPDATE"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	err = controllers.ItemIDUpdate(uint(id), data)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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

	*/

	if err := AuthUser(c, "AUTH_GET_ITEM_ID"); err != nil {
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	item, err := controllers.GetItemByID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	items, err := controllers.GetAllItems()
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
		c.Status(403)

		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}
	item, err := controllers.GetItemFamilyByID(uint(id))
	if err != nil {
		c.Status(400)

		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    item,
	})
}
