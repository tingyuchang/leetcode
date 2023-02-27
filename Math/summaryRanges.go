package Math

import "strconv"

func SummaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if i-1 != start {
				result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i-1]))
			} else {
				result = append(result, strconv.Itoa(nums[start]))
			}
			start = i
		}

		if i == len(nums)-1 {
			if i != start {
				result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i]))
			} else {
				result = append(result, strconv.Itoa(nums[i]))
			}
		}
	}

	return result
}
