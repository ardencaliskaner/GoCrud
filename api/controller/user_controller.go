package controller

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/service"
	"strconv"

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
	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	id, errParse := strconv.Atoi(ctx.Param("id"))

	if errParse != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errParse.Error(), helper.EmptyObj{})
		ctx.JSON(response.Code, response)
		return
	}

	user, err := controller.userService.GetById(id)

	if err != nil {
		response := helper.BuildErrorResponseWithoutMessage(err.Error(), helper.EmptyObj{})
		ctx.JSON(response.Code, response)
		return
	}

	response := helper.BuildResponse(true, helper.Success, user)
	ctx.JSON(response.Code, response)
}

func (controller *userController) GetAll(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	users, err := controller.userService.GetAll()
	if err != nil {
		response := helper.BuildErrorResponseWithoutMessage(err.Error(), helper.EmptyObj{})
		ctx.JSON(response.Code, response)
		return
	}

	response := helper.BuildResponse(true, helper.Success, users)
	ctx.JSON(response.Code, response)
}

func (controller *userController) UpdateUser(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}
	var userModel model.User

	id, errParse := strconv.Atoi(ctx.Param("id"))
	errModel := ctx.ShouldBindJSON(&userModel)

	if errModel != nil && errParse != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errModel.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	updateUserModel, errUpdate := controller.userService.UpdateUser(id, &userModel)

	if errUpdate != nil {
		response := helper.BuildErrorResponseWithoutMessage(errUpdate.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	response := helper.BuildResponse(true, helper.Success, updateUserModel)
	ctx.JSON(response.Code, response)
}

func (controller *userController) DeleteUser(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	id, errParse := strconv.Atoi(ctx.Param("id"))

	if errParse != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errParse.Error(), helper.EmptyObj{})
		ctx.JSON(response.Code, response)
		return
	}

	errDelete := controller.userService.DeleteUser(id)

	if errDelete != nil {
		response := helper.BuildErrorResponseWithoutMessage(errDelete.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	response := helper.BuildResponse(true, helper.Success, nil)
	ctx.JSON(response.Code, response)
}
