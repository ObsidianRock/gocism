package diffsquares

//SquareOfSum finds the square sum of the fist N natural numbers
func SquareOfSum(num int) int {

	var count int

	for i := 1; i <= num; i++ {
		count += i

	}

	return count * count
}

//SumOfSquares finds the sum of squares of the fist N natural numbers
func SumOfSquares(num int) int {

	var count int
	for i := 1; i <= num; i++ {
		count += i * i
	}

	return count
}

//Difference Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(num int) int {

	return SquareOfSum(num) - SumOfSquares(num)
}
