package romannumerals

import (
	"errors"
	"fmt"
)

func ToRomanNumeral(n int) (res string, err error) {
	if n >= 3001 || n < 1 {
		return res, errors.New("Invalid input")
	}

	for n != 0 {
		switch {
		case n >= 1000:
			res += "M"
			n -= 1000
		case n >= 900:
			res += "CM"
			n -= 900
		case n >= 500:
			res += "D"
			n -= 500
		case n >= 400:
			res += "CD"
			n -= 400
		case n >= 100:
			res += "C"
			n -= 100
		case n >= 90:
			res += "XC"
			n -= 90
		case n >= 50:
			res += "L"
			n -= 50
		case n >= 40:
			res += "XL"
			n -= 40
		case n >= 10:
			res += "X"
			n -= 10
		case n >= 9:
			res += "IX"
			n -= 9
		case n >= 8:
			res += "VIII"
			n -= 8
		case n >= 6:
			res += "VI"
			n -= 6
		case n >= 5:
			res += "V"
			n -= 5
		case n >= 4:
			res += "IV"
			n -= 4
		default:
			res += "I"
			n--
		}
	}

	fmt.Println(res)
	return
}
