package longestPalindrome

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	ans, start, end := 0, 0, 0

	hashSet := make(map[rune]int)

	for i, v := range s {
		oi, ok := hashSet[v]

		if ok {
			for a := oi; a <= i; a++ {
				if string(s[a]) == string(v) {
					if checkPalindromic(s[a:i+1]) {
						ans = max(ans, i-a+1)
						if ans == i-a+1 {
							start = a
							end = i
						}
					}
				}
			}
		} else {
			hashSet[v] = i
		}
	}

	return s[start:end+1]
}


func max (i, j int) int {
	if i > j {
		return i
	}

	return j
}
func checkPalindromic(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++  {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}
