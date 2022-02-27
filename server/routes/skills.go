package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func SkillsRoute(router fiber.Router) {
	router.Post("/", controllers.CreateSkill)
	router.Get("/", controllers.GetSkills)
	router.Get("/:id", controllers.GetSkill)
	router.Put("/:id", controllers.UpdateSkill)
	router.Delete("/:id", controllers.DeleteSkill)
}
