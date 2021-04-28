// Package isogram provides functionality related to isogram words
package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not
func IsIsogram(word string) bool {
	counter := make(map[rune]int)
	word = strings.ToLower(word)

	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}

		counter[char]++
		if counter[char] > 1 {
			return false
		}
	}
	return true
}
