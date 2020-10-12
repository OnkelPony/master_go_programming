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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)
	//s.volume()
	fmt.Printf("Circle area: %v\n", s.area())
	fmt.Printf("Sphere volume: %v\n", s.(circle).volume())
	ball, ok := s.(circle)
	if ok {
		fmt.Printf("Ball volume: %v\n", ball.volume())
	}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v was circle type\n", value)
	case rectangle:
		fmt.Printf("%#v was rectangle type\n", value)
	}
}
