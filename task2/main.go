package task2

import (
	"unicode"
	"strings"
)


func CountWordFreq(input string) map[string]int  {
	count := map[string]int{}
	fields := strings.Fields(input)

	for _, field := range fields {
		word := []rune{}

		for _, r := range field {
			if unicode.IsLetter(r) {
				word = append(word, unicode.ToLower(r))
			}
		}

		count[string(word)]++
	}
	
	return count
}
