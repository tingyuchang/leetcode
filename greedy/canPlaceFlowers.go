package greedy

func CanPlaceFlowers(flowerbed []int, n int) bool {

	i := 0

	for i < len(flowerbed) {
		v := flowerbed[i]

		if v == 1 {
			i++
			continue
		}

		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i += 2
		} else {
			i++
		}
		if n == 0 {
			return true
		}

	}

	return false
}
