package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	var oldChar, newChar string
	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite o caractere substituto: ")
	fmt.Scanln(&newChar)

	result := strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", result)
}
