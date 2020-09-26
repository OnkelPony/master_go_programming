package main

import "fmt"

func main() {
	s1 := "Quick brown fox"
	fmt.Println(s1[0:3])
	fmt.Printf("%q\n", s1[2:6])

	s2 := "Příliš žluťoučký kůň"
	fmt.Printf("%q\n", s2[2:6])
	rs2 := []rune(s2)
	fmt.Printf("%q\n", string(rs2[2:6]))

}
