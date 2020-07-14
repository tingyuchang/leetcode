package generateParenthesis

import "fmt"

var tmp []string
var N int
var count int
func GenerateParenthesis(n int) []string {
	N = n
	count = 1
	tmp = []string{}
	backtrack("", 0, 0)
	return tmp
}

func backtrack (s string, l, r int){
	fmt.Printf("Count: %v\tS: %v L: %v R: %v\n", count, s, l , r)
	count++
	if len(s) == 2* N {
		tmp = append(tmp, s)
		return
	}
	if l < N {
		backtrack(fmt.Sprintf("%v(", s), l+1, r)
	}
	if r < l {
		fmt.Printf("-------\n")
		backtrack(fmt.Sprintf("%v)", s), l, r+1)
	}

}