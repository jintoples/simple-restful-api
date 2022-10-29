package main

import (
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jintoples/simple-restful-api/app"
	"github.com/jintoples/simple-restful-api/controller"
	"github.com/jintoples/simple-restful-api/exception"
	"github.com/jintoples/simple-restful-api/helper"
	"github.com/jintoples/simple-restful-api/middleware"
	"github.com/jintoples/simple-restful-api/repository"
	"github.com/jintoples/simple-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
