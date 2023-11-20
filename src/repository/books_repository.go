package repository

import (
	"context"
	"errors"
)

type IBookRepository interface {
	GetBook(ctx context.Context, id string) (any, error)
}

type BookRepository struct {
	DBMock map[string]any
}

func (repo *BookRepository) GetBook(ctx context.Context, id string) (any, error) {
	book, ok := repo.DBMock[id]
	if !ok {
		return nil, errors.New("no books found")
	}

	return book, nil
}
