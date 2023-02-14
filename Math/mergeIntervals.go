package Math

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	result := make([][]int, 0)
	// star:end
	mStartToEnd := make(map[int]int)
	starts := make([]int, 0)
	for _, v := range intervals {
		_, ok := mStartToEnd[v[0]]
		if ok {
			if mStartToEnd[v[0]] < v[1] {
				mStartToEnd[v[0]] = v[1]
			}
		} else {
			mStartToEnd[v[0]] = v[1]
			starts = append(starts, v[0])
		}
	}

	sort.Ints(starts)

	current := -1
	for i := 0; i < len(starts); i++ {
		if i == 0 {
			current = starts[0]
		}

		if starts[i] > current && starts[i] <= mStartToEnd[current] {
			if mStartToEnd[current] < mStartToEnd[starts[i]] {
				mStartToEnd[current] = mStartToEnd[starts[i]]
			}
		} else if starts[i] > current && starts[i] > mStartToEnd[current] {
			result = append(result, []int{current, mStartToEnd[current]})
			current = starts[i]
		}
	}

	result = append(result, []int{current, mStartToEnd[current]})

	return result
}
