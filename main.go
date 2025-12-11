package main

import (
	"net/http"
	"xriot/learn-golang-dependency-injection/app"
	"xriot/learn-golang-dependency-injection/controller"
	"xriot/learn-golang-dependency-injection/helper"
	"xriot/learn-golang-dependency-injection/middleware"
	"xriot/learn-golang-dependency-injection/repository"
	"xriot/learn-golang-dependency-injection/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
