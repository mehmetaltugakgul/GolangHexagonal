package handler

import (
	"golangHexagonal/internal/app/model"
	"golangHexagonal/internal/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthActions interface {
	AuthenticateUser(email, password string) (*model.User, error)
}

type JWTActions interface {
	GenerateToken(userID uint) (string, error)
	VerifyToken(tokenString string) (*service.Claims, error)
}

type AuthHandler struct {
	userService AuthActions
	jwtService  JWTActions
}

func NewAuthHandler(userService AuthActions, jwtActions JWTActions) *AuthHandler {
	return &AuthHandler{userService: userService, jwtService: jwtActions}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/login", h.Login)
	app.Post("/logout", h.Logout)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input model.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.userService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logout"})
}
