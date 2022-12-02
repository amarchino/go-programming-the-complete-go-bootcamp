package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Coding Exercise #1
	// Consider the following declarations:
	// var i int = 3
	// var f float64 = 3.2
	// Write a Go program that converts i to float64 and f to int.
	// Print out the type and the value of the variables created after conversion.
	// fmt.Printf("i: %T %v\n", float64(i), float64(i))
	// fmt.Printf("f: %T %v\n", int(f), int(f))

	// Consider the following declarations:
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"
	// Write a Go program that converts:
	// 1. i to string (int to string)
	// 2. s2 to int (string to int)
	// 3. f to string (float64 to string)
	// 4. s1 to float64  (string to float64)
	// 5. Print the value and the type for each variable created after conversion.
	i_conv := fmt.Sprintf("%d", i)
	s2_conv, _ := strconv.Atoi(s2)
	f_conv := fmt.Sprintf("%f", f)
	s1_conv, _ := strconv.ParseFloat(s1, 64)

	fmt.Printf("i: %T %v\n", i_conv, i_conv)
	fmt.Printf("s2: %T %v\n", s2_conv, s2_conv)
	fmt.Printf("f: %T %v\n", f_conv, f_conv)
	fmt.Printf("s1: %T %v\n", s1_conv, s1_conv)

	// Coding Exercise #3
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	// x, y := 4, 5.1
	// z := float64(x) * y
	// fmt.Println(z)

	// a := 5
	// b := 6.2 * float64(a)
	// fmt.Println(b)

	// Coding Exercise #4
	// Create a Go program that computes how long does it take for the Sunlight to reach the Earth if we know that the distance from the Sun to Earth
	// is 149.6 million km and the speed of light is 299,792,458 m / s
	// The formula used is: time = distance / speed
	var distance = 149.6
	var speedOfLight = 299792458
	fmt.Println((distance * float64(1000000) * float64(1000)) / float64(speedOfLight))

	// Coding Exercise #5
	// Write the correct logical operator (&&, || , !) inside the expression so that result1 will be false and result2 will be true.
	x, y := 0.1, 5
	var z float64
	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.
	result1 := x < z && int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1*5 || int(z) == 0
	fmt.Println(result2)

}
