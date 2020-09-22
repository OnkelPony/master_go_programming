package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]int{}
	var a2 = [3]int{108, 666, 9}
	var a3 = [4]string{"Martin", "Dave", "Andy", "Alan"}
	var a4 = [4]string{"x", "y"}

	var a5 = [...]int{1, 4, 5}
	fmt.Printf("The length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3}

	_, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

	numbers[0] = 7
	fmt.Printf("%v\n", numbers)
	x := numbers[0]
	fmt.Println("x is", x)

	for i, v := range numbers {
		fmt.Println("index:", i, " value:", v)
	}

	fmt.Println(strings.Repeat("%", 27))

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, " value:", numbers[i])
	}

	balances := [2][3]int{
		[3]int{2, 3, 8},
		{11, 12, 20},
	}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	m := [...]int{1, 2, 3}
	n := m
	fmt.Printf("m is equal to n: %v\n", m == n)
	n[1] = 7
	fmt.Printf("m is equal to n: %v\n", m == n)

}
