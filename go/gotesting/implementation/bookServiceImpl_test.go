package implementation

import (
	"context"
	"os"
	"testing"

	"gitlab.com/eliotandelon/gotesting/database"
	"gitlab.com/eliotandelon/gotesting/domain/book"
)

var s book.BookService

func TestMain(m *testing.M) {
	repo := &database.BookRepositoryMock{}
	s = NewBookServiceImplementation(repo)

	code := m.Run()
	os.Exit(code)
}

func TestListBooks(t *testing.T) {
	testCases := []struct {
		Name       string
		ExpectEror error
	}{
		{Name: "List Books", ExpectEror: nil},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := s.ListBooks(ctx)
			if err != tc.ExpectEror {
				t.Errorf("Error %v, got %v", tc.ExpectEror, err)
			}
		})
	}
}

func TestBook(t *testing.T) {
	testCases := []struct {
		Name       string
		Id         string
		ExpectEror error
	}{
		{Name: "Book by id", Id: "B001", ExpectEror: nil},
		{Name: "Invalid Id", Id: "", ExpectEror: ErrorInvalidBookId},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := s.Book(ctx, tc.Id)
			if err != tc.ExpectEror {
				t.Errorf("Error %v, got %v", tc.ExpectEror, err)
			}
		})
	}
}
