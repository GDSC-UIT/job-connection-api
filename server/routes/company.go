package routes

import (
	"github.com/GDSC-UIT/job-connection-api/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func CompanyRoute(router fiber.Router) {
	router.Post("/", controllers.CreateCompany)
	router.Get("/", controllers.GetCompanies)
	router.Get("/:id", controllers.GetCompanyByID)
	router.Put("/:id", controllers.UpdateCompany)
}
