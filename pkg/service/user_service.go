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
	CreateUser(registerModel model.Register) (model.User, error)
	UpdateUser(id int, userModel *model.Update) (model.User, error)
	DeleteUser(id int) error
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

	user := model.User{Id: int(userEntity.ID), Name: userEntity.Name, Email: userEntity.Email, Password: userEntity.Password}

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
				Id:       int(userEntity.ID),
				Name:     userEntity.Name,
				Email:    userEntity.Email,
				Password: userEntity.Password,
			})
	}

	return users, nil
}

func (service *userService) CreateUser(registerModel model.Register) (model.User, error) {

	if registerModel.Email == "" || registerModel.Name == "" || registerModel.Password == "" {
		return model.User{}, errors.New(helper.BadRequest)
	}

	existingUser, _ := service.userRepository.GetByEmail(registerModel.Email)

	if existingUser.Email == registerModel.Email {
		return model.User{}, errors.New(helper.UserExist)
	}

	userEntity := entity.User{
		Name:     registerModel.Name,
		Email:    registerModel.Email,
		Password: registerModel.Password,
	}

	err := service.userRepository.CreateUser(&userEntity)
	if err != nil {
		return model.User{}, err
	}

	return model.User{Id: int(userEntity.ID), Name: userEntity.Name, Email: userEntity.Email, Password: registerModel.Password}, nil
}

func (service *userService) UpdateUser(id int, userModel *model.Update) (model.User, error) {

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
	} else if (entity.User{}) == *userEntity {
		return model.User{}, errors.New(helper.UserNotFound)
	}

	existingUser, _ := service.userRepository.GetByEmail(userModel.Email)

	if existingUser.Email == userModel.Email && existingUser.ID != uint(id) {
		return model.User{}, errors.New(helper.IdMatchedAnotherEmail)
	}

	userEntity.Name = userModel.Name
	userEntity.Email = userModel.Email
	userEntity.Password = userModel.Password

	errUpdate := service.userRepository.UpdateUser(userEntity)

	if errUpdate != nil {
		return model.User{}, errUpdate
	}

	return model.User{Id: int(userEntity.ID), Name: userEntity.Name, Email: userEntity.Email, Password: userModel.Password}, nil
}

func (service *userService) DeleteUser(id int) error {

	userEntity, err := service.userRepository.GetById(id)

	if err != nil {
		return err
	}

	errUpdate := service.userRepository.DeleteUser(userEntity)

	return errUpdate
}
