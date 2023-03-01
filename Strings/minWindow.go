package Strings

import "math"

/*
same idea like 424. Longest Repeating Character Replacement

*/
func MinWindow(s string, t string) string {
	l, r, lengthOfT, minVal, start := 0, 0, len(t), math.MaxInt, 0
	count := make(map[uint8]int, 0)

	for i := 0; i < len(t); i++ {
		count[t[i]]++
	}

	for r < len(s) {
		if count[s[r]] > 0 {
			lengthOfT--
		}
		count[s[r]]--
		r++

		for lengthOfT == 0 {
			if (r - l) < minVal {
				minVal = r - l
				start = l
			}

			count[s[l]]++
			if count[s[l]] > 0 {
				lengthOfT++
			}
			l++
		}
	}

	if minVal != math.MaxInt {
		return s[start : start+minVal]
	}

	return ""
}
