package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("Aletta is anal queen!", "anal")
	p(result)

	result = strings.ContainsAny("aůoerifjaq", "sy")
	p(result)

	result = strings.ContainsRune("Heather is oral queen, tečka!", 'č')
	p(result)

	n := strings.Count("Ó, fretná chrochtobuznosti...", "o")
	p(n)

	n = strings.Count("pět", "")
	p(n)

	p(strings.ToLower("Go Python Java"))
	p(strings.ToUpper("Go Python Java"))

	p(strings.EqualFold("Praseté!", "pRaSeTé!"))

	p(strings.Repeat("čaPí", 27))

	myStr := strings.Replace("192.168.108.9", ".", ":", 2)
	p(myStr)

	myStr = strings.Replace("192.168.108.9", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("befelemepesseveze", "e", "3")
	p(myStr)

	myval := strings.Split("befelemepesseveze", "e")
	fmt.Printf("%T, %#v\n", myval, myval)

	myval = strings.Split("befelemepesseveze", "")
	fmt.Printf("%T, %#v\n", myval, myval)

	myval = []string{"I", "learn", "Golang"}
	myStr = strings.Join(myval, "*")
	p(myStr)

	myStr = "Orange\tGreen\nBlue Yellow"
	myval = strings.Fields(myStr)
	fmt.Printf("%T, %#v\n", myval, myval)

	myStr = strings.TrimSpace("\n Goodbye windows, welcome Linux!\t ")
	fmt.Printf("%q\n", myStr)

	myStr = strings.Trim("... Hello, Gophers?!!", ".,?! ")
	fmt.Printf("%q\n", myStr)

}
