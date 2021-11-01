package controller

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController() *userController {
	return &userController{
		userService: service.NewUserService(),
		jwtService:  service.NewJWTService(),
	}
}

func (controller *userController) GetById(ctx *gin.Context) {
	response := helper.BuildErrorResponse(helper.ServerError, helper.ServerError, helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

func (controller *userController) GetAll(ctx *gin.Context) {
	response := helper.BuildErrorResponse(helper.ServerError, helper.ServerError, helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

func (controller *userController) UpdateUser(ctx *gin.Context) {
	response := helper.BuildErrorResponse(helper.ServerError, helper.ServerError, helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}

func (controller *userController) DeleteUser(ctx *gin.Context) {
	response := helper.BuildErrorResponse(helper.ServerError, helper.ServerError, helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
