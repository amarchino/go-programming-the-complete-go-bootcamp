package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("Number of items inside os.Args:", len(os.Args))

	var result, _ = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T %v\n", os.Args[1], os.Args[1])
	fmt.Printf("%T %v\n", result, result)
}
