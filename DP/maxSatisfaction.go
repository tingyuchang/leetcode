package DP

import (
	"sort"
)

func MaxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	if satisfaction[len(satisfaction)-1] < 0 {
		return 0
	}

	res := 0

	l := 0
	r := len(satisfaction) - 1
	minPositive := 0

	for r > l {

		mid := int(uint(l+r) >> 1)
		if satisfaction[mid] < 0 && satisfaction[mid+1] >= 0 {
			minPositive = mid + 1
			break
		}

		if satisfaction[mid] < 0 {
			l = mid
		} else {
			r = mid
		}

	}

	sum := 0
	for i := minPositive; i < len(satisfaction); i++ {
		sum += satisfaction[i]
		res += satisfaction[i] * (i + 1 - minPositive)
	}

	//fmt.Println("sum: ", sum, "minPositive ", minPositive, "res: ", res)
	for i := minPositive - 1; i >= 0; i-- {
		temp := res + sum
		currentNegtiveSum := 0
		for j := minPositive - 1; j >= i; j-- {
			currentNegtiveSum += satisfaction[j]
		}
		//fmt.Println("currentNegtiveSum", currentNegtiveSum)
		temp += currentNegtiveSum
		//fmt.Println("temp", temp)
		if temp > res {
			res = temp
		} else {
			break
		}

	}

	if res < 0 {
		return 0
	}

	return res
}
