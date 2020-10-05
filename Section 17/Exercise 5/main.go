package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	checkErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Println(lines)
	}
	err = scanner.Err()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
