package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, pattern string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	fmt.Print("Digite um padrão: ")
	fmt.Scanln(&pattern)

	indices := findPatternIndices(str, pattern)

	if len(indices) > 0 {
		fmt.Println("Índices do padrão:", indices)
	} else {
		fmt.Println("O padrão não foi encontrado na string.")
	}
}

func findPatternIndices(str, pattern string) []int {
	indices := []int{}
	startIndex := 0
	for {
		index := strings.Index(str[startIndex:], pattern)
		if index == -1 {
			break
		}
		indices = append(indices, startIndex+index)
		startIndex += index + len(pattern)
	}
	return indices
}
