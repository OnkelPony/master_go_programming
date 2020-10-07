package main

import (
	"fmt"
	"strings"
)

func main() {
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Miller",
		age:       27,
	}
	fmt.Printf("%#v\n", diana)

	type Book struct {
		string
		float32
		bool
	}
	b1 := Book{"Exodus by Leon Uris", 10.8, true}
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}
	e1 := Employee{
		name:   "John",
		salary: 108666,
		bool:   false,
	}
	fmt.Printf("%#v\n", e1)
	e1.bool = true

	fmt.Println(strings.Repeat("*", 27))

	type Contact struct {
		email, address string
		phone          int
	}

	type Customer struct {
		name        string
		turnover    int
		contactInfo Contact
	}
	john := Customer{
		name:     "John Holmes",
		turnover: 666108,
		contactInfo: Contact{
			email:   "john@holmes.cock",
			address: "Filthy boulevard 3",
			phone:   16549551,
		},
	}
	fmt.Printf("%#v\n", john)
	fmt.Printf("%+v\n", john)
	fmt.Println("john's turnover:", john.turnover)
	fmt.Println("john's email:", john.contactInfo.email)
	john.contactInfo.email = "mrdaak@gmail.com"
	fmt.Println("john's new email is:", john.contactInfo.email)
}
