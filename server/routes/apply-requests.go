package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func ApplyRequestsRoute(router fiber.Router) {
	router.Post("/", controllers.CreateApplyRequests)
	router.Get("/", controllers.GetApplyRequests)
	router.Get("/:id", controllers.GetApplyRequest)
}
