package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)
	var name = "Dan"
	fmt.Printf("My name is: %s\n", name)

	s := "learning golang!"
	fmt.Printf("what i do? %s\n", s)

	car, cost := "Lexus", 666108
	fmt.Printf("%s costs %d\n", car, cost)

	car, year := "Honda", 2015
	fmt.Printf("%s is from year %d\n", car, year)

	var (
		bossName string
		bossAge  int
		gender   string
	)
	fmt.Println(bossName, bossAge, gender)

	var i, j int
	i, j = 108, 666
	i, j = j, i
	fmt.Println("i=", i, "j=", j)
}
