package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a, b := 108, 66.6

	// fmt.Println(a * b)
	fmt.Println(float64(a) * b)
	fmt.Println(a * int(b))
	c := int64(b)
	// fmt.Println(a == c)
	fmt.Println(int64(a) == c)

	fmt.Println(string(108))
	fmt.Println(string(666))

	// s1 = string(65.4)
	s1 := fmt.Sprintf("%f", 10.8)
	fmt.Println(s1)

	s2 := fmt.Sprintf("%d", 666)
	fmt.Println(s2)

	var result, err = strconv.ParseFloat("10.8", 32)
	if err != nil {
		log.Println("Can't convert to float32!")
	} else {
		fmt.Printf("type: %T, value: %v\n", result, result)
	}

	i, _ := strconv.Atoi("-666")
	fmt.Printf("type: %T, value: %v", i, i)
	s := strconv.Itoa(108)
	fmt.Printf("type: %T, value: %v", s, s)
}
