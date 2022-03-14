package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateSkill(c *fiber.Ctx) error {
	var skill models.Skill
	if err := c.BodyParser(&skill); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBInstance.Db.Create(&skill)
	return c.JSON(json{
		Data:    skill,
		Message: "Successfully create skill",
	})
}

func GetSkills(c *fiber.Ctx) error {
	// var skills []models.Skill

	// database.DBInstance.Db.Find(&skills)
	// return c.JSON(json{
	// 	Data: skills,
	// })

	model := database.DBInstance.Db.Model(&models.Skill{})

	page := pg.With(model).Request(c.Request()).Response(&[]models.Skill{})
	// pg.Response(model, c.Request(), &[]models.Skill{})
	return c.JSON(page)
}

func GetSkill(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")
	}

	var skill models.Skill
	result := database.DBInstance.Db.First(&skill, "id = ?", id)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.JSON(json{
		Data: skill,
	})

}

func UpdateSkill(c *fiber.Ctx) error {
	var update_skill models.Skill
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")

	}

	if err := c.BodyParser(&update_skill); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var current_skill models.Skill

	result := database.DBInstance.Db.First(&current_skill, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_skill.Name = update_skill.Name

	result = database.DBInstance.Db.Save(&current_skill)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	return c.JSON(json{
		Data:    current_skill,
		Message: "Skill update successful",
	})
}

func DeleteSkill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that: id is an integer")
	}

	var skill models.Skill

	skill.ID = id
	if err := database.DBInstance.Db.Delete(&skill).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(json{
		Message: "Successfully Delete Skill",
	})

}
