package Strings

func IsIsomorphic(s, t string) bool {

	if len(s) != len(t) {
		return false
	}
	m := make(map[uint8]int, len(s))
	m2 := make(map[uint8]int, len(s))

	for i := 0; i < len(s); i++ {
		s1, ok1 := m[s[i]]
		t1, ok2 := m2[t[i]]
		if (ok1 && ok2 && s1 != t1) || ok1 != ok2 {
			return false
		}
		m[s[i]] = i
		m2[t[i]] = i
	}

	return true
}
