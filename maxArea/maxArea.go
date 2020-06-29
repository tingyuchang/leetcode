package maxArea

func MaxArea(height []int) int {
	var area int
	start := 0
	end := len(height) - 1
	for {

		startValue := height[start]
		endValue := height[end]

		area = max(area, (end-start)*min(startValue, endValue))

		if end-start == 1 {
			break
		}

		if startValue > endValue {
			end--
		} else {
			start++
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
