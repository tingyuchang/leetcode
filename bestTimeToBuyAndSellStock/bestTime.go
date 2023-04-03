package bestTimeToBuyAndSellStock

func MaxProfit2(prices []int) int {
	minVal := prices[0]
	maxVal := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		} else if prices[i]-minVal > maxVal {
			maxVal = prices[i] - minVal
		}
	}

	return maxVal
}

type Stock struct {
	Day   int
	Price int
}

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max, min, previousMin := 0, 0, 0

	for i := 0; i < len(prices); i++ {

		if prices[i] > prices[max] {
			max = i
			previousMin = min
			continue
		}
		if prices[i] < prices[min] {
			// ex: 4,7, 2,1
			// only put min to previousMin when it's smaller than max
			if min < max {
				previousMin = min
			}
			min = i
		} else {
			if prices[i]-prices[min] > prices[max]-prices[previousMin] {
				max = i
				previousMin = min
			}
		}
	}
	//fmt.Println(prices, "max: ", max, "min: ", min, "previous: ", previousMin)
	if min > max {
		if previousMin < max {
			return prices[max] - prices[previousMin]
		}
		return 0
	}
	return prices[max] - prices[min]
}

func MaxProfitSlow(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max, min := 1, 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[min] {
			// check max - min
			if max > i {
				min = i
				continue
			}
			temp := i
			for j := i + 1; j < len(prices); j++ {
				if prices[j] > prices[temp] {
					temp = j
				}
			}
			if prices[temp]-prices[i] > prices[max]-prices[min] {
				min = i
				max = temp
			}
		}
	}

	return prices[max] - prices[min]
}
