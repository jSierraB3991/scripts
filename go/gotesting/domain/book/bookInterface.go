package book

import "context"

type DBBookInteractor interface {
	AllBooks(ctx context.Context) ([]*Book, error)
	GetBookById(ctx context.Context, id string) (*Book, error)
	SaveBook(ctx context.Context, book *Book) error
}

type BookService interface {
	ListBooks(ctx context.Context) ([]*Book, error)
	Book(ctx context.Context, id string) (*Book, error)
	SaveBook(ctx context.Context, book *Book) error
}
