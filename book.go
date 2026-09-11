package happyfunbooks

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
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

func (b *Book) SetCopies(n int) error {
	if n < 0 {
		return fmt.Errorf("negative number of copies: %d", n)
	}
	b.Copies = n
	return nil
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

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var catalog Catalog
	err = json.NewDecoder(file).Decode(&catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}
