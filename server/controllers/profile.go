package controllers

import (
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/middlewares"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
)

type Profile struct {
	Profile interface{} `json:"profile"`
	Type    string      `json:"type"`
}

func GetProfile(c *fiber.Ctx) error {
	info := c.Locals("info").(middlewares.UserInfo)
	profile_type := c.Locals("type").(string)

	if profile_type == "user" {
		var user_profile models.User
		result := database.DBInstance.Db.First(&user_profile, "id = ?", info.ID)
		if result.RowsAffected != 0 && result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
		}
		if result.RowsAffected == 0 {
			user_profile.ID = info.ID
			user_profile.Name = info.Name
			user_profile.Email = info.Email
			user_profile.Photo = info.Photo

			result := database.DBInstance.Db.Create(&user_profile)
			if result.Error != nil {
				return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
			}
		}
		return c.JSON(Profile{
			Profile: user_profile,
			Type:    profile_type,
		})
	} else {
		return c.JSON(Profile{
			Type: profile_type,
		})
	}
}

func UpdateProfile(c *fiber.Ctx) error {
	info := c.Locals("info").(middlewares.UserInfo)
	profile_type := c.Locals("type").(string)

	if profile_type == "user" {
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
		return c.JSON(Profile{
			Type: profile_type,
		})
	}
}
