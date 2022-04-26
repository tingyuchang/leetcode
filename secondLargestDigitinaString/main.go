package main

/*
Given an alphanumeric string s, return the second largest numerical digit that appears in s,
or -1 if it does not exist.
An alphanumeric string is a string consisting of lowercase English letters and digits.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(secondHighest("dfa12321afd"))
	fmt.Println(secondHighest("abc1111"))
	fmt.Println(secondHighest("ck077"))
}

func secondHighest(s string) int {
	var max uint8 = 47
	var second uint8 = 47
	for i := 0; i < len(s); i++ {
		// skip same numbers
		if s[i] == max || s[i] == second {
			continue
		}
		// 48 - 57 = 0 - 9
		if s[i] >= 48 && s[i] <= 57 {
			if s[i] > max {
				second = max
				max = s[i]
			} else if s[i] > second {
				second = s[i]
			}
		}
	}
	if second == 47 || max == second {
		return -1
	}
	result, _ := strconv.Atoi(string(second))
	return result
}
