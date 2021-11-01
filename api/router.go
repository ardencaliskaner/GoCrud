package api

import (
	"GoCrud/api/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	jwtMiddleware := middleware.NewJwtMiddleware()

	router := s.router

	v1 := router.Group("/v1/api")
	{
		auth := v1.Group("auth")
		{
			auth.POST("/login", s.authController.Login)
		}

		register := v1.Group("/users")
		{
			register.PUT("/", s.authController.Register)
		}

		user := v1.Group("/users")
		{
			user.Use(jwtMiddleware.AuthorizeJWT())
			user.PATCH("/:id", s.userController.UpdateUser)
			user.DELETE("/:id", s.userController.DeleteUser)
			user.GET("/:id", s.userController.GetById)
			user.GET("/", s.userController.GetAll)
		}
	}
	return router
}
