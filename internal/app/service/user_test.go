package service

import (
	"errors"
	"golangHexagonal/internal/app/model"
	"golangHexagonal/internal/app/service/mocks"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUserByID(uint(1)).Return(&model.User{
			Email:    "test@gmail.com",
			ID:       1,
			Name:     "test",
			Password: "123456",
		}, nil)

		srv := NewUserService(mockRepo)
		user, err := srv.GetUserByID(1)
		assert.Nil(t, err)

		assert.Equal(t, user.ID, uint(1))
	})

	t.Run("should return error when get user fail", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUserByID(uint(1)).Return(nil, errors.New("error"))

		srv := NewUserService(mockRepo)
		user, err := srv.GetUserByID(1)
		assert.NotNil(t, err)

		assert.Nil(t, user)
	})
}

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should create user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

		srv := NewUserService(mockRepo)
		err := srv.CreateUser(&model.User{
			Email:    "test@gmail.com",
			ID:       1,
			Name:     "test",
			Password: "123456",
		})
		assert.Nil(t, err)
	})

	t.Run("should return error when create user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().CreateUser(gomock.Any()).Return(errors.New("error"))

		srv := NewUserService(mockRepo)
		err := srv.CreateUser(&model.User{
			Email:    "test@gmail.com",
			ID:       1,
			Name:     "test",
			Password: "123456",
		})
		assert.NotNil(t, err)
	})
}

func TestService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should update user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().UpdateUser(gomock.Any()).Return(nil)

		srv := NewUserService(mockRepo)
		err := srv.UpdateUser(&model.User{
			Email:    "update@gmail.com",
			ID:       1,
			Name:     "update",
			Password: "123456",
		})
		assert.Nil(t, err)
	})

	t.Run("should return error when update user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("error"))

		srv := NewUserService(mockRepo)
		err := srv.UpdateUser(&model.User{
			Email:    "update@gmail.com",
			ID:       1,
			Name:     "update",
			Password: "123456",
		})
		assert.NotNil(t, err)
	})
}

func TestService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should delete user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().DeleteUser(uint(1)).Return(nil)

		srv := NewUserService(mockRepo)
		err := srv.DeleteUser(1)
		assert.Nil(t, err)
	})

	t.Run("should return error when delete user", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().DeleteUser(uint(1)).Return(errors.New("error"))

		srv := NewUserService(mockRepo)
		err := srv.DeleteUser(1)
		assert.NotNil(t, err)
	})
}

func TestService_AuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return user", func(t *testing.T) {
		email := "test@gmail.com"
		password := "123456"
		hashedPassword := "$2b$12$gBaSY3Lr5QHhLx/yiq/oeu4xw7lUuMNracbFbIJUAFYaY4/Ic0xb."
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUserByEmail(email).Return(&model.User{
			Email:    "test@gmail.com",
			ID:       uint(1),
			Name:     "test",
			Password: hashedPassword,
		}, nil)

		srv := NewUserService(mockRepo)
		user, err := srv.AuthenticateUser(email, password)
		assert.Nil(t, err)

		assert.Equal(t, user.ID, uint(1))
	})

	t.Run("should return error when authenticate user fail", func(t *testing.T) {
		email := "test@gmail.com"
		password := "wrong password"
		hashedPassword := "$2b$12$gBaSY3Lr5QHhLx/yiq/oeu4xw7lUuMNracbFbIJUAFYaY4/Ic0xb."
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUserByEmail(email).Return(&model.User{
			ID:       uint(1),
			Email:    "test@gmail.com",
			Name:     "test",
			Password: hashedPassword,
		}, nil)

		srv := NewUserService(mockRepo)
		user, err := srv.AuthenticateUser(email, password)
		assert.NotNil(t, err)

		assert.Nil(t, user)
	})

	t.Run("should return error when find user fail", func(t *testing.T) {
		email := "notfound@gmail.com"
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUserByEmail(email).Return(nil, errors.New("error"))

		srv := NewUserService(mockRepo)
		user, err := srv.AuthenticateUser(email, "123456")
		assert.NotNil(t, err)

		assert.Nil(t, user)
	})
}

func TestService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return users", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUsers().Return([]*model.User{
			{
				Email:    "test@gmail.com",
				ID:       1,
				Name:     "test",
				Password: "123456",
			},
			{
				Email:    "test2@gmail.com",
				ID:       2,
				Name:     "test2",
				Password: "123456",
			},
		}, nil)

		srv := NewUserService(mockRepo)
		users, err := srv.GetUsers()
		assert.Nil(t, err)

		assert.Equal(t, len(users), 2)
	})

	t.Run("should return error when get users fail", func(t *testing.T) {
		mockRepo := mocks.NewMockIRepository(ctrl)
		mockRepo.EXPECT().FindUsers().Return(nil, errors.New("error"))

		srv := NewUserService(mockRepo)
		users, err := srv.GetUsers()
		assert.NotNil(t, err)

		assert.Nil(t, users)
	})
}
