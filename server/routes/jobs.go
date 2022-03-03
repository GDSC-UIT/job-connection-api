package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func JobsRoute(router fiber.Router) {
	router.Post("/", controllers.CreateJob)
	router.Get("/", controllers.GetJobs)
	router.Get("/:id", controllers.GetJobByID)
	router.Put("/:id", controllers.UpdateJob)
}
