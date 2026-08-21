package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")

	title, author, copies := "Master and Commander", "Patrick O'Brian", 17
	printBook(title, author, copies)

	title, author, copies = "A Morbid Taste for Bones", "Ellis Peters", 42
	printBook(title, author, copies)
}

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")
}
