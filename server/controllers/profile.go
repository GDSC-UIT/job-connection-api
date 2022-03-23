package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/middlewares"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Profile struct {
	Profile interface{} `json:"profile"`
	Type    string      `json:"type"`
}

func GetProfile(c *fiber.Ctx) error {
	info := c.Locals("info").(middlewares.UserInfo)

	accountType := new(models.AccountType)
	err := database.DBInstance.Db.First(&accountType, "user_id = ?", info.ID).Error
	if err != nil || accountType.Type == "user" {
		return getUserProfile(c, err == nil)
	}
	return getCompanyProfile(c)
}

func getUserProfile(c *fiber.Ctx, existed bool) error {
	info := c.Locals("info").(middlewares.UserInfo)

	var user_profile models.User
	result := database.DBInstance.Db.First(&user_profile, "id = ?", info.ID)
	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		if err := database.DBInstance.Db.Transaction(func(tx *gorm.DB) error {
			user_profile.ID = info.ID
			user_profile.Name = info.Name
			user_profile.Email = info.Email
			user_profile.Photo = info.Photo

			if err := tx.Create(&models.AccountType{UserID: info.ID, Type: "user"}).Error; err != nil {
				return err
			}
			result := tx.Create(&user_profile)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err)
		}

	}
	return c.JSON(Profile{
		Profile: user_profile,
		Type:    "user",
	})
}

func getCompanyProfile(c *fiber.Ctx) error {
	info := c.Locals("info").(middlewares.UserInfo)
	company := new(models.Company)
	database.DBInstance.Db.Where("id = ?", info.ID).First(company)
	return c.JSON(Profile{
		Profile: company,
		Type:    "company",
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	info := c.Locals("info").(middlewares.UserInfo)

	accountType := new(models.AccountType)
	err := database.DBInstance.Db.First(&accountType, "user_id = ?", info.ID).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if accountType.Type == "user" {
		var update_profile models.User

		if err := c.BodyParser(&update_profile); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		var current_profile models.User
		result := database.DBInstance.Db.First(&current_profile, "id = ?", info.ID)
		if result.RowsAffected != 0 && result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
		}

		// update
		current_profile.Name = update_profile.Name
		current_profile.Email = update_profile.Email
		current_profile.Phone = update_profile.Phone
		current_profile.Photo = update_profile.Photo
		current_profile.Bio = update_profile.Bio
		current_profile.Location = update_profile.Location
		current_profile.FacebookURL = update_profile.FacebookURL
		current_profile.LinkedinURL = update_profile.LinkedinURL
		current_profile.GPA = update_profile.GPA
		current_profile.NumberOfYears = update_profile.NumberOfYears
		current_profile.Degree = update_profile.Degree

		result = database.DBInstance.Db.Save(&current_profile)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
		}
		return c.JSON(json{
			Data:    current_profile,
			Message: "Profile update successful",
		})
	} else {
		var update_profile models.Company

		if err := c.BodyParser(&update_profile); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		var current_profile models.Company
		result := database.DBInstance.Db.First(&current_profile, "id = ?", info.ID)
		if result.RowsAffected != 0 && result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
		}

		// update
		current_profile.Name = update_profile.Name
		current_profile.Email = update_profile.Email
		current_profile.Photo = update_profile.Photo
		current_profile.Address = update_profile.Address
		current_profile.WorkingTime = update_profile.WorkingTime
		current_profile.Description = update_profile.Description

		result = database.DBInstance.Db.Save(&current_profile)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
		}
		return c.JSON(json{
			Data:    current_profile,
			Message: "Profile update successful",
		})
	}
}
