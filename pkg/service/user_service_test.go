package service_test

import (
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/model"
	"GoCrud/pkg/repository/mocks"
	"GoCrud/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetById(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: "1324"}
		mockUserRepository.On("GetById", 1).Return(mockUser, nil)

		result, err := userService.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, "Arden", result.Name)
	})
}

func TestGetAll(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: "1324"}
		mockUserRepository.On("GetAll").Return([]*entity.User{mockUser}, nil)

		result, err := userService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, "Arden", result[0].Name)
	})
}

func TestCreateUser(t *testing.T) {

	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		password := "1324"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: string(hash)}
		registerModel := model.Register{Name: "Arden", Email: "arden@test.com", Password: string(hash)}
		email := mockUser.Email

		mockUserRepository.On("GetByEmail", email).Return(&entity.User{}, nil)

		mockUserRepository.On("CreateUser", mockUser).Return(nil)

		_, err := userService.CreateUser(registerModel)

		// mockUserRepository.AssertExpectations(t)

		assert.NoError(t, err)
	})
}

func TestUpdateUser(t *testing.T) {

	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		password := "1324"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: string(hash)}
		mockUserRepository.On("UpdateUser", mockUser).Return(nil)

		userModel := &model.User{Name: "Arden", Email: "arden@test.com", Password: string(hash)}
		result, err := userService.UpdateUser(1, userModel)

		assert.Equal(t, *userModel, result)
		assert.Nil(t, err)
	})
}

func TestValidate(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		user := model.User{Id: 1, Name: "Arden", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NoError(t, err)
	})
	t.Run("Error Empty User", func(t *testing.T) {

		err := userService.Validate(nil)

		assert.NotNil(t, err)

		assert.Equal(t, "user is empty", err.Error())
	})
	t.Run("Error Empty User Name", func(t *testing.T) {

		user := model.User{Id: 1, Name: "", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NotNil(t, err)

		assert.Equal(t, "user name is empty", err.Error())
	})
}
