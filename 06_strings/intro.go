package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello!\"")
	fmt.Println(`He says. "Hello!"`)

	// Raw string
	s2 := `I like \n Go!`
	fmt.Println(s2)

	fmt.Println("Print: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike`)

	fmt.Println(`C:\Users\march`)
	fmt.Println("C:\\Users\\march")

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])

	// s3[5] = 'x'

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}
