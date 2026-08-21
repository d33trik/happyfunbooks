package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func printBook(b Book) {
	fmt.Println(b.Title, "by", b.Author, "-", b.Copies, "copies")
}

func main() {
	fmt.Println("Books in stock:")

	book := Book{
		Title:  "Master and Commander",
		Author: "Patrick O'Brian",
		Copies: 17,
	}
	printBook(book)

	book = Book{
		Title:  "A Morbid Taste for Bones",
		Author: "Ellis Peters",
		Copies: 42,
	}
	printBook(book)
}
