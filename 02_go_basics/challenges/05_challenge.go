package main

import "fmt"

type duration int
type mile float64
type kilometer float64

var hour duration

const m2km = 1.609

func main() {
	// Coding Exercise #1
	// 1. Declare a new type type called duration. Have the underlying type be an int.
	// 2. Declare a variable of the new type called hour using the var keyword
	// 3. In function main:
	//   a. print out the value of the variable hour
	//   b. print out the type of the variable hour
	//   c. assign 3600 to the variable hour using the  = operator
	//   d. print out the value of hour
	fmt.Println(hour)
	fmt.Printf("%T\n", hour)
	hour = 3600
	fmt.Println(hour)

	// Coding Exercise #2
	// There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.
	var hour duration = 3600
	minute := 60
	fmt.Println(int(hour) != minute)

	// Coding Exercise #3
	// 1. Declare two defined types called mile and kilometer. Have the underlying type be an float64.
	// 2. Declare a constant called m2km equals 1.609  ( 1mile=1.609km)
	// 3. In function main:
	//   a. declare a variable called mileBerlinToParis of type mile with value 655.3
	//   b. declare a variable called kmBerlinToParis of type kilometer
	//   c. calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km. Assign the result to kmBerlinToParis
	//   d. print out the distance in km between Berlin and Paris
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(mileBerlinToParis) * kilometer(m2km)
	fmt.Println(kmBerlinToParis)
}
