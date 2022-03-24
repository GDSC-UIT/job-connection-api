package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateExperience(c *fiber.Ctx) error {
	var experience models.Experience
	if err := c.BodyParser(&experience); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.DBInstance.Db.Create(&experience)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(json{
		Data: experience,
	})
}

func GetExperiences(c *fiber.Ctx) error {
	user_id := c.Query("user_id")

	var experiences []models.Experience

	database.DBInstance.Db.Preload("Company").Find(&experiences, "user_id = ?", user_id)
	return c.JSON(json{
		Data: experiences,
	})

}

func GetExperience(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")

	}

	var experience models.Experience
	result := database.DBInstance.Db.First(&experience, "id = ?", id)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(json{
		Data: experience,
	})

}

func UpdateExperience(c *fiber.Ctx) error {
	var update_experience models.Experience
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")

	}

	if err := c.BodyParser(&update_experience); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var current_experience models.Experience

	result := database.DBInstance.Db.First(&current_experience, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_experience.UserID = update_experience.UserID
	current_experience.CompanyID = update_experience.CompanyID
	current_experience.JobTitle = update_experience.JobTitle
	current_experience.Description = update_experience.Description
	current_experience.SkillIds = update_experience.SkillIds
	current_experience.From = update_experience.From
	current_experience.To = update_experience.To

	result = database.DBInstance.Db.Save(&current_experience)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	return c.JSON(json{
		Data:    current_experience,
		Message: "Experience update successful",
	})
}

func DeleteExperience(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")
	}

	var experience models.Experience

	experience.ID = uint(id)
	if err := database.DBInstance.Db.Delete(&experience).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(json{
		Message: "Successfully Delete Experience",
	})

}
