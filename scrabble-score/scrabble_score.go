//  Package scrabble has functionality for the game Scrabble
package scrabble

import "strings"

// Score computes the Scrabble score for a given word
func Score(word string) (score int) {
	word = strings.ToLower(word)
	for _, char := range word {
		switch {
		case strings.ContainsRune("aeioulnrst", char):
			score++
		case strings.ContainsRune("dg", char):
			score += 2
		case strings.ContainsRune("bcmp", char):
			score += 3
		case strings.ContainsRune("fhvwy", char):
			score += 4
		case strings.ContainsRune("k", char):
			score += 5
		case strings.ContainsRune("jx", char):
			score += 8
		case strings.ContainsRune("qz", char):
			score += 10
		}
	}
	return score
}
