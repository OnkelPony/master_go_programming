package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() {
	for i, val := range n {
		fmt.Println(i, val)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)
	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Println("Seconds in a day:", seconds)
	girls := names{"Lucie", "Natálka", "Cyklistka"}
	girls.print()
	names.print(girls)
}
