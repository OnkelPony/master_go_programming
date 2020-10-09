package main

import "fmt"

func main() {
	name := "Aletta"
	fmt.Println(&name)

	var x int = 108
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr)

	var ptr1 *float32
	_ = ptr1

	p := new(int)
	x = 100
	p = &x

	fmt.Printf("p is of type %T with value %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x)

	*p = 90
	fmt.Println(*p)
	fmt.Println("*p == x:", *p == x)
	fmt.Println("Value of x:", *p, "Value of *p:", x)

	*p = 10
	*p /= 2
	fmt.Println(x)
}
