package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", a+b)
	time.Sleep(time.Millisecond * 108)
	wg.Done()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	const n = 3
	wg.Add(n)
	go sum(1.08, 6.66, &wg)
	go sum(10.8, 6.66, &wg)
	go sum(10.8, 66.6, &wg)
	wg.Wait()
	fmt.Println("Main took:", time.Since(start))
}
