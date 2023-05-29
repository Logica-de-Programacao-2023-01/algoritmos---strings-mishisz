package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := replaceVowels(input, '*')

	fmt.Println("Resultado:", result)
}

func replaceVowels(str string, replacement rune) string {
	vowels := "aeiouAEIOU"
	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, string(vowel), string(replacement))
	}
	return str
}
