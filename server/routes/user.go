package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router) {
	router.Get("/", controllers.GetUsers)
	router.Get("/:id",controllers.GetUser)
}
