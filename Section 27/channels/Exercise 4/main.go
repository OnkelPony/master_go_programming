package main

import (
	"fmt"
	"time"
)

func power(i int, s chan int) {
	s <- i * i
	time.Sleep(time.Second)
}
func main() {
	c := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, c)
		fmt.Println(<-c)
	}
}
