package routes

import "github.com/gofiber/fiber/v2"

func UserProfileRoute(router fiber.Router) {
	router.Get("/",func (c *fiber.Ctx) error{
		return c.SendString("get user profiles")
	})
}