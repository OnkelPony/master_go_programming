package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Value of GOOS", runtime.GOOS)
	fmt.Println("Value of GOARCH", runtime.GOARCH)
	fmt.Println("Value of GOMAXPROCS", runtime.GOMAXPROCS(0))
}
