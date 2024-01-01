package router

import (
	"github.com/gofiber/fiber/v2"
	"golangHexagonal/internal/app/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Post("/login", authHandler.Login)
	app.Get("/logout", authHandler.Logout)
	app.Get("/users", userHandler.GetUsers)
}
