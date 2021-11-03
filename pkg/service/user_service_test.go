package service_test

import (
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/repository/mocks"
	"GoCrud/pkg/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)
		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: "1324"}
		mockUserRepository.On("GetById", 1).Return(mockUser, nil)

		result, err := userService.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, "Arden", result.Name)
	})
	t.Run("Error_User_Not_Found", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)
		mockUserRepository.On("GetById", 1).Return(&entity.User{}, errors.New(helper.UserNotFound))

		_, err := userService.GetById(1)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserNotFound)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: "1324"}
		mockUserRepository.On("GetAll").Return([]*entity.User{mockUser}, nil)

		result, err := userService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, "Arden", result[0].Name)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		registerModel := model.Register{Name: "Arden", Email: "arden@test.com", Password: "1234"}
		mockUser := entity.User{Name: "Arden", Email: "arden@test.com", Password: "1234"}

		mockUserRepository.On("GetByEmail", registerModel.Email).Return(&entity.User{}, nil)

		mockUserRepository.On("CreateUser", &mockUser).Return(nil)

		createdUser, err := userService.CreateUser(registerModel)

		assert.Equal(t, createdUser.Email, mockUser.Email)

		assert.NoError(t, err)
	})
	t.Run("Error_User_Already_Exist", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		mockUser := entity.User{Name: "Arden", Email: "arden@test.com", Password: "1234"}
		registerModel := model.Register{Name: "Arden", Email: "arden@test.com", Password: "1234"}

		mockUserRepository.On("GetByEmail", registerModel.Email).Return(&mockUser, nil)

		mockUserRepository.On("CreateUser", &mockUser).Return(mock.Anything)

		_, err := userService.CreateUser(registerModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserExist)
	})
	t.Run("Error_Email_Null", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		registerModel := model.Register{Name: "Arden", Email: "", Password: "1234"}

		// mockUserRepository.On("GetByEmail", registerModel.Email).Return(mock.Anything)
		// mockUserRepository.On("CreateUser", mock.Anything).Return(mock.Anything)
		_, err := userService.CreateUser(registerModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.BadRequest)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "Arden", Email: "arden@test.com", Password: "1234"}
		mockUser := entity.User{Name: "Arden", Email: "arden@test.com", Password: "1234"}

		mockUserRepository.On("GetById", 0).Return(&mockUser, nil)
		mockUserRepository.On("GetByEmail", userModel.Email).Return(&mockUser, nil)

		mockUserRepository.On("UpdateUser", &mockUser).Return(nil)

		updatedUser, err := userService.UpdateUser(0, &userModel)

		assert.Equal(t, updatedUser.Email, mockUser.Email)

		assert.NoError(t, err)
	})
	t.Run("Error_User_Not_Found_With_Given_Id", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "Arden", Email: "arden@test.com", Password: "1234"}
		// mockUser := entity.User{Name: "Arden", Email: "arden@test.com", Password: "1234"}

		mockUserRepository.On("GetById", 1).Return(&entity.User{}, nil)

		// mockUserRepository.On("GetByEmail", userModel.Email).Return(&entity.User{}, nil)
		// mockUserRepository.On("UpdateUser", &mockUser).Return(nil)

		_, err := userService.UpdateUser(1, &userModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserNotFound)
	})
	t.Run("Error_User_Not_Found_With_Given_Id", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "Arden", Email: "arden@test.com", Password: "1234"}
		mockUser := entity.User{Name: "Arden", Email: "arden@test.com", Password: "1234"}

		mockUserRepository.On("GetById", 1).Return(&mockUser, nil)
		mockUserRepository.On("GetByEmail", userModel.Email).Return(&mockUser, nil)

		mockUserRepository.On("UpdateUser", &mockUser).Return(nil)

		_, err := userService.UpdateUser(1, &userModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.IdMatchedAnotherEmail)
	})
	t.Run("Error_Email_Null", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "Arden", Email: "", Password: "1234"}
		_, err := userService.UpdateUser(0, &userModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserEmailEmpty)
	})
	t.Run("Error_Name_Null", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "", Email: "arden@mail.com", Password: "1234"}
		_, err := userService.UpdateUser(0, &userModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserNameEmpty)
	})
	t.Run("Error_Password_Null", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		userModel := model.Update{Name: "Arden", Email: "arden@mail.com", Password: ""}
		_, err := userService.UpdateUser(0, &userModel)

		assert.NotNil(t, err)
		assert.EqualError(t, err, helper.UserPassEmpty)
	})
}

func TestValidate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		user := model.User{Id: 1, Name: "Arden", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NoError(t, err)
	})
	t.Run("Error_Empty_User", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)
		err := userService.Validate(nil)

		assert.NotNil(t, err)

		assert.Equal(t, "user is empty", err.Error())
	})
	t.Run("Error_Empty_User_Name", func(t *testing.T) {
		mockUserRepository := new(mocks.UserRepository)
		userService := service.NewUserServiceMock(mockUserRepository)

		user := model.User{Id: 1, Name: "", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NotNil(t, err)

		assert.Equal(t, "user name is empty", err.Error())
	})
}
