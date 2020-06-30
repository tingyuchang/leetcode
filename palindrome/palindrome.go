package palindrome

import (
	"strconv"
)

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	n := len(s)

	for i := 0; i < n/2; i++ {
		j := n - 1 - i

		if s[i] != s[j] {
			return false
		}
	}

	return true
}
