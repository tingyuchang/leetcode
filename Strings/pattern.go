package Strings

import (
	"strings"
)

func RepeatedSubstringPattern(s string) bool {
	i := 0
	n := len(s)

	switch {
	case n == 0:
		return false
	case n == 1:
		return false
	}

	c0 := s[i]

	for i < n {
		start := 0
		next := strings.IndexByte(s[i+1:], c0)
		if next == -1 {
			return false
		}
		next += i + 1
		ok := false
		length := next - start

		for next+length <= n {
			if s[start:next] != s[next:next+length] {
				break
			}
			start = next
			next = next + length

			if next == n {
				ok = true
			}
		}

		if ok {
			return true
		}
		i++
	}

	return false
}

func WordPattern(pattern string, s string) bool {
	// ex: a : dog
	// b: cat
	patternString := make(map[uint8]string, 0)
	subStrings := make([]string, 0)

	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if start != i {
				subStrings = append(subStrings, s[start:i])
			}
			start = i + 1
		}

		if i == len(s)-1 {
			subStrings = append(subStrings, s[start:i+1])
		}
	}

	if len(subStrings) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		v, ok := patternString[pattern[i]]
		subString := subStrings[i]

		if !ok {
			patternString[pattern[i]] = subString
		} else {
			if v != subString {
				return false
			}
		}
	}

	// check
	check := make(map[string]uint8, 0)
	for k, v := range patternString {
		k2, ok := check[v]
		if !ok {
			check[v] = k
		} else {
			if k != k2 {
				return false
			}
		}
	}

	return true
}
