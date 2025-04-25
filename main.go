package main

import (
	"app-backend/data/database"
	"app-backend/data/redis"
	"app-backend/domain/user"
	"app-backend/handler"
	"app-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app:=fiber.New()

	// Initialize Data clients
	dbConn := database.InitPostgres()
	redis.InitRedis("localhost:6379")

	// Construct User Domain
	userRepo    := user.NewRepository(dbConn)
	userService := user.NewService(userRepo)
	userHandler := &handler.UserHandler{Service: userService}

	// Register Routes
	routes.RegisterUserRoutes(app, userHandler)

	app.Listen(":8000")
}