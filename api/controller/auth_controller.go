package controller

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(),
		jwtService:  service.NewJWTService(),
	}
}

func (controller *authController) Login(ctx *gin.Context) {
	var loginModel model.Login
	errDTO := ctx.ShouldBind(&loginModel)
	if errDTO != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userModel, verifyErr := controller.authService.VerifyCredential(loginModel.Email, loginModel.Password)

	if verifyErr != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, helper.InvalidCredential, helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(uint64(userModel.ID), 10))
	response := helper.BuildResponse(true, "OK!", generatedToken)
	ctx.JSON(http.StatusOK, response)
}

func (controller *authController) Register(ctx *gin.Context) {
	var registerModel model.Register

	errModel := ctx.ShouldBindJSON(&registerModel)

	if errModel != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errModel.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userModel, createuserErr := controller.authService.CreateUser(registerModel)

	if createuserErr != nil {
		if createuserErr.Error() == helper.UserExist {
			response := helper.BuildErrorResponse(helper.UserExist, createuserErr.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusForbidden, response)
		} else {
			response := helper.BuildErrorResponse(helper.ServerError, createuserErr.Error(), helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		}
		return
	}

	generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(uint64(userModel.ID), 10))
	response := helper.BuildResponse(true, "OK!", generatedToken)
	ctx.JSON(http.StatusOK, response)

}
