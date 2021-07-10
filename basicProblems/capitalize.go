package basicProblems

import (
	"fmt"
	"strings"
)

func capitalize(str string) string  {
	ss := strings.Split(str, " ")
	result := ""
	for i,v := range ss {
		upperStr := fmt.Sprintf("%v%v", strings.ToUpper(string(v[0])), string(v[1:]))
		if i == 0 {
			result = upperStr
		} else  {
			result = fmt.Sprintf("%v %v", result, upperStr)
		}

	}


	return result
}