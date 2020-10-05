package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file doesn't exist!")
		}
	}
	err = os.Rename("info.txt", "data.txt")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
