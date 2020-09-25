package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Printf("%s\n", s1)
	fmt.Printf("%q\n", s1)

	fmt.Println("He says \"Hello!\"")
	fmt.Println(`He says "Hello!"`)

	fmt.Println("Price: 108\nBrand: Asics")
	fmt.Println(`Price: 666
Brand: 4SR`)

	fmt.Println(`C:\Users\někdo`)
	fmt.Println("C:\\Users\\někdo")

	var s3 = "I love" + " Go " + "programming"
	fmt.Println(s3 + "!")
	fmt.Println("Element at index zero:", string(s3[0]))
}
