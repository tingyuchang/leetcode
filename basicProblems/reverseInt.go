package basicProblems

import (
	"fmt"
	"math"
	"strconv"
)

// --- Directions
// Given an integer, return an integer that is the reverse
// ordering of numbers.
// --- Examples
//   reverseInt(15) === 51
//   reverseInt(981) === 189
//   reverseInt(500) === 5
//   reverseInt(-15) === -51
//   reverseInt(-90) === -9

func reverseInt(n int) int {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	var results []int
	for n > 0 {
		results = append(results, n%10)
		n = n/10
	}
	sum := 0
	for i,v := range results {
		sum += v*int(math.Pow10(len(results)-1-i))
	}

	if isNegative {
		sum = -sum
	}
	return sum
}

func reverseInt2(n int) int {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	result,_ := strconv.Atoi(reverseString(fmt.Sprintf("%v",n)))
	if isNegative {
		result = -result
	}
	return result
}