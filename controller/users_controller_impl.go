package controller

import (
	"golang-jwt/helper"
	"golang-jwt/model/web"
	"golang-jwt/model/web/users"
	"golang-jwt/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UsersControllerImpl struct {
	UsersService service.UsersService
}

func NewUsersControllerImpl(userService service.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: userService,
	}
}

func (controller *UsersControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := users.UsersCreateRequest{}
	helper.JsonDecode(request, &userCreateRequest)

	response := controller.UsersService.Create(request.Context(), userCreateRequest)
	webResponse := web.GeneralResponse{
		Status: "success",
		Code:   200,
		Data:   response,
	}

	helper.JsonEncode(writer, webResponse)
}

func (controller *UsersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := users.UsersUpdateRequest{}
	helper.JsonDecode(request, &userUpdateRequest)

	response := controller.UsersService.Update(request.Context(), userUpdateRequest)
	webResponse := web.GeneralResponse{
		Status: "success",
		Code:   200,
		Data:   response,
	}

	helper.JsonEncode(writer, webResponse)
}

func (controller *UsersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("id"))
	helper.SetPanicError(err)

	controller.UsersService.Delete(request.Context(), int16(userId))
	response := web.GeneralResponse{
		Status: "success",
		Code:   200,
	}
	helper.JsonEncode(writer, response)
}

func (controller *UsersControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("id"))

	helper.SetPanicError(err)

	data := controller.UsersService.GetById(request.Context(), int16(userId))

	response := web.GeneralResponse{
		Status: "success",
		Code:   200,
		Data:   data,
	}
	helper.JsonEncode(writer, response)
}

func (controller *UsersControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datas := controller.UsersService.GetAll(request.Context())

	response := web.GeneralResponse{
		Status: "success",
		Code:   200,
		Data:   datas,
	}

	helper.JsonEncode(writer, response)
}
