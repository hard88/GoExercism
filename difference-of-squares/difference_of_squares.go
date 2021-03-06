package diffsquares

func SumOfSquares(input int) int {
	sum := (input * (input + 1) * (2 * input + 1))/6
	return sum
}

func SquareOfSum(input int) int {
	sum := (1 + input) * input / 2
	return sum * sum
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
