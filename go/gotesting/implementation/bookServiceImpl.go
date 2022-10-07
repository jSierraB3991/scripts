package implementation

import (
	"context"
	"errors"

	"github.com/segmentio/ksuid"
	"gitlab.com/eliotandelon/gotesting/domain/book"
)

var ErrorInvalidBookId = errors.New("Invalid book id")

type BookServiceImplementation struct {
	Database book.DBBookInteractor
}

func NewBookServiceImplementation(database book.DBBookInteractor) *BookServiceImplementation {
	return &BookServiceImplementation{Database: database}
}

func (this *BookServiceImplementation) ListBooks(ctx context.Context) ([]*book.Book, error) {
	return this.Database.AllBooks(ctx)
}

func (this *BookServiceImplementation) Book(ctx context.Context, id string) (*book.Book, error) {
	if id == "" {
		return nil, ErrorInvalidBookId
	}
	return this.Database.GetBookById(ctx, id)
}

func (this *BookServiceImplementation) SaveBook(ctx context.Context, book *book.Book) error {
	id, err := ksuid.NewRandom()
	if err != nil {
		return err
	}

	book.Id = id.String()
	return this.Database.SaveBook(ctx, book)
}
