package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	checkErr(err)

	defer file.Close()

	byteSlice := []byte("I learn Golang, pyčo!")
	bytesWritten, err := file.Write(byteSlice)
	checkErr(err)

	fmt.Printf("Bytes written: %d\n", bytesWritten)

	bs := []byte("Golang is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
