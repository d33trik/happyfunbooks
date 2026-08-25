package happyfunbooks_test

import (
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
