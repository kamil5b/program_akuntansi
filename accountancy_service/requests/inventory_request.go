package requests

import (
	"log/slog"
	"program_akuntansi/accountancy_service/controllers"

	"github.com/gofiber/fiber/v2"
)

//CREATE

func InventoryOpenItem(c *fiber.Ctx) error { //GET
	/*
		Authorization Header
		open_item	: uint

	*/

	if err := AuthUser(c, "AUTH_INVENTORY_OPEN"); err != nil {
		c.Status(403)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)
		slog.Error("id not valid")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	unit, err := c.ParamsInt("unit", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if unit == 0 {
		c.Status(200)
		slog.Error("no inventory opened")
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "no inventory opened",
		})
	}

	inv_id, err := controllers.InventoryOpenItem(uint(id), uint(unit))
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    inv_id,
	})
}

//GET

func GetInventoryByID(c *fiber.Ctx) error { //GET

	/*
		Authorization Header

	*/

	if err := AuthUser(c, "AUTH_GET_INVENTORY_ID"); err != nil {
		c.Status(403)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)
		slog.Error("id not valid")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	inventory, err := controllers.GetInventoryByID(uint(id))
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
		c.Status(403)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	inventories, err := controllers.GetAllInventories()
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
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
		c.Status(403)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  403,
			"message": err.Error(),
		})
	}

	id, err := c.ParamsInt("id", 0)
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	if id == 0 {
		c.Status(400)
		slog.Error("id not valid")
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "id not valid",
		})
	}

	inventory, err := controllers.GetCurrentInventoryByID(uint(id))
	if err != nil {
		c.Status(400)
		slog.Error(err.Error())
		return c.JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    inventory,
	})
}
