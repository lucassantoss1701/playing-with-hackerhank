package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func splitCamelCase(inputStr string) string {
	var words []string
	var currentWord string

	for _, char := range inputStr {
		if unicode.IsUpper(char) {
			if currentWord != "" {
				words = append(words, strings.ToLower(currentWord))
			}
			currentWord = string(char)
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append(words, strings.ToLower(currentWord))
	}

	return strings.Join(words, " ")
}

func combineCamelCase(words []string, isMethod bool) string {
	var result strings.Builder

	for _, word := range words {
		result.WriteString(cases.Title(language.Und, cases.NoLower).String(word))
	}

	if isMethod {
		result.WriteString("()")
	}

	return result.String()
}

func camelCaseOperations(inputLines []string) {
	for _, line := range inputLines {
		parts := strings.Split(line, ";")
		operation, type_, data := parts[0], parts[1], parts[2]

		if operation == "S" {
			result := splitCamelCase(data)
			fmt.Println(result)
		} else if operation == "C" {
			isMethod := type_ == "M"
			result := combineCamelCase(strings.Fields(data), isMethod)
			fmt.Println(result)
		}
	}
}

func main() {
	// Sample Input
	inputLines := []string{
		"S;M;plasticCup()",
		"C;V;mobile phone",
		"C;C;coffee machine",
		"S;C;LargeSoftwareBook",
		"C;M;white sheet of paper",
		"S;V;pictureFrame",
	}

	// Execute the program with the sample input
	camelCaseOperations(inputLines)
}
