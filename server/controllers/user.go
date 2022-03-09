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
func UpdateUser(c *fiber.Ctx) error {
	var update_user_info models.User
	id := c.Params("id")
	if err := c.BodyParser(&update_user_info); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var current_user_info models.User
	result := database.DBInstance.Db.First(&current_user_info, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_user_info.Name = update_user_info.Name
	current_user_info.Email = update_user_info.Email
	current_user_info.Photo = update_user_info.Photo
	current_user_info.Bio = update_user_info.Bio
	current_user_info.Location = update_user_info.Location
	current_user_info.FacebookURL = update_user_info.FacebookURL
	current_user_info.LinkedinURL = update_user_info.LinkedinURL
	current_user_info.GPA = update_user_info.GPA
	current_user_info.NumberOfYears = update_user_info.NumberOfYears
	current_user_info.Degree = update_user_info.Degree


	result = database.DBInstance.Db.Save(&current_user_info)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	
	return c.JSON(json{
		Data:    current_user_info,
		Message: "Informations update successful",
	})
}
