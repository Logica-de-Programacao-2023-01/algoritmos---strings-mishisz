package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	uniqueLetters := findUniqueLetters(input)

	fmt.Println("Letras únicas:", uniqueLetters)
}

func findUniqueLetters(str string) string {
	unique := ""
	for _, char := range str {
		if strings.Count(str, string(char)) == 1 {
			unique += string(char)
		}
	}
	return unique
}
