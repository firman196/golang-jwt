package controller

import (
	"golang-jwt/model/web"
	"golang-jwt/model/web/users"
	"golang-jwt/service"
	"golang-jwt/utils"
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
	utils.JsonDecode(request, &userCreateRequest)

	response := controller.UsersService.Create(request.Context(), userCreateRequest)
	webResponse := web.GeneralResponse{
		Status: "success",
		Data:   response,
	}

	utils.JsonEncode(writer, webResponse)
}

func (controller *UsersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := users.UsersUpdateRequest{}
	utils.JsonDecode(request, &userUpdateRequest)
	userId, err := strconv.Atoi(params.ByName("id"))
	utils.SetPanicError(err)

	response := controller.UsersService.Update(request.Context(), int16(userId), userUpdateRequest)
	webResponse := web.GeneralResponse{
		Status: "success",
		Data:   response,
	}

	utils.JsonEncode(writer, webResponse)
}

func (controller *UsersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("id"))
	utils.SetPanicError(err)

	controller.UsersService.Delete(request.Context(), int16(userId))
	response := web.GeneralResponse{
		Status: "success",
	}
	utils.JsonEncode(writer, response)
}

func (controller *UsersControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := strconv.Atoi(params.ByName("id"))

	utils.SetPanicError(err)

	data := controller.UsersService.GetById(request.Context(), int16(userId))

	response := web.GeneralResponse{
		Status: "success",
		Data:   data,
	}
	utils.JsonEncode(writer, response)
}

func (controller *UsersControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datas := controller.UsersService.GetAll(request.Context())

	response := web.GeneralResponse{
		Status: "success",
		Data:   datas,
	}

	utils.JsonEncode(writer, response)
}

func (controller *UsersControllerImpl) Auth(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userAuthRequest := users.UserAuthRequest{}
	utils.JsonDecode(request, &userAuthRequest)

	response := controller.UsersService.Auth(request.Context(), userAuthRequest)

	webResponse := web.GeneralResponse{
		Status: "success",
		Data:   response,
	}

	utils.JsonEncode(writer, webResponse)
}

func (controller *UsersControllerImpl) RefreshToken(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	token := request.Header.Get("Authorization")

	response := controller.UsersService.RefreshToken(request.Context(), token)

	webResponse := web.GeneralResponse{
		Status: "Success",
		Data:   response,
	}

	utils.JsonEncode(writer, webResponse)
}
