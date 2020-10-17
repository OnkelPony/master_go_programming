package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a float64, wg *sync.WaitGroup) {
		fmt.Printf("Square root of %f is %f\n", a, math.Sqrt(a))
		wg.Done()
	}(108, &wg)
	wg.Wait()
}
