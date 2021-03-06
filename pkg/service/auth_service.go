package service

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/repository"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) (model.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService() AuthService {
	return &authService{userRepository: repository.NewUserRepository()}
}

func NewAuthServiceMock(repo repository.UserRepository) AuthService {
	return &authService{userRepository: repo}
}

func (service *authService) VerifyCredential(email string, password string) (model.User, error) {
	userEntity, verifyErr := service.userRepository.GetByEmail(email)

	if verifyErr == nil {
		comparedPassword := comparePassword(userEntity.Password, []byte(password))

		if userEntity.Email == email && comparedPassword {
			return model.User{Id: int(userEntity.ID), Name: userEntity.Name, Email: userEntity.Email}, nil
		}
		return model.User{}, errors.New(helper.EmailOrPassNotMatch)
	}

	return model.User{}, errors.New(helper.UserNotFound)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
