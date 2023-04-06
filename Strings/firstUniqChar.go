package Strings

import "strings"

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}

func firstUniqChar(s string) int {

	cache := make(map[byte]bool)
	for i := 0; i < len(s); i++ {

		if !cache[s[i]] {
			res := strings.IndexByte(string([]byte(s)[i+1:]), s[i])
			if res == -1 {
				return i
			}
		}
		cache[s[i]] = true

	}

	return -1

}
