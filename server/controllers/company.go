package controllers

import (

	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCompany(context *fiber.Ctx) error {
	p := new(models.Company)
	if err := context.BodyParser(p); err != nil {
		return err
	}
	database.DBInstance.Db.Create(p)
	return context.JSON(json{Data: p})
}
func GetCompanies(context *fiber.Ctx) error{
	model := database.DBInstance.Db.Model(&models.Company{})
	
	page := pg.Response(model, context.Request(), &[]models.Company{})
	return context.JSON(page)
	
}
func GetCompanyByID(context *fiber.Ctx) error{
	id := context.Params("id")
	company := new(models.Company)
	database.DBInstance.Db.Where("id = ?", id).First(company)
	return context.JSON(json{Data: company})
}
func UpdateCompany(context *fiber.Ctx) error{
	var update_info models.Company
	id := context.Params("id")
	if err := context.BodyParser(&update_info); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var current_info models.Company
	result := database.DBInstance.Db.First(&current_info, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_info.Name = update_info.Name
	current_info.Email = update_info.Email
	current_info.Photo = update_info.Photo
	current_info.Address = update_info.Address
	current_info.WorkingTime = update_info.WorkingTime
	current_info.Description = update_info.Description

	result = database.DBInstance.Db.Save(&current_info)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	
	return context.JSON(json{
		Data:    current_info,
		Message: "Informations update successful",
	})
}