package Strings

import (
	"strings"
)

func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	l := 0
	temp := 0
	for l < len(s) {
		temp = strings.IndexByte(t, s[l])
		t = string([]byte(t)[temp+1:])
		if temp == -1 {
			return false
		}

		l++
	}
	return true
}

// strings.Index(haystack, needle)
func StrStr(haystack string, needle string) int {
	if needle == haystack {
		return 0
	}

	lengthOfN := len(needle)

	for i, _ := range haystack {
		if i < len(haystack)-lengthOfN+1 {
			s := haystack[i : i+lengthOfN]
			if s == needle {
				return i
			}
		}
	}

	return -1
}

func IndexOfSubString(s, substr string) int {
	n := len(substr)

	switch {
	case n == 0:
		return 0
	case n == 1:
		return strings.IndexByte(s, substr[0])
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	c0 := substr[0]
	c1 := substr[1]
	i := 0
	// sliding window, only compare t times
	t := len(s) - n + 1
	//fails := 0
	for i < t {
		if s[i] != c0 {
			// IndexByte is faster than bytealg.IndexString, so use it as long as
			o := strings.IndexByte(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i+1] == c1 && s[i:i+n] == substr {
			return i
		}
		i++
		//fails++
		//if fails >= 4+i>>4 && i < t {
		//	// See comment in ../bytes/bytes.go.
		//	j := bytealg.IndexRabinKarp(s[i:], substr)
		//	if j < 0 {
		//		return -1
		//	}
		//	return i + j
		//}
	}
	return -1
}
