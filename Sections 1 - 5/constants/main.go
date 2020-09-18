package main

import "fmt"

func main() {
	const days int = 7
	const pi = 3.14

	const (
		a         = 108
		b float32 = 66.6
	)

	const n, m int = 5, 6

	const (
		min1 = -500
		min2
		min3
	)
	fmt.Println(min1, min2, min3)

	const l1 = len("vejce")

	const x = 5
	const y = 1.1
	fmt.Printf("x times y is %v type of x ís %T, type of y is %T\n", x*y, x, y)
	fmt.Printf("type of x is %T\n", x)

	const (
		hu1 = iota
		hu2
		hu3
	)
	fmt.Println(hu1, hu2, hu3)

	const (
		pej1 = iota * iota
		pej2
		pej3
		pej4
		_
		pej5
	)
	fmt.Println(pej1, pej2, pej3, pej4, pej5)
}
