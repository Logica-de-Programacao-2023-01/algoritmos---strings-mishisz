package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	noSpaces := strings.ReplaceAll(input, " ", "")

	fmt.Println("Resultado:", noSpaces)
}
