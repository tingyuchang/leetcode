package lengthOfLongestSubstring


func LengthOfLongestSubstring(s string) int {
	hashSet := make(map[rune]int)

	ans, start := 0, 0

	for i, v := range s {
		_, ok := hashSet[v]

		if ok {
			start = max(start, hashSet[v]+1)
		}

		hashSet[v] = i

		ans = max(ans, i-start+1)

	}

	return ans
}

func max(n int, m int) int {
	if n > m {
		return n
	}
	return m
}