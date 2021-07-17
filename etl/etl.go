package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	var score int
	for k, v := range in {
		score = k
		for _, letter := range v {
			out[strings.ToLower(letter)] = score
		}
	}

	return out
}
