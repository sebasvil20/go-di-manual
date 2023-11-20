package providers

import (
	"github.com/sebasvil20/go-di/src/repository"
	"github.com/sebasvil20/go-di/src/service"
)

func ProvideBookService(repo *repository.BookRepository) *service.BookService {
	return &service.BookService{BookRepository: repo}
}
