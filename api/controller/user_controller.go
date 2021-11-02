package controller

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/service"
	"net/http"
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
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id, errParse := strconv.Atoi(ctx.Param("id"))

	if errParse != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errParse.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := controller.userService.GetById(id)

	if err != nil {
		if err.Error() == helper.UserNotFound {
			response := helper.BuildErrorResponse(helper.UserNotFound, err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusNotFound, response)
		} else {
			response := helper.BuildErrorResponse(helper.ServerError, err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusInternalServerError, response)
		}
		return
	}

	response := helper.BuildResponse(true, helper.Success, user)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) GetAll(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	users, err := controller.userService.GetAll()
	if err != nil {

		if err.Error() == helper.UserNotFound {
			response := helper.BuildErrorResponse(helper.UserNotFound, err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusNotFound, response)
		} else {
			response := helper.BuildErrorResponse(helper.ServerError, err.Error(), helper.EmptyObj{})
			ctx.JSON(http.StatusInternalServerError, response)
		}
		return
	}

	response := helper.BuildResponse(true, helper.Success, users)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) UpdateUser(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var userModel model.User

	id, errParse := strconv.Atoi(ctx.Param("id"))
	errModel := ctx.ShouldBindJSON(&userModel)

	if errModel != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errModel.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if errParse != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errParse.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updateUserModel, errUpdate := controller.userService.UpdateUser(id, &userModel)

	if errUpdate != nil {
		if errUpdate.Error() == helper.UserNotFound {
			response := helper.BuildErrorResponse(helper.UserNotFound, errUpdate.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		} else {
			response := helper.BuildErrorResponse(helper.ServerError, errUpdate.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		}
		return
	}

	response := helper.BuildResponse(true, helper.Success, updateUserModel)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) DeleteUser(ctx *gin.Context) {
	response := helper.BuildErrorResponse(helper.ServerError, helper.ServerError, helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
