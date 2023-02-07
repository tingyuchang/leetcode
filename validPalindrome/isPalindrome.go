package validPalindrome

import (
	"strings"
)

/*
non-alphanumeric characters,
0-9: 48-57
A-Z: 65 - 90
a-z: 97-122
*/
func IsPalindrome(s string) bool {
	n := len(s)

	for i, v := range s {
		if i > n/2 {
			break
		}
		if v < 48 || (v > 57 && v < 65) || (v > 90 && v < 97) || v > 122 {
			n++
			continue
		}
		r := s[n-i-1]
		for r < 48 || (r > 57 && r < 65) || (r > 90 && r < 97) || r > 122 {
			n--
			r = s[n-i-1]
		}
		if strings.ToLower(string(v)) != strings.ToLower(string(r)) {
			return false
		}

	}

	return true
}
