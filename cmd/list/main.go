package main

import (
	"fmt"

	"github.com/d33trik/happyfunbooks"
)

func main() {
	fmt.Println("Books in stock:")

	book := happyfunbooks.Book{
		Title:  "Master and Commander",
		Author: "Patrick O'Brian",
		Copies: 17,
	}
	fmt.Println(happyfunbooks.BookToString(book))

	book = happyfunbooks.Book{
		Title:  "A Morbid Taste for Bones",
		Author: "Ellis Peters",
		Copies: 42,
	}
	fmt.Println(happyfunbooks.BookToString(book))
}
