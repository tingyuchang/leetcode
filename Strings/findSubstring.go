package Strings

import "fmt"

func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)

	for _, v := range words {
		wordsMap[v]++
	}

	ans := make([]int, 0)

	m := len(words[0])
	n := len(words)
	for i := 0; i <= len(s)-m*n; i++ {
		if _, ok := wordsMap[s[i:i+m]]; ok {
			wordsMap[s[i:i+m]]--
			start := i + m
			for j := 1; j < n; j++ {
				if used, ok := wordsMap[s[start:start+m]]; ok && used > 0 {
					wordsMap[s[start:start+m]]--
					start = start + m
				} else {
					break
				}
			}

			// check wordsMap
			fmt.Println(wordsMap)
			ok := true
			for k, v := range wordsMap {
				if v != 0 {
					ok = false
				}
				wordsMap[k] = 0
			}
			for _, v := range words {
				wordsMap[v]++
			}

			if ok {
				ans = append(ans, i)
			}

		}
	}

	return ans

}
