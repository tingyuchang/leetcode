package __InterviewBit

import (
	"strings"
)

func reverseString(A string) string {
	strs := strings.Split(A, " ")

	res := ""

	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] == "" {
			continue
		}
		res += strs[i]
		if i > 0 {
			res += " "
		}
	}

	if res[len(res)-1] == ' ' {
		return string(res[:len(res)-1])
	}

	return res
}
