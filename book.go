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

func (b *Book) SetCopies(n int) {
	b.Copies = n
}

type Catalog map[string]Book

func (c Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(c))
}

func (c Catalog) GetBook(ID string) (Book, bool) {
	book, ok := c[ID]
	return book, ok
}

func (c Catalog) AddBook(book Book) {
	c[book.ID] = book
}

func GetCatalog() Catalog {
	return Catalog{
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
