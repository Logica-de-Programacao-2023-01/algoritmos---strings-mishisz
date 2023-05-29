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
	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite a letra substituta: ")
	fmt.Scanln(&newChar)

	result := strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", result)
}
