package service

import (
	"context"

	"github.com/sebasvil20/go-di/src/repository"
)

type IBookService interface {
	GetBook(ctx context.Context, id string) (any, error)
}

type BookService struct {
	BookRepository repository.IBookRepository
}

func (srv *BookService) GetBook(ctx context.Context, id string) (any, error) {
	return srv.BookRepository.GetBook(ctx, id)
}
