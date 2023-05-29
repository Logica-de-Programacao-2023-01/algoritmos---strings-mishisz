package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	str := strings.ReplaceAll(input, " ", "")

	isAscending := checkAscendingSequence(str)

	if isAscending {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}

func checkAscendingSequence(str string) bool {
	previousNum := -1
	for _, char := range str {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		if num <= previousNum {
			return false
		}
		previousNum = num
	}
	return true
}
