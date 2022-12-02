package main

import "fmt"

func main() {
	a, b := 4, 2

	r := (a + b) / (a - b) * 2
	fmt.Println(r)

	r = 9 % a
	fmt.Println(r)

	a, b = 2, 3
	a += b
	fmt.Println(a)

	b -= 2
	fmt.Println(b)

	b *= 10
	fmt.Println(b)

	b /= 5
	fmt.Println(b)

	a %= 3
	fmt.Println(a)

	x := 1
	x += 1
	x++
	fmt.Println(x)

	x--
	fmt.Println(x)
	// fmt.Println(5 + x--)
}
