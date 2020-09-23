package main

import "fmt"

func main() {

	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]
	s3[1] = 108
	fmt.Println(s1, s4)

	a1 := [4]int{10, 20, 30, 40}
	sl1, sl2 := a1[0:2], a1[1:3]
	a1[1] = 666
	fmt.Println(a1, sl1, sl2)

	cars := []string{"Lexus", "Accura", "Mazda", "Subaru"}
	newCars := []string{}

	newCars = append(newCars, cars[1:3]...)
	fmt.Println(newCars)

}
