package main

import "fmt"

func main() {
	var a int = 10
	var b = 5.4
	c := "Gopher"
	_, _, _ = a, b, c

	var (
		value    int
		price    float64
		name     string
		validity bool
	)
	fmt.Printf("value = %d, price = %f, name = %s, validity = %v", value, price, name, validity)
}
