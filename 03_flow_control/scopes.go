package main

// File scope
import (
	"fmt"
	f "fmt"
)

// Package scope
const done = false

var b int = 10

func main() {
	// Block scope
	var task = "Running"
	fmt.Println(task, done)

	const done = true
	fmt.Printf("done in main() is %v\n", done)

	f1()

	f.Println("Bye Bye!")
}

func f1() {
	// const done = true
	fmt.Printf("done in f1(): %v\n", done)

	a := 10
	_ = a
}
