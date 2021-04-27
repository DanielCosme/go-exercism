//  Package scrabble has functionality for the game Scrabble
package scrabble

import "strings"

// Score computes the Scrabble score for a given word
func Score(word string) (score int) {
	word = strings.ToUpper(word)
	for _, char := range word {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}
