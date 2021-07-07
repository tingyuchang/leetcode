package divide

import (
	"fmt"
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
	fmt.Println("Start:", dividend, divisor)

	if dividend == 0 {
		return 0
	}

	quotient := 0

	isNegative := false

	if dividend < 0  {
		isNegative  = !isNegative
	}

	if divisor < 0  {
		isNegative  = !isNegative
	}

	absDividend, absDivisor := abs(dividend), abs(divisor)

	for absDividend >= absDivisor {
		tmp, m := absDivisor, 1
		for absDividend >= tmp <<1 {
			tmp, m = tmp << 1, m <<1
			fmt.Printf("tmp:\t %d\n", tmp)
			fmt.Printf("m:\t %d\n", m)
		}
		absDividend -= tmp
		quotient += m
	}

	if isNegative {
		quotient = - quotient
	}

	if quotient < -2147483648 {
		return -2147483648
	} else if quotient > 2147483647 {
		return 2147483647
	}

	return quotient
}

func abs(n int) int64 {
	absn := int64(n)
	if absn < 0 {
		return -absn
	}
	return absn
}
