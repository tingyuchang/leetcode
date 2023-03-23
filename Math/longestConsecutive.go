package Math

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	res := 0
	minVal, maxVal := nums[0], nums[0]

	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}

		if v < minVal {
			minVal = v
		}
	}

	count := make(map[int]int)

	for _, v := range nums {
		count[v] = count[v] + 1
	}

	visited := make(map[int]int)
	for k, v := range count {

		if v >= 1 {
			i := k
			temp := 0
			for {
				if visited[i] > 0 {
					temp += visited[i]
					break
				}

				if count[i] >= 1 {
					temp++
				} else {
					break
				}
				i++
			}

			visited[k] = temp

			if temp > res {
				res = temp
			}
		}

	}

	return res
}

/*  counting
if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	res := 0
	minVal, maxVal := nums[0], nums[0]

	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}

		if v < minVal {
			minVal = v
		}
	}

	count := make(map[int]int)

	for _, v := range nums {
		count[v] = count[v] + 1
	}
	i := minVal
	temp := 0
	for i <= maxVal {
		if count[i] >= 1 {
			temp++

			if temp > res {
				res = temp
			}
		} else {
			temp = 0
		}
		i++
	}

	return res
*/
