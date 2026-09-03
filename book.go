package happyfunbooks

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func (b Book) String() string {
	return fmt.Sprintf("%v, by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func GetAllBooks(catalog map[string]Book) []Book {
	return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}

func AddBook(catalog map[string]Book, book Book) {
	catalog[book.ID] = book
}

func GetCatalog() map[string]Book {
	return map[string]Book{
		"abc": {
			ID:     "abc",
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		"xyz": {
			ID:     "xyz",
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
}
