package Matrix

import (
	"sort"
)

func FindMinArrowShots(points [][]int) int {
	return findMinArrowShots(points)
}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	current := points[0][1]
	ans := 1

	for i := 1; i < len(points); i++ {
		if current >= points[i][0] {
			continue
		}
		ans++
		current = points[i][1]
	}

	return ans
}
