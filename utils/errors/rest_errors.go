package errors

import "net/http"

type RestErr struct {
	Message string
	Code int
	Error string
}

func NewBadRequestError(message string)(*RestErr) {
	return &RestErr{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}



func NewNotFoundError(message string)(*RestErr) {
	return &RestErr{
		Message: message,
		Code: http.StatusNotFound,
		Error: "not_found",
	}
}