package database

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/eliotandelon/gotesting/domain/book"
)

func (this MysqlRepository) AllBooks(ctx context.Context) ([]*book.Book, error) {
	rows, err := this.DB.QueryContext(ctx, "SELECT id, title, author FROM book")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var books = []*book.Book{}

	for rows.Next() {
		var singleBook = book.Book{}
		if err = rows.Scan(&singleBook.Id, &singleBook.Title, &singleBook.Author); err == nil {
			books = append(books, &singleBook)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}
	return books, nil

}

func (this MysqlRepository) GetBookById(ctx context.Context, id string) (*book.Book, error) {
	rows, err := this.DB.QueryContext(ctx, "SELECT id, title, author FROM book WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var book = book.Book{}

	for rows.Next() {
		if err = rows.Scan(&book.Id, &book.Title, &book.Author); err == nil {
			return &book, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}
	return &book, nil

}

func (this MysqlRepository) SaveBook(ctx context.Context, book *book.Book) error {
	_, err := this.DB.ExecContext(ctx,
		"INSERT INTO book(id, title, author) VALUES(?, ?, ?)",
		book.Id, book.Title, book.Author)
	return err

}
