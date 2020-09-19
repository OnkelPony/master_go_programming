package main

import "fmt"

func main() {
	language := "c++"

	switch language {
	case "Go", "golang":
		fmt.Println("You're learning Go language!")
	case "python":
		fmt.Println("You're learning Python")
	default:
		fmt.Println("You're learning other language.")
	}

	n := -108
	switch {
	case n%2 == 0:
		fmt.Println("Even")
	case n%2 != 0:
		fmt.Println("Odd")
	default:
		fmt.Println("Miracle happened")

	}

	switch q := -0; true {
	case q < 0:
		fmt.Println("Negative")
	case q > 0:
		fmt.Println("Positive!")
	default:
		fmt.Println("Zero")
	}
}
