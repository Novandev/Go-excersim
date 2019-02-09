package luhn

import (
	"strconv"
	"strings"
	// "reflect"
)

func Valid(str string) bool{
	var err error
	str = strings.Trim(str,"")
	str = strings.Replace(str, " ", "",-1)
	str_array := strings.Split(str,"")
	num_array := make([]int, len(str_array))
	if len(str) < 2 {
		return false
	}
	for pos,char := range str_array{
		if pos % 2 == 0{
			num_array[pos], _ = strconv.Atoi(char)
			num_array[pos] = num_array[pos]*2
			if num_array[pos] > 9{
				num_array[pos] = num_array[pos] -9
			}
			if err != nil {
				return false
			}
		}else{
			num_array[pos], _ = strconv.Atoi(char)
			if num_array[pos] > 9{
				num_array[pos] = num_array[pos] -9
			}
			if err != nil {
				return false
			}
		}

	}
	sum := 0
	for num := range num_array{
		sum += num
	}
	return sum %10 == 0
}