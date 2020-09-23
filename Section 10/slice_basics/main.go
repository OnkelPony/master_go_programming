package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil:", cities == nil)
	fmt.Printf("cities is %#v\n", cities)

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Println(nums)

	type names []string
	friends := names{"Aletta", "Katsumi"}
	fmt.Printf("%#v\n", friends)

	x := numbers[0]
	fmt.Println("x is", x)

	numbers[1] = 108
	fmt.Printf("%#v\n", numbers)

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])
	}

	var n []int
	n = numbers
	_ = n

	var nn []int
	fmt.Printf("nn equals nil: %v\n", nn == nil)

	mm := []int{}
	fmt.Printf("mm equals nil: %v\n", mm == nil)
	fmt.Printf("mm is: %#v\n", mm)

	a, b := []int{1, 2, 3}, []int{1, 2, 108}
	var eq bool = true
	for i, valA := range a {
		if valA != b[i] {
			eq = false
			break
		}
	}

	if eq {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

}
