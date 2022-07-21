package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "heyy711@"

	//calculate the byte
	fmt.Println(len(s))

	//calculate the length
	fmt.Println(utf8.RuneCountInString(s))
}
