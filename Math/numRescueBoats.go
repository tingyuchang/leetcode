package Math

import (
	"sort"
)

func NumRescueBoats(people []int, limit int) int {
	return numRescueBoats(people, limit)
}

func numRescueBoats(people []int, limit int) int {
	res := 0
	sort.Sort(sort.Reverse(sort.IntSlice(people)))

	l := 0
	r := len(people) - 1

	for r >= l {
		if people[l] == limit {
			res++
		} else {
			if people[l]+people[r] <= limit {
				r--
			}
			res++
		}
		l++

	}

	return res
}
