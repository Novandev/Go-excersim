package diffsquares


func SumOfSquares(num int)int{
	sumOfSquares := 0
	for i:=1; i < num + 1; i++{

		sumOfSquares += (i*i)
	}
	return sumOfSquares
}
func SquareOfSum(num int) int{
	squareOfSums := 0
	for i:=1; i < num + 1; i++{
		squareOfSums += i
	}
	squareOfSums *=squareOfSums
	return squareOfSums
}

func Difference(num int)int{

	return  SquareOfSum(num) - SumOfSquares(num)
}