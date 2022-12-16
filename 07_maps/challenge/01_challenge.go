package main

import "fmt"

func main() {
	// Coding Exercise #1
	// 1. Declare a map called m1 which value is nil. Print out its type and value.
	// 2. Declare a map called m2. It's keys are of type int and values of type string. Initialize the map using a map literal with two key:value pairs.
	// 3. Add the following key: value to the map: 10: "Abba"
	// 4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.
	var m1 map[string]string
	fmt.Printf("%T - %#v\n", m1, m1)
	m2 := map[int]string{1: "One", 2: "Two"}
	m2[10] = "Abba"
	fmt.Println(m2[1])
	fmt.Println(m2[100])

	// Coding Exercise #2
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	var m11 map[int]bool = make(map[int]bool)
	m11[5] = true

	m21 := map[int]int{3: 10, 4: 40}
	m31 := map[int]int{3: 10, 4: 40}

	fmt.Println(fmt.Sprintf("%v", m21) == fmt.Sprintf("%v", m31))

	// Coding Exercise #3
	// Consider the following map declaration:
	// m := map[int]bool{"1": true, 2: false, 3: false}
	// 1. The above map declaration has an error. Solve the error!
	// 2. Delete a key:value pair from the map.
	// 3. Iterate over the map and print out each key and value.
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 3)
	for k, v := range m {
		println(k, v)
	}
}
