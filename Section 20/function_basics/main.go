package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1() function")
}

func f2(a, b int) {
	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float32, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	fmt.Println("s =", s)
	s = a + b
	return
}
func main() {
	f1()
	f2(108, 666)
	f3(1, 3, 5, 66.6, 10.8, "golang")
	fmt.Println(f4(10.8))
	fmt.Println(f5(108, 666))
	ss := sum(4, 5)
	fmt.Println(ss)
}
