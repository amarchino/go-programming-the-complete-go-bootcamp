package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	// Coding Exercise #1
	// Consider the following type and interface declaration.
	// type vehicle interface {
	// 	License() string
	// 	Name() string
	// }
	// type car struct {
	// 	licenceNo string
	// 	brand     string
	// }
	// 1. Create a Go program where car type implements the vehicle interface.
	// 2. Create a variable of type vehicle that holds a car struct value.
	// 3. Call the methods (Licence and Name) on the interface value declared at step 2
	// 4. Run the program without errors.
	var myCar vehicle
	myCar = car{licenceNo: "12348A", brand: "BMW"}
	fmt.Printf("License: %v, Name: %v\n", myCar.License(), myCar.Name())

	// Coding Exercise #2
	// 1. Declare a variable called empty of type empty interface. Print out its type.
	// 2. Assign an int value to the variable called empty. Print out its type.
	// 3. Assign a float64 value to empty. Print out its type.
	// 4. Assign an int slice value to empty. Print out its type.
	// 5. Add a new int value to the slice (empty variable).
	// 6. Print out the slice (empty variable).
	var empty interface{}
	empty = 5
	fmt.Printf("%T\n", empty)
	empty = 5.
	fmt.Printf("%T\n", empty)
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)
	empty = append(empty.([]int), 7)
	fmt.Printf("%v\n", empty)

	// Coding Exercise #3
	// There is an error in the following Go program. Try to identify the error, change the code and run the program without errors.
	// type cube struct {
	// 	edge float64
	// }
	// func volume(c cube) float64 {
	// 	return c.edge * c.edge * c.edge
	// }
	// func main() {
	// 	var x interface{}
	// 	x = cube{edge: 5}
	// 	v := volume(x)
	// 	fmt.Printf("Cube Volume: %v\n", v)
	// }
	var x interface{}
	x = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
