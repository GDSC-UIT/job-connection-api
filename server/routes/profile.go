package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/GDSC-UIT/job-connection-api/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProfileRoute(router fiber.Router) {
	router.Get("/", middlewares.FirebaseAuthHandler(), controllers.GetProfile)
	router.Put("/",middlewares.FirebaseAuthHandler(),controllers.UpdateProfile)
}