package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func ExperienceRoute(router fiber.Router) {
	router.Post("/", controllers.CreateExperience)
	router.Get("/", controllers.GetExperiences)
	router.Get("/:id", controllers.GetExperience)
	router.Put("/:id", controllers.UpdateExperience)
	router.Delete("/:id", controllers.DeleteExperience)
}
