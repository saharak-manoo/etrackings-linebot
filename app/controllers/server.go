package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/saharak/etrackings-linebot/app/utils/response"
)

type Server struct {
	App *fiber.App
}

func (server *Server) Initialize() {
	server.App = fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println("[ETrackings GO] error -> ", err)

			return response.BadRequest(c)
		},
	})
	server.App.Use(logger.New())
}

func (server *Server) Run(addr string) {
	server.SetupRoutes()

	log.Fatal(server.App.Listen(addr))
}
