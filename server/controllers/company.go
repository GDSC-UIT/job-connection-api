package controllers

import (
	"context"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"
	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func CreateCompany(ctx *fiber.Ctx) error {
	p := new(models.Company)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	opt := option.WithCredentialsJSON([]byte(conf.Config.FirebaseServiceAccount))
	fireApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err)
	}

	client, err := fireApp.Auth(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err)
	}
	_, err = client.GetUserByEmail(context.Background(), p.Email)
	if err == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Email has already been registered")
	}
	userToCreate := new(auth.UserToCreate)
	userToCreate.Email(p.Email)
	userToCreate.Password("123456")
	userCreated, err := client.CreateUser(context.Background(), userToCreate)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err)
	}

	if err = database.DBInstance.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models.AccountType{UserID: userCreated.UID, Type: "company"}).Error; err != nil {
			return err
		}
		p.ID = userCreated.UID

		if p.Photo == "" {
			p.Photo = "https://i.imgur.com/FuoDVfD.png"
		}

		if err := tx.Create(p).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(json{Data: p, Message: "Your password is 123456"})
}
func GetCompanies(context *fiber.Ctx) error {
	model := database.DBInstance.Db.Model(&models.Company{})

	page := pg.Response(model, context.Request(), &[]models.Company{})
	return context.JSON(page)

}
func GetCompanyByID(context *fiber.Ctx) error {
	id := context.Params("id")
	company := new(models.Company)
	database.DBInstance.Db.Where("id = ?", id).First(company)
	return context.JSON(json{Data: company})
}
func UpdateCompany(context *fiber.Ctx) error {
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
	current_info.Approved = update_info.Approved
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
