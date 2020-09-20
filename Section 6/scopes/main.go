package main

import "fmt"
import f "fmt"

const done = false

func main() {
	var task = "Running!"
	fmt.Println(task, done)
	f.Println("Bajbaj!")

	const done = true
	fmt.Printf("done in main(): %v\n", done)
	f1()
}

func f1() {
	fmt.Printf("done in f1(): %v\n", done)
}
