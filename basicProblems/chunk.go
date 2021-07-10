package basicProblems

// Given an array and chunk size, divide the array into many subarrays
// where each subarray is of length size
// --- Examples
// chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
// chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
// chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) -->
// [[ 1, 2, 3], [4, 5, 6], [7, 8]]
// chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
// chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]

func chunk(n []int, size int) [][]int {
	var chunked [][]int
	index := 0

	for index < len(n) {
		if index+size > len(n) {
			chunked = append(chunked, n[index:])
		} else {
			chunked = append(chunked, n[index:index+size])
		}
		index+=size
	}
	return chunked
}
func chunk2(n []int, size int) [][]int {
	capacity := len(n)/size
	isRemainder := false
	if len(n)%size > 0 {
		isRemainder = true
	}

	var result [][]int
	quotient :=0
	for quotient < capacity {
		result = append(result, n[quotient*size:quotient*size+size])
		quotient++
	}
	if isRemainder {
		result = append(result, n[quotient*size:])
	}

	return result
}
