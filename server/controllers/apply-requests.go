package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateApplyRequests(c *fiber.Ctx) error {
	var applyrequests models.ApplyRequest
	if err := c.BodyParser(&applyrequests); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBInstance.Db.Create(&applyrequests)
	return c.JSON(json{
		Data: applyrequests,
	})
}

func GetApplyRequests(c *fiber.Ctx) error {
	var applyrequests []models.ApplyRequest

	database.DBInstance.Db.Find(&applyrequests)
	return c.JSON(json{
		Data: applyrequests,
	})

}

func GetApplyRequest(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")

	}

	var applyrequest models.ApplyRequest
	result := database.DBInstance.Db.First(&applyrequest, "id = ?", id)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(json{
		Data: applyrequest,
	})

}
