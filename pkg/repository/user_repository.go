package repository

import (
	"GoCrud/pkg/db"
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/helper"
	"errors"
	"log"
	"time"

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
	GetAll() (users []*entity.User, err error)
	CreateUser(user *entity.User) error
	UpdateUser(updateUserEntity *entity.User) error
	DeleteUser(delUserEntity *entity.User) error
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

func (repo *userRepository) GetAll() (users []*entity.User, err error) {
	err = repo.DB.Find(&users).Error

	if err != nil {
		return nil, err
	} else if len(users) == 0 {
		return nil, errors.New(helper.UserNotFound)
	}

	return users, err
}

func (repo *userRepository) CreateUser(user *entity.User) error {

	user.Password = hashAndSalt([]byte(user.Password))

	result := repo.DB.Create(&user)

	return result.Error
}

func (repo *userRepository) UpdateUser(updateUserEntity *entity.User) error {

	updateUserEntity.Password = hashAndSalt([]byte(updateUserEntity.Password))

	result := repo.DB.Save(&updateUserEntity)
	return result.Error
}

func (repo *userRepository) DeleteUser(delUserEntity *entity.User) error {

	delUserEntity.DeletedAt = &time.Time{}

	result := repo.DB.Save(&delUserEntity)
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
