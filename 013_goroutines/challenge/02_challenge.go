package main

import (
	"fmt"
	"strings"
)

func power(n int, ch chan int) {
	ch <- n * n
}

func main() {
	/*
			Coding Exercise #1
		1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64
		2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3. Both work with data of type rune.
		3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.
		4. Print out the type of all the channels declared.
	*/
	var c1 chan float64
	c2 := make(<-chan rune)
	c3 := make(chan<- rune)
	c4 := make(chan int, 10)
	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
	fmt.Println(strings.Repeat("#", 10))

	/*
			Coding Exercise #2
		Create a function literal (a.k.a. anonymous function) that sends the string value if receives as argument to main func using a channel.
	*/
	ch := make(chan string)
	go func(a string) {
		ch <- a
	}("Gophers")
	fmt.Printf("Value received: %s\n", <-ch)
	fmt.Println(strings.Repeat("#", 10))

	/*
			Coding Exercise #3
		There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
		package main

		import (
			"fmt"
		)

		func main() {
			c := make(<-chan int)

			go func(n int) {
				c <- n
			}(100)

			fmt.Println(<-c)
		}
	*/
	c := make(chan int)
	go func(n int) {
		c <- n
	}(100)
	fmt.Println(<-c)
	fmt.Println(strings.Repeat("#", 10))

	/*
			Coding Exercise #4
		Create a goroutine named power() that has one parameter of type int, calculates the square value of its parameter and then sends  the result into a channel.
		In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.
		Print out the square values.
		A square(or raising to power 2) is the result of multiplying a number by itself. e.g., 25 is the square of 5.
	*/
	for i := 1; i <= 50; i++ {
		go power(i, c)
	}
	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}
	fmt.Println(strings.Repeat("#", 10))

	/*
			Coding Exercise #5
		Change the program from Exercise #4 and calculate the square of all values between 1 and 50 included using an anonymous function.
	*/
	for i := 1; i <= 50; i++ {
		go func(n int) {
			c <- n * n
		}(i)
	}
	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}
}
