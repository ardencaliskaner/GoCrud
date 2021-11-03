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

// GetUserById  - (Authorization Required)
// @Summary Get User With Given Id
// @Description Get User With Given Id
// @Tags GetUserById
// @Accept json
// @Produce json
// @Param   id    path    int     true        "ID"
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 400 {object} helper.Response{status=bool,message=string,errors=object,data=object} "bad request"
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/users/{id} [get]
// @Security bearerAuth
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

// GetAllUsers  - (Authorization Required)
// @Summary List of Users
// @Description List of Users
// @Tags GetUsers
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/users [get]
// @Security bearerAuth
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

// UpdateUser Credentials With Given Id Of User - (Authorization Required)
// @Description UpdateUser
// @Summary update user with given user id
// @Tags UpdateUser
// @Accept json
// @Produce json
// @Param   id    path    int     true        "ID"
// @Param update body model.Update true "User Update Model"
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 400 {object} helper.Response{status=bool,message=string,errors=object,data=object} "bad request || user with that email already exists, please check id or use another email"
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/users/{id} [PATCH]
// @Security bearerAuth
func (controller *userController) UpdateUser(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	_, errToken := controller.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		response := helper.BuildErrorResponse(helper.CheckCredential, errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(response.Code, response)
		return
	}
	var userModel model.Update

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

// DeleteUser With Given Id Of User - (Authorization Required)
// @Description DeleteUser
// @Summary deleteUser user with given user id
// @Tags DeleteUser
// @Accept json
// @Produce json
// @Param   id    path    int     true        "ID"
// @Success 200 {object} helper.Response{status=bool,message=string,errors=object,data=object}
// @Failure 400 {object} helper.Response{status=bool,message=string,errors=object,data=object} "bad request"
// @Failure 401 {object} helper.Response{status=bool,message=string,errors=object,data=object} "unauthorized please check again your credential"
// @Failure 404 {object} helper.Response{status=bool,message=string,errors=object,data=object} "user with that id does not exist"
// @Router /v1/api/users/{id} [DELETE]
// @Security bearerAuth
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
