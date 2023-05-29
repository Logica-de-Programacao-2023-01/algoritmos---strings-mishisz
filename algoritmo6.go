package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	words := strings.Fields(input)
	count := len(words)

	fmt.Println("A string contém", count, "palavra(s).")
}
