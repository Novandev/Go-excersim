package grains

import (
	//"fmt"
	"errors"
)
func Square(num_of_squares int)(uint64, error){
	if (num_of_squares > 64 || num_of_squares < 1){
		return 0,errors.New("Incorrect Input")
	}
	total_grains := uint64(1)
	for i := 1; i < num_of_squares; i++{
		total_grains= total_grains * uint64(2)
	}
	return total_grains,nil
}

func Total() uint64 {
	return uint64(18446744073709551615)
}