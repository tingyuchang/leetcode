package Math

import (
	"math"
)

func maximum69Number(num int) int {
	n := num
	count := 0
	changing := -1
	for n > 0 {
		if n%10 == 6 {
			changing = count
		}
		n = n / 10
		count++
	}

	if changing == -1 {
		return num
	}

	return num + int(math.Pow(float64(10), float64(changing)))*3
}
