package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateJob(c *fiber.Ctx) error {
	p := new(models.Job)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	database.DBInstance.Db.Create(p)
	return c.JSON(json{Data: p})
}
func GetJobs(c *fiber.Ctx) error{
	model := database.DBInstance.Db.Model(&models.Job{})
	
	page := pg.Response(model, c.Request(), &[]models.Job{})
	return c.JSON(page)
	
}
func GetJobByID(c *fiber.Ctx) error{
	id := c.Params("id")
	job := new(models.Job)
	database.DBInstance.Db.Where("id = ?", id).First(job)
	return c.JSON(json{Data: job})
}
func UpdateJob(c *fiber.Ctx) error{
	var update_job_info models.Job
	id := c.Params("id")
	if err := c.BodyParser(&update_job_info); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var current_job_info models.Job
	result := database.DBInstance.Db.First(&current_job_info, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_job_info.Title = update_job_info.Title
	current_job_info.Address = update_job_info.Address
	current_job_info.Description = update_job_info.Description
	current_job_info.SkillIds = update_job_info.SkillIds

	result = database.DBInstance.Db.Save(&current_job_info)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	
	return c.JSON(json{
		Data:    current_job_info,
		Message: "Informations update successful",
	})
}
