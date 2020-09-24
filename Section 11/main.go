package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]
	no := len(numbers)
	if no < 2 || no > 10 {
		log.Fatal("Enter between 2 and 10 numbers!")
	}
	product, sum := 1., 0.
	for _, numStr := range numbers {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
		product *= num

	}
	fmt.Printf("Sum: %f, Product: %f\n", sum, product)
}
