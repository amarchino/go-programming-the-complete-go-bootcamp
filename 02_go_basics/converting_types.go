package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1

	// x = x * y
	x = x * int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5
	fmt.Printf("%T\n", a)

	var b int64 = 2
	fmt.Printf("%T\n", b)

	// a = b
	a = int(b)

	s := string(99)
	fmt.Printf("%T %v\n", s, s)

	// s1 := string(44.2)
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	f1, err := strconv.ParseFloat(s1, 64)
	_ = err

	fmt.Printf("%T\n", f1)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i Type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 Type is %T, s2 value is %q\n", s2, s2)
}
