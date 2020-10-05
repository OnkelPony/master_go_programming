package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("info.txt")
	checkErr(err)
	defer file.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
