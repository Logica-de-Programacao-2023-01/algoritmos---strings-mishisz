package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
  isCamelCase, numWords := checkCamelCase(input)
  if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", numWords)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func checkCamelCase(str string) (bool, int) {
	if len(str) == 0 {
		return false, 0
	}

	numWords := 1
	for i, char := range str {
		if i == 0 {
			if !unicode.IsUpper(char) {
				return false, numWords
			}
		} else {
			if unicode.IsUpper(char) {
				numWords++
			}
		}
	}

	return true, numWords
}
