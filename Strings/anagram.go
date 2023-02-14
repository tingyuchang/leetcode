package Strings

func IsAnagram(s string, t string) bool {
	ms := make(map[int32]int, 0)

	for _, v := range s {
		val, ok := ms[v]
		if !ok {
			ms[v] = 1
		} else {
			ms[v] = val + 1
		}
	}

	for _, v := range t {
		val, ok := ms[v]

		if !ok {
			return false
		} else {
			ms[v] = val - 1
		}
	}

	for _, v := range ms {
		if v != 0 {
			return false
		}
	}

	return true
}
