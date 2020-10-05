package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	checkErr(err)
	defer file.Close()
	infoByte, err := ioutil.ReadAll(file)
	infoTxt := string(infoByte)
	fmt.Println(infoTxt)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
