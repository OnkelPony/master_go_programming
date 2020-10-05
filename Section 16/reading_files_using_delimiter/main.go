package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	checkErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords)
	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
		checkErr(err)
		log.Println("Scan was completed and it reached End Of File")

	}
	fmt.Println("First Line found:", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
