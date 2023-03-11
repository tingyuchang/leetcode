package Math

import (
	"fmt"
	"leetcode/DP"
)

func MaxProduct(nums []int) int {
	minVal, maxVal, maxProduct := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if v < 0 {
			minVal, maxVal = maxVal, minVal
		}

		minVal = DP.min(minVal*v, v)
		maxVal = max(maxVal*v, v)

		if maxProduct < maxVal {
			maxProduct = maxVal
		}

	}

	return maxProduct
}

func MaxProductSlow(nums []int) int {
	fmt.Println("Start: ", nums)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]
	singleMax := -11
	// index: value

	currentNegativeCount := make([]int, 0)
	sum := 1
	preZeroIndex := -1
	for i, v := range nums {
		if v > singleMax {
			singleMax = v
		}
		if v == 0 || i == len(nums)-1 {
			if v == 0 {
				v = 1
			}

			if v < 0 {
				currentNegativeCount = append(currentNegativeCount, i)
			}

			if len(currentNegativeCount) == 1 {
				sum1 := product(nums[preZeroIndex+1 : currentNegativeCount[0]])

				if sum1 > max {
					max = sum1
				}

				if currentNegativeCount[0]+1 <= i {
					sum2 := product(nums[currentNegativeCount[0]+1 : i])
					sum2 = sum2 * v
					if sum2 > max {
						max = sum2
					}
				}

			} else if len(currentNegativeCount) >= 3 {
				if len(currentNegativeCount)%2 == 0 {
					sum1 := product(nums[preZeroIndex+1 : i])
					sum1 = sum1 * v
					if sum1 > max {
						max = sum1
					}
				} else {
					sum1 := product(nums[preZeroIndex+1 : currentNegativeCount[len(currentNegativeCount)-1]])
					if sum1 > max {
						max = sum1
					}

					sumLast := product(nums[currentNegativeCount[0]+1 : i])
					sumLast = sumLast * v
					if sumLast > max {
						max = sumLast
					}
				}

			} else {
				// if count == 2 or 0
				if i > preZeroIndex+1 {
					sum = sum * v
					if sum > max {
						max = sum
					}
				}
			}
			currentNegativeCount = make([]int, 0)
			sum = 1
			preZeroIndex = i
		} else if v < 0 {
			currentNegativeCount = append(currentNegativeCount, i)
			sum = sum * v
		} else {
			if sum == 0 {
				sum = v
			} else {
				sum = sum * v
			}
		}

	}

	if singleMax > max {
		return singleMax
	}
	return max
}

func product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 1
	for _, v := range nums {
		sum = sum * v
	}
	return sum
}
