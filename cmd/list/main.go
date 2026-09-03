package main

import (
	"fmt"

	"github.com/d33trik/happyfunbooks"
)

func main() {
	catalog := happyfunbooks.GetCatalog()
	fmt.Println("Books in stock:")
	for _, book := range happyfunbooks.GetAllBooks(catalog) {
		fmt.Println(book)
	}
}
