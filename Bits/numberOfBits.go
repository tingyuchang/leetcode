package Bits

func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		if num%2 == 1 {
			result++
		}
		num = num >> 1
	}

	return result
}
