// Package leap has funtionality to manage leap years
package leap

// IsLeapYear determines if a year is a leap year
func IsLeapYear(year int) (isLeap bool) {
	evenlyDiv := func(n int) bool {
		return year%n == 0
	}

	if evenlyDiv(4) && !evenlyDiv(100) || evenlyDiv(400) {
		isLeap = true
	}

	return isLeap
}
