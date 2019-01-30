
package isogram

func isIsogram( str string) bool{
	isWordIsogram := false
	var bagOfWords = []string
	for i := range str{
		for x := range bagOfWords{
			if string(str[i]) == bagOfWords[x]{
				isWordIsogram = true
			}
		}

	}
	return isWordIsogram
}