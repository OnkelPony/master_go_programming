package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 3)
	fmt.Println("Channel capacity:", cap(c1))

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending\n", i)
		}
		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps for 2s")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine received value from channel", v)
	}
	fmt.Println(<-c1)
	//c1 <- 108
}
