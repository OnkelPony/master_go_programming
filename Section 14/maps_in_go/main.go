package main

import (
	"fmt"
)

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)
	fmt.Printf("No. of elements: %d\n", len(employees))
	fmt.Printf("Value of key %q is %q\n", "Dan", employees["Dan"])

	people := map[string]float64{}
	people["John"] = 10.8
	people["Marry"] = 66.6
	fmt.Printf("%v\n", people)

	map1 := make(map[int]int)
	fmt.Printf("map1: %#v\n", map1)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 10.8,
		"EUR": 66.6,
		"CHF": 6.66,
	}
	fmt.Println(balances)

	m := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m

	balances["USD"] = 666
	balances["GBP"] = 108
	v, ok := balances["CZK"]

	if ok {
		fmt.Println("The CZK balance is ", v)
	} else {
		fmt.Println("The CZK is not in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}
	fmt.Printf("balances: %v\n", balances)

	delete(balances, "USD")

	a := map[string]string{"A": "666", "C": "Z", "B": "Y"}
	b := map[string]string{"A": "666", "B": "Y", "C": "Z"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	if s1 == s2 {
		fmt.Println("Maps are equal!")
	} else {
		fmt.Println("Maps are not equal!")
	}

	friends := map[string]int{"Dan": 40, "Maria": 30}
	neighbors := friends
	friends["Dan"] = 666
	fmt.Println(neighbors)

	colleagues := make(map[string]int)
	for k, v := range friends {
		colleagues[k] = v
	}
}
