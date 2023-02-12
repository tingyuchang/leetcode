package Array

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))

	for _, v := range nums {
		exist := m[v]
		if exist == true {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		exist, ok := m[v]

		if ok {
			// i should always greater than exist
			if (i - exist) <= k {
				return true
			}
		}

		m[v] = i

	}

	return false
}