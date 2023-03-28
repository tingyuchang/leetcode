package Sort

func sortColor(nums []int) {
	counter := [3]int{}

	for _, v := range nums {
		counter[v]++
	}

	for i, j := 0, 0; i < len(nums); {
		if counter[j] == 0 {
			j++
			continue
		}

		nums[i] = j
		counter[j]--
		i++
	}
}
