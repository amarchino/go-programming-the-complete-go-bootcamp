package main

import "fmt"

func main() {
	// Coding Exercise #1
	// Consider the following variable declarations:
	// x, y, z := 10, 15.5, "Gophers"
	// score := []int{10, 20, 30}
	// Using fmt.Printf():
	// 1. Print each variable using a specific verb for its type
	// 2. Print the string value enclosed by double quotes ("Gophers")
	// 3. Print each variable using the same verb
	// 4. Print the type of y and score
	// fmt.Printf("x: %d, y: %f, z: %s\n", x, y, z)
	// fmt.Printf("z: %q\n", z)
	// fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	// fmt.Printf("y: %T, score: %T\n", y, score)

	// Coding Exercise #2
	// Consider the following constant declaration:
	const x float64 = 1.422349587101
	// Write a Go program that prints the value of x with 4 decimal points.
	fmt.Printf("%.4f\n", x)

	// Coding Exercise #3
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape

}
