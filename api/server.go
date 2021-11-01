package api

import (
	"GoCrud/api/controller"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router         *gin.Engine
	userController controller.UserController
	authController controller.AuthController
}

func NewServer() *Server {
	return &Server{
		router:         gin.Default(),
		userController: controller.NewUserController(),
		authController: controller.NewAuthController(),
	}
}

func (s *Server) Run() error {
	r := s.Routes()
	r.Use(cors.Default())

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
