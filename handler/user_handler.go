package handler

import (
	user "app-backend/domain/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *user.Service
}

type RegisterRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Grade       string `json:"grade"`
	Location    string `json:"location"`
}

func (h *UserHandler) RegisterStudent(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	u := &user.User{
		Name:  req.Name,
		Email: req.Email,
	}

	s := &user.Student{
		DateOfBirth: req.DateOfBirth,
		Grade:       req.Grade,
		Location:    req.Location,
	}

	if err := h.Service.RegisterStudent(u, s, req.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "registered"})
}
