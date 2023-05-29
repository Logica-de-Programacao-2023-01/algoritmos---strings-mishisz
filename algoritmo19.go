package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	reversed := reverseString(input)

	fmt.Println("String invertida:", reversed)
}

func reverseString(str string) string {
	reversed := ""
	for _, char := range str {
		reversed = string(char) + reversed
	}
	return reversed
}
