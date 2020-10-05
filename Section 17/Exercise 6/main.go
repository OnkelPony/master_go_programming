package main

import (
	"io/ioutil"
	"log"
)

func main() {
	sentence := []byte("The Go gopher is an iconic mascot! A není to kůň!")
	err := ioutil.WriteFile("info.txt", sentence, 0644)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
