package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now().UnixNano() / 1000000

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- "Ahoj!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}

	e := time.Now().UnixNano() / 1_000_000
	fmt.Println("Operation took:", e-s)
}
