package main

import "fmt"

func main() {
	grades := [3]int{
		1: 108,
		0: 666,
		2: 42,
	}
	fmt.Println(grades)

	accounts := [3]int{
		2: 42,
	}
	fmt.Println(accounts)

	names := [...]string{
		6: "Rocco",
	}
	fmt.Println(names)

	cities := [...]string{
		5: "Tel Aviv",
		"Netanya",
		1: "Jerusalem",
	}
	fmt.Printf("%#v", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Printf("%#v", weekend)
}
