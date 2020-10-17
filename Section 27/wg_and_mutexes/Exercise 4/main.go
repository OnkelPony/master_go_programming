package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100; i <= 149; i++ {
		go func(a int, wg *sync.WaitGroup) {
			fmt.Printf("Square root of %d is %f\n", a, math.Sqrt(float64(a)))
			time.Sleep(time.Second)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
