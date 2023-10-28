package requests

import (
	"program_akuntansi/controllers"
	"program_akuntansi/utilities"

	"github.com/gofiber/fiber/v2"
)

//UPDATE

func InventoryOpenItem(c *fiber.Ctx) error { //POST
	var data map[string]string
	/*
		Authorization Header
		id 	: uint
		open_item	: uint

	*/
	if err := c.BodyParser(&data); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err,
		})
	}

	if err := AuthUser(c, "AUTH_INVENTORY_OPEN"); err != nil {
		return err
	}
	dataint := utilities.MapStringToInt(data)
	id, err := controllers.InventoryOpenItem(uint(dataint["id"]), uint(dataint["open_item"]))
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
		"data":    id,
	})
}

//GET

func GetInventoryByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVENTORY_ID"); err != nil {
		return err
	}

	id := c.QueryInt("id", 0)

	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Inventory not found in query",
		})
	}

	inventory, err := controllers.GetInventoryByID(uint(id))
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
		"data":    inventory,
	})
}

func GetAllInventory(c *fiber.Ctx) error { //GET
	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_ALL_INVENTORY"); err != nil {
		return err
	}

	inventories, err := controllers.GetAllInventories()
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
		"data":    inventories,
	})
}

func GetCurrentInventoryByID(c *fiber.Ctx) error { //GET

	if err := AuthUser(c, "AUTH_GET_CURRENT_INVENTORY_ID"); err != nil {
		return err
	}

	id := c.QueryInt("id", 0)

	if id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Inventory not found in query",
		})
	}

	inventory, err := controllers.GetCurrentInventoryByID(uint(id))
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
		"data":    inventory,
	})
}
