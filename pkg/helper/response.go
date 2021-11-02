package helper

import (
	"fmt"
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	findErrorEnum(splittedError[0])

	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

func findErrorEnum(message string) {
	test := StatusMap(message)
	fmt.Println(test)
}

const (
	RecordNotFound = "record not found"
)

var statusMap = map[string]int{
	RecordNotFound: 403,
}

func StatusMap(code string) int {
	return statusMap[code]
}

// var statusText = map[string]int{
// 	RecordNotFound: 403,
// }

// func StatusText(code string) int {
// 	return statusText[code]
// }
