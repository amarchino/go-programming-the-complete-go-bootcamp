package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14159
	const secondsInHour = 3600

	duration := 234

	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// x, y := 5, 0
	// fmt.Println(x / y)

	const a = 5
	const b = 0
	// fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 4, 5

	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)

	// Constant rules
	// 1. you cannot change a constant
	// 2. you cannot initialize a contant at runtime
	// 3. you cannot use a variable to initialize a constant
	// 4. you can use len if the argument is a constant
}
