package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(b Book) string {
	return ""
}

func TestBookToString_FormatsBookInfoAsString() {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := "Sea Room, by Adam Nicolson - 2 copies"
	got := BookToString(input)
	if want != got {
		panic("BookToString: wrong result")
	}
}

func main() {
	TestBookToString_FormatsBookInfoAsString()

	fmt.Println("Books in stock:")

	book := Book{
		Title:  "Master and Commander",
		Author: "Patrick O'Brian",
		Copies: 17,
	}
	fmt.Println(BookToString(book))

	book = Book{
		Title:  "A Morbid Taste for Bones",
		Author: "Ellis Peters",
		Copies: 42,
	}
	fmt.Println(BookToString(book))
}
