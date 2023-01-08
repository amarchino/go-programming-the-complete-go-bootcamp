package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	// var ch chan int
	// fmt.Println(ch)

	// ch = make(chan int)
	// fmt.Println(ch)

	// c := make(chan int)

	// // SEND
	// c <- 10

	// // RECEIVE
	// num := <-c
	// _ = num

	// fmt.Println(<-c)

	// close(c)

	c := make(chan int)
	// only for receiving
	c1 := make(<-chan string)
	// only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)
	n := <-c
	fmt.Println("Value received:", n)
	fmt.Println("Exiting main")
}
