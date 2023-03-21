package Math

func ZeroFilledSubarray(nums []int) int64 {
	res := 0
	cache := make(map[int]int)
	sum := func(n int) int {
		result := 0
		for i := 1; i <= n; i++ {
			result += i
		}
		return result
	}

	start := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if start == -1 {
				start = i
			}
			if i == len(nums)-1 {
				v, ok := cache[i-start+1]

				if ok {
					res += v
				} else {
					temp := sum(i - start + 1)
					cache[i-start+1] = temp
					res += temp
				}
			}

		} else {
			if start != -1 {
				v, ok := cache[i-start]

				if ok {
					res += v
				} else {
					temp := sum(i - start)
					cache[i-start] = temp
					res += temp
				}
			}
			start = -1
		}
	}

	return int64(res)
}
