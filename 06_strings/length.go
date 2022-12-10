package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Codruța"
	fmt.Println(len(name))

	fmt.Println(len("友達"))

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("友達")

	fmt.Println(n, m)
}
