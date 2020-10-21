package main

import (
	"fmt"
	"github.com/OnkelPony/smazat/v3"
	"log"
)

func main() {
	count := 111132
	malas, err := arithmetic.MalaCounts(count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("To achieve %d repetitions you'll need to complete %d malas.", count, malas)
}
