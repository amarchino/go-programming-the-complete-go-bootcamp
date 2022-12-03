package main

import "fmt"

func main() {
	// Coding Exercise #1
	// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Coding Exercise #2
	// Change the code from the previous exercise and use the continue statement to print out all the numbers divisible by 7 between 1 and 50.
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Coding Exercise #3
	// Change the code from the previous exercise and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
	for i, count := 1, 0; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
		count++
		if count >= 3 {
			break
		}
	}
	fmt.Println()

	// Coding Exercise #4
	// Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
	for i := 1; i <= 500; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Coding Exercise #5
	// Using a for loop print out all the years from your birthday to the current year.
	// Use a variant of for loop where the post statement is moved inside the for block of code.
	for year := 1988; year <= 2022; {
		fmt.Printf("%d ", year)
		year++
	}
	fmt.Println()

	// Coding Exercise #6
	// Consider the following Go program.
	// Change the code and use a switch statement instead of an if statement.
	age := -9
	// if age < 0 || age > 100 {
	// 	fmt.Println("Invalid Age")
	// } else if age < 18 {
	// 	fmt.Println("You are minor!")
	// } else if age == 18 {
	// 	fmt.Println("Congratulations! You've just become major!")
	// } else {
	// 	fmt.Println("You are major!")
	// }
	switch true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

}
