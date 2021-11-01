package repository

import (
	"GoCrud/pkg/db"
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/helper"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	DB *gorm.DB
}

//go:generate mockery --name=UserRepository
type UserRepository interface {
	GetById(id int) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

func NewUserRepository() UserRepository {
	return &userRepository{DB: db.ConnectDB()}
}

func (repo *userRepository) GetById(id int) (*entity.User, error) {
	var user entity.User
	res := repo.DB.Where("id = ?", id).Take(&user)
	if res.Error != nil {
		return &entity.User{}, errors.New(helper.UserNotFound)
	}
	return &user, nil
}

func (repo *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	res := repo.DB.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return &entity.User{}, errors.New(helper.UserNotFound)
	}
	return &user, nil
}

func (repo *userRepository) CreateUser(user *entity.User) error {
	user.Password = hashAndSalt([]byte(user.Password))

	result := repo.DB.Create(&user)

	return result.Error
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
