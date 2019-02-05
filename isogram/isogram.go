
package isogram

import (
	"sort"
	"strings"
)

func IsIsogram( str string) bool{
	found:= true
	s := strings.Split(str, "")
	sort.Strings(s)
	for i:=0; i < len(s) -1; i++{
		if strings.ToLower(s[i]) == strings.ToLower(s[i+1]) {
			found = false
		}
	}
	return found
}