package Sort

import (
	"fmt"
)

func MoveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			count++
			continue
		}
		if count != 0 {
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
	// insertion sort's thinking
	/*
		for i := len(nums) - 1; i >= 0; i-- {
				key := nums[i]
				if key != 0 {
					continue
				}
				for j := i + 1; j < len(nums); j++ {
					if nums[j] != 0 {
						nums[j-1] = nums[j]
						if j == len(nums)-1 {
							nums[j] = key
						}
					} else {
						nums[j-1] = key
						break
					}
				}
			}
	*/
}

func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		for j := i - 1; j >= 0; j-- {
			if list[j] > key {
				list[j+1] = list[j]
				if j == 0 {
					list[j] = key
				}
			} else {
				list[j+1] = key
				break
			}
		}
	}
	return list
}

func InsertionSortV2(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}

func InsertionSortDesc(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && key > list[j] {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}

// Search return index of target in input list, if not found, return error
func LinearSearch(list []int, target int) (int, error) {
	for i, v := range list {
		if v == target {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%d was Not found", target)
}

// BinaryAdd implements n emelemts list
// A = Σa[i]*2^i
// B = Σb[i]*2^i
// c = a + b = Σc[i]*2^i
func AddBinary(a, b []int) []int {
	c := make([]int, len(a)+1)
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + carry
		switch sum {
		case 2:
			carry = 1
			c[i+1] = 0
		case 3:
			carry = 1
			c[i+1] = 1
		default:
			carry = 0
			c[i+1] = sum
		}
	}

	if carry == 1 {
		c[0] = 1
	}
	return c
}

// sum from  0 to n in list
func SumList(list []int, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += list[i]
	}
	return sum
}
