package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive!")
	}
	_ = inStock

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's expensive!")
	}

	age := 50
	if age >= 0 && age < 18 {
		fmt.Printf("You canno vote! Prise return in %d years!\n", 18-age)
	} else if age == 18 {
		fmt.Printf("This is your first vote!\n")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's important!\n")
	} else {
		fmt.Printf("Invalid age!\n")
	}
}
