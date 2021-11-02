package helper

import (
	"net/http"
	"strings"
)

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Code:    http.StatusOK,
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	_, code := findErrorEnum(message)

	res := Response{
		Code:    code,
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
func BuildErrorResponseWithoutMessage(err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	message, code := findErrorEnum(splittedError[0])

	res := Response{
		Code:    code,
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func findErrorEnum(errorMessage string) (string, int) {
	code := StatusMap(errorMessage)

	if code == 0 {
		code = 500
	}

	if errorMessage == "" {
		errorMessage = "Server Error"
	}
	return errorMessage, code
}

const (
	Success = "succes"

	BadRequest            = "bad request"
	UserEmpty             = "user is empty"
	UserNameEmpty         = "user name is empty"
	UserEmailEmpty        = "user email is empty"
	UserPassEmpty         = "user password is empty"
	IdMatchedAnotherEmail = "user with that email already exists, please check id or use another email"

	EmailOrPassNotMatch = "email or password not match"
	CheckCredential     = "please check again your credential"
	TokenNotFound       = "token not found"
	TokenNotValid       = "token is not valid"

	RecordNotFound = "record not found"
	UserExist      = "user with that email already exists"

	UserNotFound = "user with that id does not exist"

	ServerError = "server error"
)

var statusMap = map[string]int{
	BadRequest:            http.StatusBadRequest,
	UserEmpty:             http.StatusBadRequest,
	UserNameEmpty:         http.StatusBadRequest,
	UserEmailEmpty:        http.StatusBadRequest,
	UserPassEmpty:         http.StatusBadRequest,
	IdMatchedAnotherEmail: http.StatusBadRequest,

	EmailOrPassNotMatch: http.StatusUnauthorized,
	CheckCredential:     http.StatusUnauthorized,
	TokenNotFound:       http.StatusUnauthorized,
	TokenNotValid:       http.StatusUnauthorized,

	RecordNotFound: http.StatusForbidden,
	UserExist:      http.StatusForbidden,

	UserNotFound: http.StatusNotFound,

	ServerError: http.StatusInternalServerError,
}

func StatusMap(code string) int {
	return statusMap[code]
}
