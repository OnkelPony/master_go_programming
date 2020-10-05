package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text:", text)
	fmt.Println("Input bytes:", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		bytes = scanner.Bytes()
		fmt.Println("You entered:", text)
		fmt.Printf("in bytes:%v\n", bytes)
		if text == "exit" {
			fmt.Println("Exiting the scanning ...")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
