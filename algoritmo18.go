package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isNumeric := checkNumeric(input)

	if isNumeric {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

func checkNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
