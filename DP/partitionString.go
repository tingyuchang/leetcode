package DP

func partitionString(s string) int {
	res := 0
	cache := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		cache[s[i]]++
		if cache[s[i]] >= 2 {
			res++
			cache = make(map[byte]int)
			cache[s[i]]++
		}
	}

	res++

	return res
}
