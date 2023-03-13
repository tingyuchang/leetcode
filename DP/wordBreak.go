package DP

import (
	"fmt"
	"strings"
)

func WordBreakII(s string, wordDict []string) []string {
	res := make([]string, 0)
	cache := make(map[string]struct{})
	for _, v := range wordDict {
		cache[v] = struct{}{}
	}
	currentsString := make([]string, 0)
	wordBreakII(s, cache, 0, &res, &currentsString)
	return res
}

func wordBreakII(s string, cache map[string]struct{}, start int, res *[]string, currentString *[]string) {
	if start == len(s) {
		*res = append(*res, strings.Join(*currentString, " "))
		return
	}

	for end := start + 1; end <= len(s); end++ {
		if _, ok := cache[s[start:end]]; ok {
			*currentString = append(*currentString, s[start:end])
			wordBreakII(s, cache, end, res, currentString)
			*currentString = (*currentString)[:len(*currentString)-1]
		}
	}

}

// Dp
func WordBreak(s string, wordDict []string) bool {
	cache := make(map[string]struct{})
	for _, v := range wordDict {
		cache[v] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			_, ok := cache[s[j:i]]
			if dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// BFS
func WordBreakBFS(s string, wordDict []string) bool {
	cache := make(map[string]struct{})
	for _, v := range wordDict {
		cache[v] = struct{}{}
	}
	queue := make([]int, 0)
	visited := make([]bool, len(s))
	queue = append(queue, 0)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:len(queue)]

		if visited[start] {
			fmt.Println(start)
			continue
		}

		for end := start + 1; end <= len(s); end++ {
			if _, ok := cache[s[start:end]]; ok {
				queue = append(queue, end)
				if end == len(s) {
					return true
				}
			}
		}

		visited[start] = true
	}
	return false
}

/*
Time complexity : O(n^3). Size of recursion tree can go up to n^2.
Space complexity : O(n). The depth of recursion tree can go up to n.
 .

Space complexity : O(n)O(n)O(n). The depth of recursion tree can go up to nnn.
*/
func WordBreakMemo(s string, wordDict []string) bool {
	cache := make(map[string]struct{})
	for _, v := range wordDict {
		cache[v] = struct{}{}
	}
	memo := make([]int, len(s))

	for i, _ := range memo {
		memo[i] = 2
	}
	res := checkWithMemo(s, cache, 0, &memo)
	fmt.Println(s, &memo)
	return res
}

func checkWithMemo(s string, cache map[string]struct{}, start int, memo *[]int) bool {
	if len(s) == start {
		return true
	}

	if (*memo)[start] != 2 {
		if (*memo)[start] == 0 {
			return false
		}
		return true
	}

	for end := start + 1; end <= len(s); end++ {
		if _, ok := cache[s[start:end]]; ok {
			if checkWithMemo(s, cache, end, memo) {
				(*memo)[start] = 1
				return true
			}
		}
	}

	(*memo)[start] = 0
	return false
}

func WordBreakSlow(s string, wordDict []string) bool {
	cache := make(map[string]struct{})
	for _, v := range wordDict {
		cache[v] = struct{}{}
	}
	return check(s, cache, 0)
}

func check(s string, cache map[string]struct{}, index int) bool {
	if len(s) == index {
		return true
	}

	for end := index + 1; end <= len(s); end++ {
		if _, ok := cache[s[index:end]]; ok {
			if check(s, cache, end) {
				return true
			}
		}
	}

	return false
}
