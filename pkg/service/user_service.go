package service

import (
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/repository"
	"errors"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	GetById(id int) (model.User, error)
	Validate(post *model.User) error
}

func NewUserService() UserService {
	return &userService{userRepository: repository.NewUserRepository()}
}

func NewUserServiceMock(repo repository.UserRepository) UserService {
	return &userService{userRepository: repo}
}

func (service *userService) Validate(user *model.User) error {
	if user == nil {
		err := errors.New(helper.UserEmpty)
		return err
	}
	if user.Name == "" {
		err := errors.New(helper.UserNameEmpty)
		return err
	}

	return nil
}

func (service *userService) GetById(id int) (model.User, error) {

	userEntity, err := service.userRepository.GetById(id)

	if err != nil {
		return model.User{}, err
	}

	if (entity.User{}) == *userEntity {
		return model.User{}, errors.New(helper.UserNotFound)
	}

	user := model.User{Name: userEntity.Name, Email: userEntity.Email, Password: userEntity.Password}

	return user, nil
}
