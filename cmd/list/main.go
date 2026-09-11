package main

import (
	"fmt"

	"github.com/d33trik/happyfunbooks"
)

func main() {
	catalog, err := happyfunbooks.OpenCatalog("testdata/catalog.json")
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
		return
	}

	fmt.Println("Books in stock:")
	for _, book := range catalog.GetAllBooks() {
		fmt.Println(book)
	}
}
