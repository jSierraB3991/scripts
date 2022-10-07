package implementation

import (
	"context"
	"fmt"

	"gitlab.com/eliotandelon/gotesting/domain/book"
)

type BookServiceMock struct{}

func (BookServiceMock) ListBooks(ctx context.Context) ([]*book.Book, error) {
	fmt.Println("llegando al Service")
	return []*book.Book{}, nil
}

func (BookServiceMock) Book(ctx context.Context, id string) (*book.Book, error) {
	return &book.Book{}, nil
}

func (BookServiceMock) SaveBook(ctx context.Context, book *book.Book) error {
	return nil
}
