package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	isAnagram := checkAnagram(str1, str2)

	if isAnagram {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func checkAnagram(str1, str2 string) bool {
	str1 = removeWhitespace(str1)
	str2 = removeWhitespace(str2)

	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	return sortedStr1 == sortedStr2
}

func removeWhitespace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func sortString(str string) string {
	strRunes := []rune(str)
	sort.Slice(strRunes, func(i, j int) bool {
		return strRunes[i] < strRunes[j]
	})
	return string(strRunes)
}
