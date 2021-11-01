package service_test

import (
	"GoCrud/pkg/db/entity"
	"GoCrud/pkg/repository/mocks"
	"GoCrud/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestVerifyCredential(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	authService := service.NewAuthServiceMock(mockUserRepository)

	t.Run("Success", func(t *testing.T) {

		password := "1324"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		mockUserResp := &entity.User{Name: "Arden", Email: "arden@test.com", Password: string(hash)}
		email := mockUserResp.Email

		mockUserRepository.On("GetByEmail", email).Return(mockUserResp, nil)

		_, err := authService.VerifyCredential(email, password)

		mockUserRepository.AssertExpectations(t)

		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {

		password := "1324444"

		mockUserResp := &entity.User{Name: "Arden", Email: "arden@test.com", Password: password}
		email := mockUserResp.Email

		mockUserRepository.On("GetByEmail", email).Return(mockUserResp, nil)

		_, err := authService.VerifyCredential(email, password)

		assert.Error(t, err)
	})
}
