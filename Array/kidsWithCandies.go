package Array

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxVal := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxVal {
			maxVal = candies[i]
		}
	}

	res := make([]bool, len(candies))
	for i, v := range candies {
		res[i] = v+extraCandies >= maxVal
	}

	return res
}
