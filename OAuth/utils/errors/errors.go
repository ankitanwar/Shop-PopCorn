package errors

import (
	"net/http"
)

//RestError : To handle the error
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

//NewBadRequest : To generate the bad request error
func NewBadRequest(message string) *RestError {
	err := &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
	return err
}

//NewNotFound : It returns an error when request obj is not present in the database
func NewNotFound(message string) *RestError {
	err := &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
	return err
}

//NewInternalServerError : It will return the internal server error
func NewInternalServerError(message string) *RestError {
	err := &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
	return err
}
