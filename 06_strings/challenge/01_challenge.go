package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Coding Exercise #1
	// 1. Using the var keyword declare a string called name and initialize it with your name.
	// 2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
	// 3. Print the following string on multiple lines like this:
	//   Your name: `here the value of name variable`
	//   Country: `here the value of country variable`
	// 4. Print out the following strings:
	//   a) He says: "Hello"
	//   b) C:\Users\a.txt
	var name string = "Alessandro"
	country := "Italy"
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
	fmt.Println("He says: \"Hello\"")
	fmt.Println(`C\Users\a.txt`)

	// Coding Exercise #2
	// 1. Declare a rune called r that stores the non-ascii letter ă
	// 2. Print out the type of r
	// 3. Declare 2 strings that contains the values ma and m
	// 4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).
	// BTW: mamă means mother in Romanian.
	// Note: You should convert rune to string to contatenate it to another string.
	var r rune = 'ă'
	fmt.Printf("%T\n", r)
	str1, str2 := "ma", "m"
	str3 := str1 + str2 + string(r)
	fmt.Println(str3)

	// Coding Exercise #3
	// Consider the following string declaration: s1 := "țară means country in Romanian"
	// 1. Iterate over the string and print out byte by byte
	// 2. Iterate over the string and print out rune by rune and the byte index where the rune starts in the string
	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}
	fmt.Println()

	for i, v := range s1 {
		fmt.Printf("%d: %c\n", i, v)
	}
	fmt.Println()

	// Coding Exercise #4
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	s2 := "Go is cool!"
	x := s2[0]
	fmt.Println(x)

	// s2[0] = 'x'

	// printing the number of runes in the string
	fmt.Println(utf8.RuneCountInString(s2))

	// Coding Exercise #5
	// Consider the following string declaration:s := "你好 Go!"
	// 1. Convert the string to a byte slice.
	// 2. Print out the byte slice
	// 3. Iterate over the byte slice and print out each index and byte in the byte slice
	s := "你好 Go!"
	slice := []byte(s)
	fmt.Printf("%#v\n", slice)
	for i, v := range slice {
		fmt.Printf("%d: %v\n", i, v)
	}

	// Coding Exercise #6
	// Consider the following string declaration:s := "你好 Go!"
	// 1. Convert the string to a rune slice.
	// 2. Print out the rune slice
	// 3. Iterate over the rune slice and print out each index and rune in the rune slice
	runeSlice := []rune(s)
	fmt.Printf("%#v\n", runeSlice)
	for i, v := range runeSlice {
		fmt.Printf("%d: %v => %c\n", i, v, v)
	}
}
