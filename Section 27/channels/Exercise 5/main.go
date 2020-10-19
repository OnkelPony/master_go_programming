package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	for i := 1; i <= 50; i++ {
		go func(i int) {
			c <- i * i
		}(i)
		fmt.Println(<-c)
	}
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
