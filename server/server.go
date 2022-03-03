package server

import (
	"fmt"

	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/middlewares"
	"github.com/GDSC-UIT/job-connection-api/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app *fiber.App
}

func (server *Server) settings() {
	middlewares.ConnectFirebase()
	database.ConnectDb()
}

func (server *Server) middlewaresInput() {
	server.app.Use(cors.New())
	server.app.Use(logger.New())
}

func (server *Server) routes() {
	api := server.app.Group("/api")

	api.Route("/profile", routes.ProfileRoute)
	api.Route("/users", routes.UserRoute)
	api.Route("/skills", routes.SkillsRoute)
	api.Route("/company", routes.CompanyRoute)
	api.Route("/jobs", routes.JobsRoute)
}

func (server *Server) middlewaresOutput() {
	server.app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	})
}

func New() *Server {
	server := &Server{
		app: fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				// Status code defaults to 500
				code := fiber.StatusInternalServerError

				// Retrieve the custom status code if it's an fiber.*Error
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}

				// Send custom error page
				err = c.Status(code).JSON(struct {
					Data    interface{} `json:"data"`
					Message string      `json:"message"`
				}{
					Message: err.Error(),
				})
				if err != nil {
					// In case the SendFile fails
					return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
				}

				// Return from handler
				return nil
			},
		}),
	}
	server.settings()
	server.middlewaresInput()
	server.routes()
	server.middlewaresOutput()

	return server
}

func (server *Server) ListenAndServe() {
	server.app.Listen(fmt.Sprintf(":%d", conf.Config.Port))
}
