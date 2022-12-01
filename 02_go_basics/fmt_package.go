package main

import "fmt"

func main() {
	name := "Alessandro"
	fmt.Println("Hello. My name is", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean value:", (a+b)/2)

	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f\n", 5, 6.8)

	fmt.Println("He says: \"Hello Go!\"")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The circumference of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q\n", figure)

	fmt.Printf("The circumference of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("Figure is of type %T\n", figure)
	fmt.Printf("Radius is of type %T\n", radius)
	fmt.Printf("Pi is of type %T\n", pi)

	closed := true
	fmt.Printf("File closed: %t\n", closed)

	fmt.Printf("%08b\n", 55)

	fmt.Printf("%x\n", 100)

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %.3f\n", x*y)
}
