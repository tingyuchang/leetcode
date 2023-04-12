package DP

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	temp := triangle[1]
	temp[0] = temp[0] + triangle[0][0]
	temp[1] = temp[1] + triangle[0][0]
	triangle[1] = temp

	for i := 1; i < len(triangle)-1; i++ {
		temp = triangle[i+1]

		for j := 0; j < len(triangle[i])-1; j++ {
			temp[j+1] = temp[j+1] + min(triangle[i][j], triangle[i][j+1])
		}

		temp[0] = temp[0] + triangle[i][0]
		temp[len(temp)-1] = temp[len(temp)-1] + triangle[i][len(triangle[i])-1]
	}

	res := math.MaxInt

	for _, v := range temp {
		if v < res {
			res = v
		}
	}

	return res

}
