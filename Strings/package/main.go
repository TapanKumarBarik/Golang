package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]

	fmt.Println(s)

	s = s + strings.Repeat("!", len(s))

	fmt.Println(s)

	s = strings.ToUpper(s)

	fmt.Println(s)

}
