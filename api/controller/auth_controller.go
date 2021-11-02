package controller

import (
	"GoCrud/pkg/helper"
	"GoCrud/pkg/model"
	"GoCrud/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	userService service.UserService
	jwtService  service.JWTService
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(),
		userService: service.NewUserService(),
		jwtService:  service.NewJWTService(),
	}
}

func (controller *authController) Login(ctx *gin.Context) {
	var loginModel model.Login
	errDTO := ctx.ShouldBind(&loginModel)
	if errDTO != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}
	userModel, verifyErr := controller.authService.VerifyCredential(loginModel.Email, loginModel.Password)

	if verifyErr != nil {
		response := helper.BuildErrorResponseWithoutMessage(verifyErr.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(uint64(userModel.Id), 10))
	response := helper.BuildResponse(true, helper.Success, generatedToken)
	ctx.JSON(response.Code, response)
}

func (controller *authController) Register(ctx *gin.Context) {
	var registerModel model.Register

	errModel := ctx.ShouldBindJSON(&registerModel)

	if errModel != nil {
		response := helper.BuildErrorResponse(helper.BadRequest, errModel.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	userModel, createuserErr := controller.userService.CreateUser(registerModel)

	if createuserErr != nil {
		response := helper.BuildErrorResponseWithoutMessage(createuserErr.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}

	generatedToken := controller.jwtService.GenerateToken(strconv.FormatUint(uint64(userModel.Id), 10))
	response := helper.BuildResponse(true, helper.Success, generatedToken)
	ctx.JSON(response.Code, response)
}
