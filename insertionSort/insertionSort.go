package insertionSort

func Sort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		for j := i - 1; j >= 0; j-- {
			if list[j] > key {
				list[j+1] = list[j]
				if j == 0 {
					list[j] = key
				}
			} else {
				list[j+1] = key
				break
			}
		}
	}
	return list
}

func SortV2(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}
