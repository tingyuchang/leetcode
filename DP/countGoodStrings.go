package DP

func CountGoodStrings(low int, high int, zero int, one int) int {
	return countGoodStrings(low, high, zero, one)
}

func countGoodStrings(low int, high int, zero int, one int) int {

	var minVal, maxVal int

	if one > zero {
		maxVal = one
		minVal = zero
	} else {
		maxVal = zero
		minVal = one
	}

	sum := 0
	if zero == one {
		low = low / zero
		high = high / zero
		power := make([]int, high+1)
		power[0] = 1
		for i := 1; i < len(power); i++ {
			power[i] = (power[i-1] * 2) % MOD
		}

		for i := low; i <= high; i++ {
			// needs to check validation here
			sum = (sum + power[i]) % MOD
		}
		return sum

	} else {
		dp := make([]int, high+1)
		// high always greater than max(one, zero)?
		for i := 0; i < maxVal; i++ {
			if i%minVal == 0 {
				dp[i] = 1
			}
		}
		if maxVal%minVal == 0 {
			dp[maxVal] = 2
		} else {
			dp[maxVal] = 1
		}

		for i := maxVal + 1; i <= high; i++ {
			dp[i] = (dp[i-one] + dp[i-zero]) % MOD
		}

		for i := low; i <= high; i++ {
			sum = (sum + dp[i]) % MOD
		}
	}
	return sum
}

/*

if one != zero

dp[i] = dp[i-zero] + dp[i-one]

dp[0] = 1

dp[0 ~ max(zero, one)] = 1

dp[min(zero, one)] = 1

dp[max(zero, one)] = 2


if one == zero

dp[i] = 2^(i/one)
*/
