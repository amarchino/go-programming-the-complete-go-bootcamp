package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Coding Exercise #1
	// Create a new file in the current working directory called info.txt and then close the file.
	// If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Coding Exercise #2
	// Rename the file created at Exercise #1 to data.txt
	// Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
	_, err = os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist", err)
		}
	}
	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Coding Exercise #3
	// Remove the file created at exercise #1. Take care that the file is now called data.txt (it has been renamed at exercise #2).
	// Perform error handling.
	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Coding Exercise #4
	// Create a Go Program that reads the entire contents of a file called info2.txt into a string. You can use ioutil.ReadAll() or any other function you want.
	// The file exists in the current working directory.
	file, err = os.Open("info2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", bytes)

	// Coding Exercise #5
	// Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
	// The file exists in the current working directory.
	file, err = os.Open("info2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line:", line)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Coding Exercise #6
	// Create a Go Program that:
	// 1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.
	// 2. Write the string "The Go gopher is an iconic mascot!" to the file.
	file, err = os.OpenFile("info3.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("The Go gopher is an iconic mascot!")
	if err != nil {
		log.Fatal(err)
	}
}
