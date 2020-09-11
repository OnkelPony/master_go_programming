package main

import "fmt"

func main() {
	name, age := "Metusalem", 666
	fmt.Println(name, "is", age, "years old.")

	a, b, c := 108, 66.6, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s\n", a, b, c)
	fmt.Printf("c is %q\n", c)
	fmt.Printf("grades: %v\n", grades)
	fmt.Printf("grades: %#v\n", grades)
	fmt.Printf("c is of type %T and grades is of type %T\n", c, grades)
	fmt.Printf("The address of a: %p\n", &a)
	fmt.Printf("%c and %c\n", 108, 666)

	const pi float64 = 3.14159265359
	fmt.Printf("PI is %06.2f\n", pi)

	fmt.Printf("108 in base 2 is %b\n", 108)
	fmt.Printf("666 in base 16 is %x\n", 666)
	fmt.Printf("badbeef in base 10 is %d\n", 0xbadbeef)
	s := fmt.Sprintf("a is %d, b is %f, c is %s\n", a, b, c)
	fmt.Println(s)

}
