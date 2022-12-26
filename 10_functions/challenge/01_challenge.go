package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func cube(a float64) float64 {
	return math.Pow(a, 3)
}

func f1(n uint) (float64, float64) {
	f, s := float64(n), 0.
	for i := uint(1); i < n; i++ {
		f *= float64(i)
		s += float64(i)
	}

	return f, s
}

func myFunc(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Cannot convert " + s + " to int.")
	}
	return n + 11*n + 111*n
}

func sum(n ...int) (s int) {
	for _, v := range n {
		s += v
	}
	return s
}

func sum1(n ...int) (s int) {
	for _, v := range n {
		s += v
	}
	return
}

func searchItem(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func searchItem2(slice []string, s string) bool {
	for _, v := range slice {
		if strings.EqualFold(v, s) {
			return true
		}
	}
	return false
}

func print(msg string) {
	fmt.Println(msg)
}

func f2(v int) {
	fmt.Println(v)
}

func main() {
	// Coding Exercise #1
	// Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).
	fmt.Println(cube(2))

	// Coding Exercise #2
	// Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
	// a) the factorial of n
	// b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
	// Test the program by calling the function.
	fmt.Println(f1(6))

	// Coding Exercise #3
	// Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes
	// (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nnn
	// Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107
	fmt.Println(myFunc("5"))
	fmt.Println(myFunc("9"))

	// Coding Exercise #4
	// Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// Coding Exercise #5
	// Change the function from the previous exercise and use a `naked return`.
	fmt.Println(sum1(1, 2, 3, 4, 5, 6))

	// Coding Exercise #6
	// Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.
	// The function should search for the string (the second parameter) in the slice (the first parameter) and returns true
	// if it finds the string in the slice and false otherwise. The function does a case-sensitive search.
	// Call the function and see how it works.
	// Example:
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result)
	result = searchItem(animals, "pig")
	fmt.Println(result)

	// Coding Exercise #7
	// Change the function from the previous exercise to do a case-insensitive search.
	// Example:
	animals2 := []string{"Lion", "tiger", "bear"}
	result = searchItem2(animals2, "beaR")
	fmt.Println(result)
	result = searchItem2(animals2, "lion")
	fmt.Println(result)

	// Coding Exercise #8
	// Consider the following Go program that prints out:
	// The Go gopher is the iconic mascot of the Go project.
	// Hello, Go playground!
	// func print(msg string) {
	// 	fmt.Println(msg)
	// }
	// func main() {
	// 	print("The Go gopher is the iconic mascot of the Go project.")
	// 	fmt.Println("Hello, Go playground!")
	// }
	// Modify only the line in the main() body function where the print() function is invoked so that the program will print out
	// Hello, Go playground! and then The Go gopher is the iconic mascot of the Go project.
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")

	// Coding Exercise #9
	// Create a function that takes in an int value and prints out that value.
	// Assign the function to a variable, print out the type of the variable and then call that function through the variable name.
	f := f2
	fmt.Printf("%T\n", f)
	f(5)
}
