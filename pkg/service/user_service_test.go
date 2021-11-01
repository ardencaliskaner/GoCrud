package service_test

import (
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/model"
	"GoCrud/pkg/repository/mocks"
	"GoCrud/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestById(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		mockUser := &entity.User{Name: "Arden", Email: "arden@test.com", Password: "1324"}
		mockUserRepository.On("GetById").Return(mockUser, nil)

		result, err := userService.GetById(0)

		assert.NoError(t, err)
		assert.Equal(t, "Arden", result.Name)
	})
}

func TestValidate(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userService := service.NewUserServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		user := model.User{ID: 1, Name: "Arden", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NoError(t, err)
	})
	t.Run("Error Empty User", func(t *testing.T) {

		err := userService.Validate(nil)

		assert.NotNil(t, err)

		assert.Equal(t, "user is empty", err.Error())
	})
	t.Run("Error Empty User Name", func(t *testing.T) {

		user := model.User{ID: 1, Name: "", Email: "arden@mail.com"}
		err := userService.Validate(&user)

		assert.NotNil(t, err)

		assert.Equal(t, "user name is empty", err.Error())
	})
}
