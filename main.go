package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")

	title, author := "Master and Commander", "Patrick O'Brian"
	printBook(title, author)

	title, author = "A Morbid Taste for Bones", "Ellis Peters"
	printBook(title, author)
}

func printBook(title, author string) {
	fmt.Println(title, "by", author)
}
