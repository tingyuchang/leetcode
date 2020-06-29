package maxArea

func MaxArea(height []int) int {
	var area int

	for i, v := range height {
		for j := i + 1; j < len(height); j++ {
			k := height[j]
			tempArea := (j - i) * min(v, k)
			area = max(area, tempArea)
		}
	}

	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
