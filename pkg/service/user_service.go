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
	GetAll() ([]model.User, error)
	Validate(post *model.User) error
	UpdateUser(id int, userModel *model.User) (model.User, error)
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

func (service *userService) GetAll() ([]model.User, error) {
	var users []model.User

	userEntites, err := service.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	for _, userEntity := range userEntites {
		users = append(
			users,
			model.User{
				Name:  userEntity.Name,
				Email: userEntity.Email,
			})
	}

	return users, nil
}

func (service *userService) UpdateUser(id int, userModel *model.User) (model.User, error) {

	if userModel.Email == "" {
		return model.User{}, errors.New(helper.UserEmailEmpty)
	}
	if userModel.Name == "" {
		return model.User{}, errors.New(helper.UserNameEmpty)
	}
	if userModel.Password == "" {
		return model.User{}, errors.New(helper.UserPassEmpty)
	}

	userEntity, err := service.userRepository.GetById(id)

	if err != nil {
		return model.User{}, err
	}

	if (entity.User{}) == *userEntity {
		return model.User{}, errors.New(helper.UserNotFound)
	}

	userEntity.Name = userModel.Name
	userEntity.Email = userModel.Email
	userEntity.Password = userModel.Password

	errUpdate := service.userRepository.UpdateUser(userEntity)

	return *userModel, errUpdate
}
