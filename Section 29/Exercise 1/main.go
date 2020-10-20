package main

import (
	"fmt"
	"packages/mypackages/arithmetic"
)

func main() {
	for i := 5; i < 108; i++ {
		if arithmetic.IsPrime(i) {
			fmt.Println(i)
		}
	}
}
