package Sort

import (
	"sort"
)

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	return successfulPairs(spells, potions, success)
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))

	sort.Ints(potions)

	for i := 0; i < len(spells); i++ {
		l := 0
		r := len(potions) - 1

		count := -1

		if spells[i]*potions[0] >= int(success) {
			res[i] = len(potions)
		} else if spells[i]*potions[len(potions)-1] < int(success) {
			res[i] = 0
		} else {
			for r > l {
				mid := int(uint(l+r) >> 1)
				if spells[i]*potions[mid] == int(success) {
					for mid >= 1 {
						if potions[mid-1] == potions[mid] {
							mid--
						} else {
							break
						}
					}
					count = mid

					break
				}

				if spells[i]*potions[mid] < int(success) && spells[i]*potions[mid+1] >= int(success) {
					count = mid + 1
					break
				}

				if spells[i]*potions[mid] > int(success) {
					r = mid
				} else {
					l = mid
				}
			}

			if count != -1 {
				res[i] = len(potions) - count
			}
		}

	}

	return res
}
