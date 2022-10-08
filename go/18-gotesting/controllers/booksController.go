package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (this *API) ListBooks(c echo.Context) error {
	books, err := this.Service.ListBooks(c.Request().Context())
	if err != nil {
		return this.tryError(c, "Error try search all books")
	}

	booksResponse := this.mapBooksToResponse(books)
	return c.JSON(http.StatusOK, booksResponse)
}

func (this *API) SaveBook(c echo.Context) error {
	var bookRequest = BookRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&bookRequest)
	if err != nil {
		return this.tryError(c, "Inavlid parameters")
	}
	err = this.Service.SaveBook(c.Request().Context(), this.mapBookToDomain(bookRequest))
	if err != nil {
		return this.tryError(c, "Error Try save book in the database")
	}
	return c.NoContent(http.StatusCreated)
}

func (this *API) GetBookById(c echo.Context) error {
	idParam := c.Param("id")
	book, err := this.Service.Book(c.Request().Context(), idParam)
	if err != nil {
		return this.tryError(c, "Invalid id book")
	}

	booksResponse := this.mapBookToResponse(book)
	return c.JSON(http.StatusOK, booksResponse)
}
