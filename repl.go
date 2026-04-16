package main

import (
	"strings"
)

func cleanInput(text string) []string {
	result := []string{}

	// removing spaces
	text = strings.TrimSpace(text)

	// to lower
	text = strings.ToLower(text)

	// splitting the string
	result = strings.Fields(text)

	// fmt.Println(result) // debug

	return result
}
