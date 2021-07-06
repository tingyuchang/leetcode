package divide

import (
	"math"
)

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8
and truncate(-2.7335) = -2.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer
range: [−2^31, 2^31 − 1]. For this problem, assume that your function returns 2^31 − 1 when the division result overflows.
 */

func Divide(dividend int, divisor int) int {
	sum := 0
	result := 0

	isNegative := false

	newDivisor := divisor
	divisorCount := 1

	if dividend < 0  {
		dividend = 0 - dividend
		isNegative  = !isNegative
	}

	if divisor < 0  {
		divisor = 0 - divisor
		newDivisor = divisor
		isNegative  = !isNegative
	}

	for sum < dividend {
		if dividend-sum < newDivisor {
			if newDivisor == divisor {
				break
			}
			newDivisor -= divisor
			divisorCount--
			continue
		}
		result += divisorCount
		sum+=newDivisor
		if dividend - sum > sum  {
			newDivisor = sum
			divisorCount = result
		}
	}

	if isNegative {
		result = 0 - result
	}

	if float64(result) > math.Pow(2, 31)-1 {
		return int(math.Pow(2, 31))-1
	}

	if float64(result) < -math.Pow(2, 31) {
		return -int(math.Pow(2, 31))
	}

	return result
}
