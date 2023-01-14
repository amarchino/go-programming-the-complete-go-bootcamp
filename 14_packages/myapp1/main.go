package main

import (
	"fmt"
	calc1 "github.com/amarchino/go-math/calc"
	"github.com/amarchino/go-math/geometry"
	calc2 "github.com/amarchino/go-math/v2/calc"
)

func main() {
	fmt.Println(geometry.CubeVolume(0))
	fmt.Println(geometry.CubeVolume(5))

	fmt.Println(calc1.Sub(5, 2))
	fmt.Println(calc2.Add(1, 2, 3))
}
