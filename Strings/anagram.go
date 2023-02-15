package Strings

import "strconv"

func GroupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	// a1e1t1 / a1b1t1 / a1n1t1
	m := make(map[string][]string, 0)
	for _, s := range strs {
		code := anagramCode(s)

		val, ok := m[code]

		if !ok {
			m[code] = []string{s}
			continue
		}
		val = append(val, s)
		m[code] = val
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func anagramCode(s string) string {
	// generate code like:
	// eat -> a1e1t1
	// cool -. c1l1o2
	var code string
	m := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	for _, v := range s {
		count, ok := m[string(v)]
		if ok {
			m[string(v)] = count + 1
		}
	}
	// a-z => 97 - 122
	for i := 97; i < 123; i++ {
		count := m[string(i)]
		if count != 0 {
			code = code + string(i) + strconv.Itoa(count)
		}
	}

	return code
}

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
