package diffsquares

func SquareOfSum(num int) int {

	var count int

	for i := 1; i <= num; i++ {
		count += i

	}

	return count * count
}

func SumOfSquares(num int) int {

	var count int
	for i := 1; i <= num; i++ {
		count += i * i
	}

	return count
}

func Difference(num int) int {

	return SquareOfSum(num) - SumOfSquares(num)
}
