package main

import (
	"golang-jwt/app"
	"golang-jwt/controller"
	"golang-jwt/helper"
	"golang-jwt/repository"
	"golang-jwt/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	//connect to mysql database
	db := app.Database()

	//validation
	validate := validator.New()

	//repository
	userRepository := repository.NewUsersRepositoryImpl()

	userService := service.NewUserServiceImpl(
		userRepository,
		db,
		validate,
	)

	//controller
	userController := controller.NewUsersControllerImpl(userService)

	//initialize
	router := httprouter.New()

	//router API
	router.POST("/api/V1/user", userController.Create)
	router.PUT("/api/V1/user/:id", userController.Update)
	router.DELETE("/api/V1/user/:id", userController.Delete)
	router.GET("/api/V1/user/:id", userController.GetById)
	router.GET("/api/V1/user", userController.GetAll)

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}

	err := server.ListenAndServe()
	//fmt.Println(reflect.TypeOf(err))
	helper.SetPanicError(err)
}
