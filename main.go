package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")

	title, author := "Master and Commander", "Patrick O'Brian"
	fmt.Println(title, "by", author)

	title, author = "A Morbid Taste for Bones", "Ellis Peters"
	fmt.Println(title, "by", author)
}
