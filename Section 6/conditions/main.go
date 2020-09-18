package main

import "fmt"

func main() {
	price, inStock := 108, true

	if price >= 80 {
		fmt.Println("Too expensive!")
	}
	if price <= 108 && inStock {
		fmt.Println("Buy it!")
	}

	if price < 108 {
		fmt.Println("It's cheap!")
	} else if price == 108 {
		fmt.Println("On the edge!")
	} else {
		fmt.Println("It's expensive!")
	}
}
