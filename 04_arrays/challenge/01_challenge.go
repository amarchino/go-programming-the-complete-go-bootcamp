package main

import "fmt"

func main() {
	// Coding Exercise #1
	// 1. Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	//   Print out the cities array and notice the value of its elements.
	var cities [2]string
	fmt.Printf("%#v\n", cities)
	// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	//   Print out the grades array and notice the value of its elements.
	var grades = [3]float64{3.14159, 2.7182}
	fmt.Printf("%#v\n", grades)
	// 3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
	var taskDone = [...]bool{true, true, false}
	fmt.Printf("%#v\n", taskDone)
	// 4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println("cities["+fmt.Sprintf("%d", i)+"]:", cities[i])
	}
	// 5. Iterate over grades using the range keyword and print out element by element.
	for i, v := range grades {
		fmt.Println("grades["+fmt.Sprintf("%d", i)+"]:", v)
	}

	// Coding Exercise #2
	// Consider the following array declaration:
	nums := [...]int{30, -1, -6, 90, -6}
	// Write a Go program that counts the number of positive even numbers in the array.
	count := 0
	for _, v := range nums {
		if v%2 == 0 {
			count++
		}
	}
	fmt.Println("The number of even numbers in nums is", count)

	// Coding Exercise #3
	// There are some errors in the following Go program. Try to identify the errors, change the code, and run the program without errors.
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6
	a := 10
	// myArray[0] = a
	myArray[0] = float64(a)
	// myArray[3] = 10.10
	myArray[2] = 10.10
	fmt.Println(myArray)

}
