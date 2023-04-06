package Strings

import "strings"

func CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	cache := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		cache[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if cache[ransomNote[i]] == 0 {
			return false
		}

		cache[ransomNote[i]]--
	}

	return true
}

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
