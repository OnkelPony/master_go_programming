package main

import "fmt"

func main() {
	a, b := 108, 66.6

	fmt.Println(a + 6)
	fmt.Println(b - 10.8)
	fmt.Println(a * 666)
	fmt.Println(b / 10.8)

	fmt.Println(a * int(b))
	fmt.Println(float64(a) / b)

	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

	a += 6
	b -= 9
	a *= 8
	b /= 2
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(108 == 108)
	fmt.Println(108 != 108)
	fmt.Println(108 >= 108)
	fmt.Println(108 <= 108)
	fmt.Println(108 > 108)
	fmt.Println(108 < 108)

	fmt.Println(10.8 < 108 && 66.6 < 666)
	fmt.Println(10.8 > 108 || 66.6 > 666)
	fmt.Println(!(108 < 666))

}
