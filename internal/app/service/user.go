package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"golangHexagonal/internal/app/model"
	"golangHexagonal/internal/app/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *UserService) CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)

}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)

}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) AuthenticateUser(email, password string) (*model.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetUsers() ([]*model.User, error) {
	return s.repo.FindUsers()
}
