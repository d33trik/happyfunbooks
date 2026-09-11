package main

import (
	"fmt"
	"os"

	"github.com/d33trik/happyfunbooks"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	ID := os.Args[1]

	catalog, err := happyfunbooks.OpenCatalog("testdata/catalog.json")
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
	}

	book, ok := catalog.GetBook(ID)
	if !ok {
		fmt.Println("Soory, I couldn't find that book in the catalog.")
		return
	}

	fmt.Println(book)
}
