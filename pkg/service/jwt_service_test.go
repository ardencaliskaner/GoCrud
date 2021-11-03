package service_test

import (
	"GoCrud/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		jwtService := service.NewJWTServiceMock()

		result := jwtService.GenerateToken("1")

		assert.NotEmpty(t, result)
	})
	t.Run("Error_UserId_Null", func(t *testing.T) {
		jwtService := service.NewJWTServiceMock()

		result := jwtService.GenerateToken("")

		assert.Empty(t, result)
	})
}
