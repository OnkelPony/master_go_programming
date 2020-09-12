package main

import "fmt"

func main() {
	var i1 int8 = -108
	fmt.Printf("%T\n", i1)

	var i2 int16 = 666
	fmt.Printf("%T\n", i2)

	var i3 int64 = 108_666
	fmt.Printf("%T %d\n", i3, i3)

	var f1, f2, f3 = -1.3, .108, 666.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var r rune = 'Đ'
	fmt.Printf("Type: %T, dec: %d, hex: %x, bin: %b, oct: %o\n", r, r, r, r, r)

	var b bool = true
	fmt.Printf("%T\n", b)

	var s string = "HEllo Go!"
	fmt.Printf("%T\n", s)

	var numbers = [6]int{108, 666, -108, -666, 0, 1}
	fmt.Printf("%T\n", numbers)

	var cities = []string{"Jerusalem", "Tel Aviv", "Eilat"}
	fmt.Printf("%T\n", cities)

	var balances = map[string]float32{
		"NIS": 4_000_000_000,
		"USD": 666.108,
	}
	fmt.Printf("%T\n", balances)

	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	var x int = 2
	var y = &x
	fmt.Printf("%T %v\n", y, y)

	fmt.Printf("%T\n", funkce)
}

func funkce() {

}
