package main

import "fmt"

func main() {
	fmt.Println("Scope means visibility")
	sayHello("Andrei")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water boiling point in degrees Fahrenheit:", tf)
}
