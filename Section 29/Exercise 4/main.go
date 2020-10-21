package main

import (
	"fmt"
	"github.com/OnkelPony/smazat"
)

func main() {
	for i := 5; i < 108; i++ {
		if arithmetic.IsPrime(i) {
			fmt.Println(i)
		}
	}

	for i := 1; i < 20; i++ {
		fmt.Printf("Faktorial %d is %d\n", i, arithmetic.Factorial(i))
	}

}
