package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	c1 := circle{radius: 6}
	r1 := rectangle{
		width:  108,
		height: 666,
	}

	print(c1)
	print(r1)
	a := c1.area()
	fmt.Println("Circle area:", a)
}
