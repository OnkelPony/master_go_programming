package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year, pages   int
	}

	lastBook := book{"Božská komedie", "Dante Alighieri", 1320}
	fmt.Println(lastBook)

	bestBook := book{author: "George Orwell", title: "Farma zvířat", year: 1945}
	_ = bestBook

	aBook := book{title: "Náhodná"}
	fmt.Printf("%#v\n", aBook)

	fmt.Println(lastBook.title)

	lastBook.author = "nejlepší"
	lastBook.year = 2020
	fmt.Printf("lastBook: %+v\n", lastBook)

	randomBook := book{title: "Random", author: "Anonymus", year: 1}
	fmt.Println(lastBook == randomBook)

	myBook := randomBook
	myBook.year = 108
	fmt.Println(myBook, randomBook)

}
