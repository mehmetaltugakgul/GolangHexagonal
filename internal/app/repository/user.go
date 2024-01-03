package repository

import (
	"golangHexagonal/internal/app/model"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateUser(user *model.User) error
	FindUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	FindUserByEmail(email string) (*model.User, error)
	FindUsers() ([]*model.User, error)
}

type UserRepository struct {
	db          *gorm.DB
	UserActions IRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindUserByID(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindUsers() ([]*model.User, error) {
	var users []*model.User
	result := r.db.Find(&users)
	return users, result.Error
}
