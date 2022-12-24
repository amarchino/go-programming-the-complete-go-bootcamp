package main

import "fmt"

func main() {
	// Coding Exercise #1
	// 1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
	// 2. Declare and initialize two values of type person, one called me and another called you.
	// 3. Print out the struct values.
	type person struct {
		name           string
		age            int
		favoriteColors []string
	}
	me := person{name: "Me", age: 30, favoriteColors: []string{"Black"}}
	you := person{name: "You", age: 20, favoriteColors: []string{"Green", "Blue"}}
	fmt.Println(me)
	fmt.Println(you)

	// Coding Exercise #2
	// Consider the code from the previous exercise and:
	// 1. Change the name or the struct value called me to "Andrei"
	// 2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
	// 3. Add a new favorite color to the second struct value called you.
	// 4. Print out the struct values.
	me.name = "Andrei"
	yourFavoriteColors := you.favoriteColors
	fmt.Printf("%T - %#v\n", yourFavoriteColors, yourFavoriteColors)
	you.favoriteColors = append(you.favoriteColors, "Red")
	fmt.Println(me)
	fmt.Println(you)

	// Coding Exercise #3
	// Consider the code from Exercise #1.
	// Iterate and print out the favorite colors of the struct value called me.
	for _, v := range me.favoriteColors {
		fmt.Println(v)
	}

	// Coding Exercise #4
	// Change the code from Exercise #1 and:
	// 1. Create a new struct type called grades with  2 fields: grade and course
	// 2. Add another field of type grades to person struct type (embedded struct).
	// 3. Change the initialization of the struct values called me and you to contain information about the grades.
	// 4. Change the grade and the course of one struct value to "Golang" and 98.
	// 5. Print out the struct values.
	type grades struct {
		grade  int
		course string
	}
	type person2 struct {
		name           string
		age            int
		favoriteColors []string
		grades         grades
	}
	me2 := person2{name: "Me", age: 30, favoriteColors: []string{"Black"}, grades: grades{grade: 100, course: "Math"}}
	you2 := person2{name: "You", age: 20, favoriteColors: []string{"Green", "Blue"}, grades: grades{grade: 85, course: "History"}}
	you2.grades.course = "Golang"
	you2.grades.grade = 98
	fmt.Println(me2)
	fmt.Println(you2)
}
