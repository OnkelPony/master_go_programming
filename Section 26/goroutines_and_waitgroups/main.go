package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1(goroutine) execution started.")
	for i := 1; i < 4; i++ {
		fmt.Println("i =", i)
		time.Sleep(time.Millisecond * 108)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
}

func f2() {
	fmt.Println("f2 execution started.")
	for i := 4; i < 7; i++ {
		fmt.Println("i =", i)
		time.Sleep(time.Millisecond * 108)
	}
	fmt.Println("f2 execution finished")
}

func main() {
	fmt.Println("main execution started")
	var wg sync.WaitGroup
	wg.Add(1)
	go f1(&wg)
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	f2()
	wg.Wait()
	fmt.Println("main execution finished")
}
