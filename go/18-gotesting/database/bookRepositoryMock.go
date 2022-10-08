package database

import (
	"context"

	"gitlab.com/eliotandelon/gotesting/domain/book"
)

type BookRepositoryMock struct{}

func (BookRepositoryMock) AllBooks(ctx context.Context) ([]*book.Book, error) {
	return []*book.Book{}, nil
}

func (BookRepositoryMock) GetBookById(ctx context.Context, id string) (*book.Book, error) {
	return &book.Book{}, nil
}

func (BookRepositoryMock) SaveBook(ctx context.Context, book *book.Book) error {
	return nil
}
