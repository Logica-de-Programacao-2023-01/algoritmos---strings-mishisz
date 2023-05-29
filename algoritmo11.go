package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	noVowels := removeVowels(input)

	fmt.Println("Resultado:", noVowels)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	noVowels := strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return -1
		}
		return r
	}, str)
	return noVowels
}
