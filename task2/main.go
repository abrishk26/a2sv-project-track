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

func IsPalindrome(input string) bool {
	buf := []rune{}

	for _, c := range input {
		if unicode.IsLetter(c) {
			buf = append(buf, unicode.ToLower(c))
		}
	}

	left, right := 0, len(buf) - 1

	for left < right {
		if buf[left] != buf[right] {
			return false
		}

		left++
		right--
	}

	return true
}
