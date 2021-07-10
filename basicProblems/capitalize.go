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

func capitalize2(str string) string  {
	result := ""

	for i,v := range str {
		if i == 0 {
			result = fmt.Sprintf("%v", strings.ToUpper(string(v)))
			continue
		}
		previousChar := string(str[i-1])
		if previousChar == " "  {
			result = fmt.Sprintf("%v%v", result, strings.ToUpper(string(v)))
		} else {
			result = fmt.Sprintf("%v%v", result, string(v))
		}
	}

	return result
}