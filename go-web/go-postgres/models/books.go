package models

import (
	"database/sql"
	"log"
)

type Book struct {
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type Bookstore struct {
	Books []Book `json:"books"`
}

func SelectBooks() (Bookstore, bool) {
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		log.Println(err)
		return Bookstore{}, false
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	books := Bookstore{}
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			log.Println(err)
		}
		books.Books = append(books.Books, bk)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return Bookstore{}, false
	}
	return books, true
}

func SelectBook(isbn string) (Book, bool) {
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)
	bk := Book{}
	err := row.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		return Book{}, false
	case err != nil:
		return Book{}, false
	}
	return bk, true
}

func InsertBook(bk *Book) bool {
	_, err := db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4);", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
