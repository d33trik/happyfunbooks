package main

import (
	"fmt"

	"github.com/d33trik/happyfunbooks"
)

func main() {
	fmt.Println("Books in stock:")

	for _, book := range happyfunbooks.GetAllBooks() {
		fmt.Println(happyfunbooks.BookToString(book))
	}
}
