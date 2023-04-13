package Math

import (
	"sort"
)

func SmallestRangeII(nums []int, k int) int {
	return smallestRangeII(nums, k)
}

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)

	minVal := nums[0] + k
	maxVal := nums[len(nums)-1] - k
	res := nums[len(nums)-1] - nums[0]

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums)-1; i++ {
		// needs to think again
		tmpMin := min(minVal, nums[i+1]-k)
		tmpMax := max(maxVal, nums[i]+k)
		res = min(res, tmpMax-tmpMin)
	}

	return res
}

/*
sort.Ints(nums)

	minVal := nums[0] + k
	maxVal := nums[len(nums)-1] - k
	defaultRes := nums[len(nums)-1] - nums[0]

	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}
	if defaultRes <= maxVal-minVal {
		return defaultRes
	}

	// fmt.Printf("default: %d, %d\n", minVal, maxVal)
	for i := 0; i < len(nums); i++ {

		if nums[i] <= minVal {
			if nums[i]+k > maxVal {
				maxVal = nums[i] + k
			}
		} else if nums[i] >= maxVal {
			if nums[i]-k < minVal {
				minVal = nums[i] - k
			}
		} else {
			// mid element
			// check it add or minus k , which one is smaller
			tmpMin := nums[i] - k
			tmpMax := nums[i] + k

			if tmpMin >= minVal || tmpMax <= maxVal {
				//continue
			} else {
				if maxVal-tmpMin > tmpMax-minVal {
					maxVal = tmpMax
				} else {
					minVal = tmpMin
				}
			}
		}

		// fmt.Printf("nums: %d, min: %d max: %d\n", nums[i], minVal, maxVal)

	}


	if maxVal-minVal > defaultRes {
		return defaultRes
	}

	return maxVal - minVal
*/
