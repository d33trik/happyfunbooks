package happyfunbooks_test

import (
	"slices"
	"testing"

	"github.com/d33trik/happyfunbooks"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := happyfunbooks.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room, by Adam Nicolson (copies: 2)"
	got := happyfunbooks.BookToString(input)
	if want != got {
		t.Fatalf("want: %q, got: %q", want, got)
	}
}

func TestGetAllBooks_RetrunsAllBooks(t *testing.T) {
	want := []happyfunbooks.Book{
		{
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
	got := happyfunbooks.GetAllBooks()
	if !slices.Equal(want, got) {
		t.Fatalf("want: %#v, got: %#v", want, got)
	}
}
