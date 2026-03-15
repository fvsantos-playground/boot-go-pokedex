package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	result := make([]string, 0, len(words))
	for _, v := range words {
		v = strings.TrimSpace(v)
		if v != "" {
			result = append(result, strings.ToLower(v))
		}
	}

	return result
}
