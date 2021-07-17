package diffsquares

import "math"

func SquareOfSum(n int) int {
	var finalSquareSum float64
	var squareSum int
	for i := 1; i <= n; i++ {
		squareSum += i
	}
	finalSquareSum = math.Pow(float64(squareSum), 2.0)

	return int(finalSquareSum)
}

func SumOfSquares(n int) int {
	var tmp float64
	var sumSquare int
	for i := 1; i <= n; i++ {
		tmp = math.Pow(float64(i), 2.0)
		sumSquare += int(tmp)
	}

	return sumSquare
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
