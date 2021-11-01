package middleware

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtMiddleware struct {
	jwtService service.JWTService
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{
		jwtService: service.NewJWTService(),
	}
}

func (mdw *JwtMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse(helper.TokenNotFound, helper.TokenNotFound, nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := mdw.jwtService.ValidateToken(authHeader)

		if err != nil {
			response := helper.BuildErrorResponse(helper.ServerError, err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			response := helper.BuildErrorResponse(helper.TokenNotValid, err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
