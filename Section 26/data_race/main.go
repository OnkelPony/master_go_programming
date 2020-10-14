package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 108
	var wg sync.WaitGroup
	wg.Add(gr * 2)
	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Millisecond * 108)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Millisecond * 108)
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("n =", n)
}
