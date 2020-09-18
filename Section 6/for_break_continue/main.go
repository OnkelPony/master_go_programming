package main

import "fmt"

func main() {
	for i := 0; i < 666; i++ {
		if i%108 != 0 {
			continue
		}
		fmt.Println(i)
	}

	count := 0
	for j := 0; true; j++ {
		if j%108 == 0 {
			fmt.Printf("%d is divisible by 108\n", j)
			count++
		}
		if count == 666 {
			break
		}
	}
	fmt.Println("Final line.")
}
