package DP

type sortString []byte

func (s sortString) Len() int {
	return len(s)
}

func (s sortString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func IsScramble(s1, s2 string) bool {
	cache := make(map[string]bool)
	return isScrambleHelper(s1, s2, cache)
}

func isScrambleHelper(s1 string, s2 string, cache map[string]bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	key := s1 + "-" + s2
	if val, ok := cache[key]; ok {
		return val
	}
	count := [26]int{}
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			cache[key] = false
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScrambleHelper(s1[:i], s2[:i], cache) && isScrambleHelper(s1[i:], s2[i:], cache) {
			cache[key] = true
			return true
		}
		if isScrambleHelper(s1[:i], s2[len(s2)-i:], cache) && isScrambleHelper(s1[i:], s2[:len(s2)-i], cache) {
			cache[key] = true
			return true
		}
	}
	cache[key] = false
	return false
}
