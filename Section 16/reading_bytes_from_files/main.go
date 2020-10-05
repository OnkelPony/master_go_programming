package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	checkErr(err)
	defer file.Close()

	byteSlice := make([]byte, 2)
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	checkErr(err)
	fmt.Printf("Number of bytes read: %d\n", numberBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
	fmt.Println(strings.Repeat("#", 27))

	file, err = os.Open("main.go")
	checkErr(err)

	data, err := ioutil.ReadAll(file)
	checkErr(err)

	fmt.Printf("Data as a string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	data, err = ioutil.ReadFile("test.txt")
	checkErr(err)
	log.Printf("Data read: %s\n", data)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
