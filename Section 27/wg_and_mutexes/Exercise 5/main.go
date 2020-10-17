package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	mx.Lock()
	*b += n
	mx.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mx *sync.Mutex) {
	mx.Lock()
	*b -= n
	mx.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mx)
		go withdraw(&balance, i, &wg, &mx)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
