package basicProblems

import "math"

func isPrime(n int) (isPrime bool) {
	isPrime = true
	squireRoot := int(math.Pow(float64(n), -1))
	for i := 2; i < squireRoot; i ++ {
		if n % 2== 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}