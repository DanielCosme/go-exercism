package luhn

import (
	"strconv"
	"strings"
)

func Valid(in string) bool {
	var sum, toDouble int64
	in = strings.ReplaceAll(in, " ", "")
	if len(in) <= 1 {
		return false
	}

	for i := len(in) - 1; i >= 0; i-- {
		var digit int64
		char := in[i]

		digit, err := strconv.ParseInt(string(char), 10, 64)
		if err != nil {
			return false
		}

		if toDouble%2 == 1 {
			digit = digit + digit
			if digit > 9 {
				digit = digit - 9
			}
		}

		sum += digit
		toDouble++
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
