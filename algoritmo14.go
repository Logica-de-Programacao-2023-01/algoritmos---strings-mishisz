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

	isDescending := checkDescendingSequence(str)

	if isDescending {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}

func checkDescendingSequence(str string) bool {
	previousNum := 10
	for _, char := range str {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		if num >= previousNum {
			return false
		}
		previousNum = num
	}
	return true
}
