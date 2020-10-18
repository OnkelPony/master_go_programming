package main

import "fmt"

func main() {
	c1 := make(chan string)
	go func(s string) {
		c1 <- s
	}("Karmapa")
	fmt.Printf("%s chenno!", <-c1)
}
