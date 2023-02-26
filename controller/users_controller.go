package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
