package happyfunbooks_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/d33trik/happyfunbooks"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := happyfunbooks.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room, by Adam Nicolson (copies: 2)"
	got := input.String()
	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestGetAllBooks_RetrunsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []happyfunbooks.Book{
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
	got := catalog.GetAllBooks()
	slices.SortFunc(got, func(a, b happyfunbooks.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}

func TestGetBook_FindBooksInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := happyfunbooks.Book{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if want != got {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}

func TestGetBook_ReturnFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	catalog.AddBook(happyfunbooks.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func getTestCatalog() happyfunbooks.Catalog {
	return happyfunbooks.Catalog{
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

func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := happyfunbooks.Book{
		Copies: 5,
	}
	book.SetCopies(12)
	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}
