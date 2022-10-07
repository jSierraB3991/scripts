package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/eliotandelon/gotesting/domain/book"
)

type API struct {
	Service book.BookService
}

func NewApi(service book.BookService) *API {
	return &API{Service: service}
}

func (API) tryError(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"message": message})
}

func (API) mapBookToDomain(bookRequest BookRequest) *book.Book {
	return &book.Book{Author: bookRequest.Author, Title: bookRequest.Title}
}

func (API) mapBookToResponse(book *book.Book) BookResponse {
	return BookResponse{Id: book.Id, Title: book.Title, Author: book.Author}
}

func (this *API) mapBooksToResponse(bookDomain []*book.Book) []BookResponse {
	if len(bookDomain) == 0 {
		return []BookResponse{}
	}
	var bookResponse []BookResponse

	for _, book := range bookDomain {
		bookResponse = append(bookResponse, this.mapBookToResponse(book))
	}

	return bookResponse
}
