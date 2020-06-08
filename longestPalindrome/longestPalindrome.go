package longestPalindrome

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	ans, start, end := 0, 0, 0

	hashSet := make(map[rune]int)
	hashSet2 := make(map[rune][]int)

	for i, v := range s {
		_, ok := hashSet[v]

		if ok {
			xi := hashSet2[v]

			for _,a := range xi {
				if ans < (i - a +1) {
					if checkPalindromic(s[a:i+1]) {
						ans = max(ans, i-a+1)
						if ans == i-a+1 {
							start = a
							end = i
						}
					}
				}
			}
			hashSet2[v] = append(xi, i)

		} else {
			hashSet[v] = i
			hashSet2[v] = []int{i}
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
