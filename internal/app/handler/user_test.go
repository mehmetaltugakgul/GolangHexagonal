package handler

import (
	"bytes"
	"encoding/json"
	"errors"

	"golangHexagonal/internal/app/handler/mocks"
	"golangHexagonal/internal/app/model"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 201 when create user success", func(t *testing.T) {
		reqBody := model.User{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "123456",
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err)

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().CreateUser(&reqBody).Return(nil)

		app := fiber.New()

		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("should return 400 when request body is invalid", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		userHandler := NewUserHandler(nil)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 500 when service return error", func(t *testing.T) {
		reqBody := model.User{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "123456",
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err)

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().CreateUser(&reqBody).Return(errors.New("error"))

		app := fiber.New()

		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandler_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 200 when get user success", func(t *testing.T) {
		expectedRespBody := model.User{
			ID:       1,
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "123456",
		}

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().GetUserByID(uint(1)).Return(&expectedRespBody, nil)

		app := fiber.New()

		req := httptest.NewRequest("GET", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)

		var respBody model.User
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.Nil(t, err)

		assert.Equal(t, expectedRespBody, respBody)
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)

		app := fiber.New()

		req := httptest.NewRequest("GET", "/users/invalid", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 500 when service return error", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().GetUserByID(uint(1)).Return(nil, errors.New("error"))

		app := fiber.New()

		req := httptest.NewRequest("GET", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandler_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 200 when update user success", func(t *testing.T) {
		reqBody := model.User{
			ID:       1,
			Name:     "update",
			Email:    "update@gmail.com",
			Password: "123456",
		}

		expectedRespBody := model.User{
			ID:       1,
			Name:     "update",
			Email:    "update@gmail.com",
			Password: "123456",
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err)

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().UpdateUser(&reqBody).Return(nil)

		app := fiber.New()

		req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)

		var respBody model.User
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		assert.Nil(t, err)

		assert.Equal(t, expectedRespBody, respBody)
	})

	t.Run("should return 400 when request body is invalid", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		userHandler := NewUserHandler(nil)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)

		app := fiber.New()

		req := httptest.NewRequest("PUT", "/users/invalid", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 500 when service return error", func(t *testing.T) {
		reqBody := model.User{
			ID:       1,
			Name:     "update",
			Email:    "update@gmail.com",
			Password: "123456",
		}

		body, err := json.Marshal(reqBody)
		assert.Nil(t, err)

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().UpdateUser(&reqBody).Return(errors.New("error"))

		app := fiber.New()

		req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandler_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 200 when delete user success", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().DeleteUser(uint(1)).Return(nil)

		app := fiber.New()

		req := httptest.NewRequest("DELETE", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 204, resp.StatusCode)
	})

	t.Run("should return 400 when id is invalid", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)

		app := fiber.New()

		req := httptest.NewRequest("DELETE", "/users/invalid", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 500 when service return error", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().DeleteUser(uint(1)).Return(errors.New("error"))

		app := fiber.New()

		req := httptest.NewRequest("DELETE", "/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandler_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 200 when get users success", func(t *testing.T) {
		users := []*model.User{
			{
				ID:       1,
				Name:     "test",
				Email:    "test@gmail.com",
				Password: "123456",
			},
			{
				ID:       2,
				Name:     "test2",
				Email:    "test2@gmail.com",
				Password: "123456",
			},
		}

		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().GetUsers().Return(users, nil)

		app := fiber.New()

		req := httptest.NewRequest("GET", "/users", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("should return 500 when service return error", func(t *testing.T) {
		mockUserService := mocks.NewMockUserActions(ctrl)
		mockUserService.EXPECT().GetUsers().Return(nil, errors.New("error"))

		app := fiber.New()

		req := httptest.NewRequest("GET", "/users", nil)
		req.Header.Set("Content-Type", "application/json")

		userHandler := NewUserHandler(mockUserService)
		userHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 500, resp.StatusCode)
	})
}
