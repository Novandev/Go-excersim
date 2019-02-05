package diffsquares

func SquareOfSum(num int) int{
	sumOfSquares := 0
	squareOfSums := 0
	diff := 0
	for i:=1; i < num + 1; i++{
		squareOfSums += i
		sumOfSquares += (i*i)
	}
	squareOfSums *=squareOfSums
	diff =(squareOfSums-sumOfSquares)
	return diff
}
