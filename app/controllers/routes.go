package controllers

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (server *Server) SetupRoutes() {

	server.App.Use(recover.New())

	// Hooks
	webhook := server.App.Group("/webhooks")
	{
		webhook.Post("/line-bot", server.LineBotHooks)
	}

}
