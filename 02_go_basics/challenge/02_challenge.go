package main

import (
	"fmt"
	"math"
)

func main() {
	// Coding Exercise #1
	// Using the const keyword declare and initialize the following constants:
	// 1. daysWeek with value 7
	// 2. lightSpeed with value 299792458
	// 3. pi with value 3.14159
	// const daysWeek = 7
	// const lightSpeed = 299792458
	// const pi = 3.14159

	// Coding Exercise #2
	// Change the code from the previous exercise and declare all 3 constants as grouped constants.
	// Make them untyped.
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	// Coding Exercise #3
	// Calculate how many seconds are in a year.
	// STEPS:
	// 1. Declare secPerDay constant and initialize it to the number of seconds in a day
	// 2. Declare daysYear constant and initialize it to 365
	// 3. Use fmt.Printf() to print out the total number of seconds in a year.
	// EXPECTED OUTPUT:
	// There are 31536000 seconds in a year.
	const secPerDay = 60 * 60 * 24
	const daysYear = 365
	fmt.Printf("%d\n", secPerDay*daysYear)

	// Coding Exercise #4
	// There are an error in the following Go program. Try to identify the error, change the code and run the program without errors.
	// const x int = 10
	var m = []int{1: 3, 4: 5, 6: 8}
	_ = m

	// Coding Exercise #5
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b
	const x = 8
	const xc int = x
	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6

	// Coding Exercise #6
	// Using Iota declare the following months of the year: Jun, Jul and Aug
	// Jun, Jul and Aug are constant and their value is 6, 7 and 8.
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println(Jun, Jul, Aug)

}
