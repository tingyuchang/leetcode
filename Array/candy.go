package Array

func candy(ratings []int) int {
	// leftToRight := make([]int, len(ratings))
	// rightToLeft := make([]int, len(ratings))
	// leftToRight[0] = 1
	// rightToLeft[len(ratings)-1] = 1

	// for i := 1; i < len(ratings); i++ {
	//     leftToRight[i] = 1
	//     rightToLeft[len(ratings)-1-i] = 1
	//     if ratings[i] > ratings[i-1] {
	//         leftToRight[i] = leftToRight[i-1]+1
	//     }

	//     if ratings[len(ratings)-1-i] >  ratings[len(ratings)-i] {
	//         rightToLeft[len(ratings)-1-i] = rightToLeft[len(ratings)-i]+1
	//     }

	// }
	// sum := 0

	// for i :=0; i < len(leftToRight); i++ {
	//     sum += max(leftToRight[i], rightToLeft[i])
	// }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sum := 0
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(candies); i++ {
		candies[i] = 1
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	for i := 0; i < len(candies); i++ {
		sum += candies[i]
	}

	return sum
}


