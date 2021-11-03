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

// Login func for creates a new token.
// @Description Create a new token.
// @Summary create a new token
// @Tags Login
// @Accept json
// @Produce json
// @Param login body model.Login true "User Login Model"
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 400 {object} helper.Response{status=bool,message=string,errors=object,data=object} "bad request"
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential || email or password not match"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/auth/login [POST]
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

// Register New User
// @Description RegisterUser (Insert)
// @Summary create a new user
// @Tags Register (Insert)
// @Accept json
// @Produce json
// @Param register body model.Register true "User Register Model"
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 400 {object} helper.Response{status=bool,message=string,errors=object,data=object} "bad request || user with that email already exists, please check id or use another email"
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential || email or password not match"
// @Failure 403 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that email already exists"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/users [PUT]
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
