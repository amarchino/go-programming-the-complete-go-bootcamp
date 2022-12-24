package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}
	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	// page := lastBook.pages

	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Leo Tolstoj"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	aBook := book{title: "Anna Karenina", author: "Leo Tolstoj", year: 1878}
	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)
}
