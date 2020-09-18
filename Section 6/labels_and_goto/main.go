package main

import "fmt"

func main() {

	outer := 666
	_ = outer

	people := [...]string{"Jan", "Petr", "Kamila", "venca", "Růžena"}
	friends := [2]string{"Kamila", "Lojza"}
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found friend %s at index %d", friend, index)
				break outer
			}
		}
	}

	i := 0

loop:
	if i < 6 {
		fmt.Println(i)
		i++
		goto loop
	}
}
