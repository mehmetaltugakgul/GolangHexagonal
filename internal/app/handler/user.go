package handler

import (
	"golangHexagonal/internal/app/model"

	"github.com/gofiber/fiber/v2"
)

type UserActions interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	GetUsers() ([]*model.User, error)
}

type UserHandler struct {
	service UserActions
}

func NewUserHandler(service UserActions) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUser)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	app.Get("/users", h.GetUsers)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user.ID = uint(id)

	if err := h.service.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)

}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users, "total": len(users), "message": "success", "status": 200, "success": true})
}
