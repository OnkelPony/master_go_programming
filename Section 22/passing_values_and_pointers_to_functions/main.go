package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 108
	price = 66.6
	name = "Dharma"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 108
	*price = 66.6
	*name = "Dharma"
	*sold = false
}

type Product struct {
	price float64
	name  string
}

func changeProduct(p Product) {
	p.price = 10.8
	p.name = "Ticket"
}

func changeProductByPointer(p *Product) {
	(*p).name = "Ticket"
	p.price = 10.8
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["A"] = 10
	m["B"] = 20
	m["C"] = 30
}

func main() {
	quantity, price, name, sold := 8, 1.9, "Beer", true
	fmt.Println("BEFORE calling changeValues()", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER calling changeValues()", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changeValuesByPointer()", quantity, price, name, sold)

	present := Product{
		price: 35.9,
		name:  "Rum",
	}
	changeProduct(present)
	fmt.Println(present)
	changeProductByPointer(&present)
	fmt.Println("AFTER calling changeProductByPointer()", present)

	prices := []int{10, 20, 30}
	changeSlice(prices)
	fmt.Println("AFTER calling changeSlice()", prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("AFTER calling changeMap()", myMap)

}
