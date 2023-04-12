package __InterviewBit

import "fmt"

func serialize(A []string) string {
	res := ""

	for i := 0; i < len(A); i++ {
		res = res + A[i] + fmt.Sprintf("%d~", len(A[i]))
	}

	return res

}
