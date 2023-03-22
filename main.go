package main

import (
	"golang-jwt/app"
	"golang-jwt/controller"
	exception "golang-jwt/exception/api"
	"golang-jwt/helper"
	"golang-jwt/middleware"
	"golang-jwt/repository"
	"golang-jwt/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	//connect to mysql database
	db := app.Database()
	//validation
	validate := exception.NewValidation(db)
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
	router.POST("/api/V1/auth", userController.Auth)
	router.PUT("/api/V1/user/:id", userController.Update)
	router.DELETE("/api/V1/user/:id", userController.Delete)
	router.GET("/api/V1/user/:id", userController.GetById)
	router.GET("/api/V1/user", userController.GetAll)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:9000",
		//Handler: router,
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.SetPanicError(err)
}
