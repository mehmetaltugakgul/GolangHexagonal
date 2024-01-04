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

func TestHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return 200 when login success", func(t *testing.T) {
		reqBody := model.LoginInput{
			Email:    "test@gmail.com",
			Password: "123456",
		}

		mockAuthActions := mocks.NewMockAuthActions(ctrl)
		mokcJWTActions := mocks.NewMockJWTActions(ctrl)
		mockAuthActions.EXPECT().AuthenticateUser(reqBody.Email, reqBody.Password).Return(&model.User{
			ID:       1,
			Email:    reqBody.Email,
			Password: reqBody.Password,
			Name:     gomock.Any().String(),
		}, nil)
		mokcJWTActions.EXPECT().GenerateToken(gomock.Any()).Return("token", nil)

		body, err := json.Marshal(&reqBody)
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		authHandler := NewAuthHandler(mockAuthActions, mokcJWTActions)
		authHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("should return 400 when request body is invalid", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/login", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		authHandler := NewAuthHandler(nil, nil)
		authHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 400, resp.StatusCode)
	})

	t.Run("should return 401 when login failed", func(t *testing.T) {
		reqBody := model.LoginInput{
			Email:    "invalidemail@gmail.com",
			Password: "123456",
		}

		mockAuthActions := mocks.NewMockAuthActions(ctrl)
		mockAuthActions.EXPECT().AuthenticateUser(reqBody.Email, reqBody.Password).Return(nil, errors.New("invalid email or password"))

		body, err := json.Marshal(&reqBody)
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		authHandler := NewAuthHandler(mockAuthActions, nil)
		authHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 401, resp.StatusCode)
	})

	t.Run("should return 500 when generate token failed", func(t *testing.T) {
		reqBody := model.LoginInput{
			Email:    "test@gmail.com",
			Password: "123456",
		}

		mockAuthActions := mocks.NewMockAuthActions(ctrl)
		mockJWTActions := mocks.NewMockJWTActions(ctrl)
		mockAuthActions.EXPECT().AuthenticateUser(reqBody.Email, reqBody.Password).Return(&model.User{
			ID:       1,
			Email:    reqBody.Email,
			Password: reqBody.Password,
			Name:     gomock.Any().String(),
		}, nil)
		mockJWTActions.EXPECT().GenerateToken(gomock.Any()).Return("", errors.New("failed to generate token"))

		body, err := json.Marshal(&reqBody)
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		app := fiber.New()

		authHandler := NewAuthHandler(mockAuthActions, mockJWTActions)
		authHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 500, resp.StatusCode)
	})
}

func TestHandler_Logout(t *testing.T) {
	t.Run("should return 200 when logout success", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/logout", nil)

		app := fiber.New()

		authHandler := NewAuthHandler(nil, nil)
		authHandler.RegisterRoutes(app)

		resp, err := app.Test(req)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)
	})
}
