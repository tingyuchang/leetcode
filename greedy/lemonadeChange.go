package greedy

func lemonadeChange(bills []int) bool {
	sum := make(map[int]int)
	exchanges := []int{5, 10, 20}

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			sum[5]++
		} else {

			exchange := bills[i] - 5

			current := len(exchanges) - 1

			for current >= 0 {
				if exchanges[current] > exchange {
					current--
					continue
				}

				if sum[exchanges[current]] > 0 {
					exchange -= exchanges[current]
					sum[exchanges[current]]--

					if exchange == 0 {
						break
					}
				} else {
					current--
				}

			}

			if exchange != 0 {
				return false
			}
			sum[bills[i]]++
		}
	}

	return true

}
