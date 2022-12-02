package main

import "fmt"

var version string = "3.1"

func main() {
	// Coding Exercise #1
	// 1. Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
	// var a int
	// var b float64
	// var c bool
	// var d string

	// 2. Using short declaration syntax declare and assign these values to variables x, y and z:
	// - 20
	// - 15.5
	// - "Gopher!"
	// x := 20
	// y := 15.5
	// z := "Gopher!"

	// 3. Using fmt.Println() print out the values of a, b, c, d, x, y and z.
	// fmt.Println(a, b, c, d, x, y, z)

	// 4. Try to run the program without error.

	// Coding Exercise #2
	// 1. Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.
	// var (
	// 	a int
	// 	b float64
	// 	c bool
	// 	d string
	// )

	// 2. Declare x, y and z on a single line -> multiple short declarations
	// x, y, z := 20, 15.5, "Gopher!"

	// 3. Remove the statement that prints out the variables. See the error!

	// 4. Change the program to run without error using the blank identifier (_)
	// _, _, _, _, _, _, _ = a, b, c, d, x, y, z

	// Coding Exercise #3
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	var a float64 = 7.1
	x, y := true, 3.7
	a, x = 5.5, false
	_, _, _ = a, x, y

	// Coding Exercise #4
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	name := "Golang"
	fmt.Println(name)
}
