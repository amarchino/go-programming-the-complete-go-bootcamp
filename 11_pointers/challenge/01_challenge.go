package main

import "fmt"

func swap(a *float64, b *float64) {
	c := *b
	*b = *a
	*a = c
}

func main() {
	// Coding Exercise #1
	// Consider the following variable declaration x := 10.10
	// 1. Print out the address of x
	// 2. Declare a pointer called ptr that stores the address of x.
	// 3. Print out the type and the value of ptr.
	// 4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	x := 10.10
	fmt.Println(&x)
	ptr := &x
	fmt.Printf("%T %v\n", ptr, ptr)
	fmt.Printf("%p %v\n", &ptr, *ptr)

	// Coding Exercise #2
	// Consider the following variable declarations:
	x, y := 10, 2.
	ptrx, ptry := &x, &y
	// Declare a new variable called z and initialize the variable by dividing x by y through the pointers.
	z := *ptrx / *ptry
	fmt.Println(z)

	// Coding Exercise #3
	// Consider the following variable declaration:
	x, y = 5.5, 8.8
	// Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5
	swap(&x, &y)
	fmt.Println(x, y)
}
