package main

import (
	"fmt"
	"numbers"
)

func main() {
	fmt.Printf("%d is even: %t\n", 40, numbers.Even(40))

	fmt.Printf("%d is prime: %t\n", 13, numbers.IsPrime(13))
	fmt.Printf("%d is prime: %t\n", 45, numbers.IsPrime(45))
}
