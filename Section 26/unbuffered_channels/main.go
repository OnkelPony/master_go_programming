package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func(c chan int) {
		fmt.Println("func goroutine start sending")
		c <- 108
		fmt.Println("func goroutine after sending")
	}(c1)

	fmt.Println("main goroutine sleeps 2s")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine start receiving")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	time.Sleep(time.Second)

}
