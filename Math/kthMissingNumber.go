package Math

import "sort"

func findKthPositive(arr []int, k int) int {
	//O(lgn)
	i := sort.Search(len(arr), func(i int) bool {
		return arr[i] > i+k
	})
	return i + k
	// O(n)
	//counter := make(map[int]bool)
	//
	//for _, v := range arr {
	//	counter[v] = true
	//}
	//temp := k
	//for i := 1; i <= len(arr)+k; i++ {
	//	if !counter[i] {
	//		temp--
	//	}
	//	if temp == 0 {
	//
	//		return i
	//	}
	//}
	//
	//return len(arr) + k + 1

}
