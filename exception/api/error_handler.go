package exception

import (
	"golang-jwt/helper"
	"golang-jwt/model/web"
	"net/http"
)

type BadRequestError struct {
	Error string
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{
		Error: error,
	}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequest(writer, request, err) {
		return
	}

	if notFound(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func badRequest(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, err := error.(BadRequestError)

	if err {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := web.GeneralResponse{
			Status: "Bad Request",
			Data:   exception.Error,
		}

		helper.JsonEncode(writer, response)
		return true
	} else {
		return false
	}
}

func notFound(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, err := error.(BadRequestError)

	if err {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := web.GeneralResponse{
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.JsonEncode(writer, response)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	response := web.GeneralResponse{
		Status: "Internal Server Error",
		Data:   error,
	}

	helper.JsonEncode(writer, response)
}
