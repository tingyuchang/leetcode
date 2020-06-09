package longestPalindrome

/*
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
 */

func checkPalindromic(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++  {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}
func max (i, j int) int {
	if i > j {
		return i
	}

	return j
}


func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0

	for i := 0; i< len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)  // if center is between 2 nodes
		ans := max(len1, len2)
		if ans > end - start {
			start = i - (ans-1)/2
			end = i + ans/2
		}
 	}


	return s[start:end+1]
}

func expandAroundCenter(s string, l, r int ) int {
	for l >=0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}
