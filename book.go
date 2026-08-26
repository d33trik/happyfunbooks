package happyfunbooks

import "fmt"

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		ID:     "xyz",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func BookToString(b Book) string {
	return fmt.Sprintf("%v, by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func GetAllBooks() []Book {
	return catalog
}

func GetBook(ID string) (Book, bool) {
	for _, book := range catalog {
		if book.ID == ID {
			return book, true
		}
	}
	return Book{}, false
}
