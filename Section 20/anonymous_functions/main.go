package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println("I'm an anonymous function.", msg)
	}("vole")

	a := increment(108)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
