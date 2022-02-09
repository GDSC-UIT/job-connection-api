package server

import (
	"fmt"
	"net/http"

	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app *fiber.App
}

func (server *Server) settings() {
	database.ConnectDb()
}

func (server *Server) middlewaresInput() {
	server.app.Use(cors.New())
	server.app.Use(logger.New())
}

func (server *Server) routes() {
	api := server.app.Group("/api")

	api.Route("/user-profiles",routes.UserProfileRoute)
}

func (server *Server) middlewaresOutput(){
	server.app.Use(func(c *fiber.Ctx) error{
		return c.SendStatus(http.StatusNotFound)
	})
}

func New() *Server {
	server := &Server {
		app: fiber.New(),
	}
	server.settings()
	server.middlewaresInput()
	server.routes()
	server.middlewaresOutput()

	return server
}

func (server *Server) ListenAndServe() {
	server.app.Listen(fmt.Sprintf(":%d",conf.Config.Port))
}
