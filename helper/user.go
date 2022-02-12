package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	ResponseCode    int         `json:"response code"`
	ResponseMessage string      `json:"response message"`
	Data            interface{} `json:"data"`
}

func APIResponse(responseCode int, responseMessage string, data interface{}) Response {
	jsonResponse := Response{
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
		Data:            data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
