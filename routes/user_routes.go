package routes

import (
	"app-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app fiber.Router, h *handler.UserHandler) {
	userGroup := app.Group("/users")
	userGroup.Get("/register", h.RegisterStudent)
	// userGroup.Get("/login", h.)
}
