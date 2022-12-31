package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type object interface {
	volume() float64
}

type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * math.Pow(c.edge, 2)
}
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}
func (c cube) getColor() string {
	return c.color
}
func measure(g geometry) (float64, float64) {
	return g.area(), g.volume()
}

func main() {
	c := cube{edge: 2.}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volume: %v\n", a, v)
}
