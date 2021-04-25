package raindrops

import "strconv"

func Convert(n int) (res string) {
	if n%3 == 0 {
		res += "Pling"
	}
	if n%5 == 0 {
		res += "Plang"
	}
	if n%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		res = strconv.Itoa(n)
	}
	return res
}
