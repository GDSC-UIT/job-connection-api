package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	model := database.DBInstance.Db.Model(&models.User{})
	return c.JSON(pg.Response(model, c.Request(), &[]models.User{}))
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DBInstance.Db.First(&user, "id = ?", id)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(json{
		Data: user,
	})
}
