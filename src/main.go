package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/go-di/src/app"
	"github.com/sebasvil20/go-di/src/controller"
	"github.com/sebasvil20/go-di/src/repository"
	"github.com/sebasvil20/go-di/src/service"
)

func main() {
	app.Run()
}

func x() {
	bookRepo := repository.BookRepository{DBMock: map[string]any{}}
	bookSrv := service.BookService{BookRepository: &bookRepo}
	bookCtrl := controller.BooksController{BooksService: &bookSrv}

	r := gin.Default()
	r.GET("/books", bookCtrl.GetBook)
	r.Run()
}
