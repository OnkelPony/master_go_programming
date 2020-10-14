package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const gr = 864000
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gr * 2)
	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Millisecond * 108)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Millisecond * 108)
			m.Lock()
			n--
			m.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("Number of Gorutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("n =", n)
}
