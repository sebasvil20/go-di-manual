package providers

import (
	"github.com/sebasvil20/go-di/src/controller"
	"github.com/sebasvil20/go-di/src/service"
)

func ProvideBookController(srv *service.BookService) *controller.BooksController {
	return &controller.BooksController{BooksService: srv}
}
