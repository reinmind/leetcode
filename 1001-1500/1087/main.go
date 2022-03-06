package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	var result []string
	s := strings.Split(text, " ")
	for k, word := range s {

		if word == second && k > 0 && k < len(s)-1 && s[k-1] == first {
			result = append(result, s[k+1])
		}
	}
	return result
}
func main() {
	text := "jkypmsxd jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa kcyxdfnoa jkypmsxd kcyxdfnoa"

	first := "kcyxdfnoa"
	second := "jkypmsxd"
	fmt.Println(findOcurrences(text, first, second))
}
