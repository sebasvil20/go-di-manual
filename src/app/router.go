package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/go-di/src/controller"
)

func HttpRouter(
	bookController *controller.BooksController,
) *gin.Engine {
	r := gin.Default()
	r.GET("/books", bookController.GetBook)
	return r
}
