package singleNumber

func SingleNumber(nums []int) int {
	n := 0
	for _, v := range nums {
		n ^= v
	}
	return n
}

func SingleNumberSlow(nums []int) int {
	m := make(map[int]int, len(nums))

	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			delete(m, v)
		}
	}
	for k, _ := range m {
		return k
	}
	return 0
}
