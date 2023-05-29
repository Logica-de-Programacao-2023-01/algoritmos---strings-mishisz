package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	uppercase := strings.ToUpper(input)

	fmt.Println("Resultado:", uppercase)
}
