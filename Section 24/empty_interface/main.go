package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Golang"
	fmt.Println(empty)

	empty = []int{1, 3, 5, 7}
	fmt.Println(empty)

	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = 108
	you.info = "Golang, learn it!"
	you.info = []int{108, 666}

	fmt.Println(you.info)
}
