package app

import "github.com/sebasvil20/go-di/src/app/providers"

func Run() {
	bookRepo := providers.ProvideBookRepository()
	bookService := providers.ProvideBookService(bookRepo)
	bookController := providers.ProvideBookController(bookService)
	router := HttpRouter(bookController)
	router.Run(":8080")
}
