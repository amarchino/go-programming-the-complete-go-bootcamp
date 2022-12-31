package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("Print: %.2f\n", m)
}
func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}
func (b *book) discount() {
	b.price *= 0.9
}

func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	// Coding Exercise #1
	// 1. Create a new type called money. Its underlying type is float64.
	// 2. Create a method called print that prints out the money value with exact 2 decimal points.
	var m money = 6.452588
	m.print()

	// Coding Exercise #2
	// Consider the money type declared at exercise #1. Create a new method for the money type called printStr
	// that returns the money value as a string with 2 decimal points.
	fmt.Println("PrintStr: ", m.printStr())

	// Coding Exercise #3
	// 1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
	// 2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
	// 3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
	b := book{price: 100, title: "My book"}
	fmt.Println("Vat: ", b.vat())
	b.discount()
	fmt.Println(b)

	// Coding Exercise #4
	// A junior Go Programmer has just developed the following Go Program.
	// You want the change the book's price by calling changePrice method.
	// However, you notice that after calling the method the price is not changed.
	// You task is to change the code in order for the changePrice method to work as expected.
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}
	// changing the price
	bestBook.changePrice(11.99)
	fmt.Printf("Book's price: %.2f\n", bestBook.price)
}
