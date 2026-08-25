package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(b Book) string {
	return fmt.Sprintf("%v, by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func main() {
	fmt.Println("Books in stock:")

	book := Book{
		Title:  "Master and Commander",
		Author: "Patrick O'Brian",
		Copies: 17,
	}
	fmt.Println(BookToString(book))

	book = Book{
		Title:  "A Morbid Taste for Bones",
		Author: "Ellis Peters",
		Copies: 42,
	}
	fmt.Println(BookToString(book))
}
