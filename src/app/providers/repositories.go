package providers

import "github.com/sebasvil20/go-di/src/repository"

func ProvideBookRepository() *repository.BookRepository {
	mockedDB := map[string]any{}
	return &repository.BookRepository{DBMock: mockedDB}
}
