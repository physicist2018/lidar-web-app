package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/physicist2018/lidar-web-app/platform/database"
)

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	db, err := database.OpenDbConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user, err := db.GetUser(int(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given id not found",
			"user":  nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
