package main

import "testing"

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room, by Adam Nicolson - 2 copies"
	got := BookToString(input)
	if want != got {
		t.Fatal("BookToString: wrong result")
	}
}
