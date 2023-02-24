package Math

import "fmt"

func MinimumDeviation(nums []int) int {
	minVal := 0
	// make all odds to even
	for i, v := range nums {
		if v%2 != 0 {
			nums[i] = v * 2
		}

		if i == 0 {
			minVal = nums[i]
		}

		if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	buildHeap(nums)
	maxVal := nums[0]
	deviation := maxVal - minVal
	for {
		if maxVal%2 == 0 {
			temp := maxVal / 2
			nums[0] = temp
			if temp < minVal {
				minVal = temp
			}
			maxHeap(nums, 0)
			maxVal = nums[0]

			if maxVal-minVal < deviation {
				deviation = maxVal - minVal
			}
		} else {
			break
		}
	}
	fmt.Println(nums, "max: ", maxVal, "min: ", minVal, "deviation: ", deviation)
	return deviation
}

func buildHeap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		maxHeap(nums, i)
	}
}

func maxHeap(nums []int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var largest int

	if l < len(nums) && nums[l] > nums[n] {
		largest = l
	} else {
		largest = n
	}

	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != n {
		nums[n], nums[largest] = nums[largest], nums[n]
		maxHeap(nums, largest)
	}
}
