package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Coding Exercise #1
	// Using a composite literal declare and initialize a slice of type string with 3 elements.
	// Iterate over the slice and print each element in the slice and its index.
	sl := []string{"A", "B", "C"}
	for i, v := range sl {
		fmt.Printf("Index: %d, Value: %q\n", i, v)
	}

	// Coding Exercise #2
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	mySlice := []float64{1.2, 5.6}
	mySlice = append(mySlice, 6)
	a := 10
	mySlice[0] = float64(a)
	mySlice = append(mySlice, 10.10)
	mySlice = append(mySlice, []float64{float64(a)}...)
	fmt.Println(mySlice)

	// Coding Exercise #3
	// 1. Declare a slice called nums with three float64 numbers.
	// 2. Append the value 10.1 to the slice
	// 3. In one statement append to the slice the values: 4.1, 5.5 and 6.6
	// 4. Print out the slice
	// 5. Declare a slice called n with two float64 values
	// 6. Append n to nums
	// 7. Print out the nums slice
	nums := []float64{0, 1, 2}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)
	n := []float64{-1, -2}
	nums = append(nums, n...)
	fmt.Println(nums)

	// Coding Exercise #4
	// Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at the command line.
	// The user should give between 2 and 10 numbers.
	sum, product := 0.0, 1.0
	if len(os.Args) > 10 {
		fmt.Println("Should give at most 10 numbers")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Should give at least 2 numbers")
		return
	}
	for idx, val := range os.Args[1:] {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Printf("Cannot parse number at index %d: %q\n", idx, val)
			return
		}
		sum += num
		product *= num
	}
	fmt.Printf("Sum: %f, Product: %f\n", sum, product)

	// Coding Exercise #5
	// Consider the following slice declaration:
	nums2 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	// Print those elements and their sum.
	sum2 := 0
	for _, val := range nums2[1 : len(nums2)-2] {
		sum2 += val
		fmt.Printf("%d ", val)
	}
	fmt.Printf("Sum: %d\n", sum2)

	// Coding Exercise #6
	// Consider the following slice declaration:
	friends := []string{"Marry", "John", "Paul", "Diana"}
	// Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	newFriends := make([]string, len(friends))
	copy(newFriends, friends)
	friends[0] = "Luke"
	newFriends[1] = "Mike"
	fmt.Printf("%v, %v\n", friends, newFriends)

	// Coding Exercise #7
	// Consider the following slice declaration:
	friends = []string{"Marry", "John", "Paul", "Diana"}
	// Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
	newFriends = append([]string{}, friends...)
	friends[0] = "Luke"
	newFriends[1] = "Mike"
	fmt.Printf("%v, %v\n", friends, newFriends)

	// Coding Exercise #8
	// Consider the following slice declaration:
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	// Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.
	// newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	newYears := append(years[:3], years[len(years)-3:]...)
	fmt.Printf("%v\n", newYears)
}
