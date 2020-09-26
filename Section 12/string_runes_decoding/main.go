package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1 := 'a'
	fmt.Printf("%T, %d\n", var1, var1)
	word := "píča"
	fmt.Println(len(word))
	fmt.Println(utf8.RuneCountInString(word))
	fmt.Println("Byte (not rune) at position 1:", word[1])

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c", word[i])
	}

	fmt.Println("\n" + strings.Repeat("*", 27))

	for i := 0; i < len(word); {
		r, size := utf8.DecodeRuneInString(word[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("*", 27))

	for i, r := range word {
		fmt.Printf("%d -> %c ", i, r)
	}
}
