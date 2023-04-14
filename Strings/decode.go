package Strings

import (
	"strconv"
)

func DecodeString(s string) string {
	return decodeString(s)
}

func decodeString(s string) string {
	stack := make([]byte, 0)

	// store indexes of '['
	cache := make([]int, 0)
	for i := 0; i < len(s); i++ {

		if s[i] == ']' {
			// find '[' in j & digit in j-1
			leftBracket := cache[len(cache)-1]
			cache = cache[:len(cache)-1]

			// find digit ex: 3, 10, or 100
			digit := string(stack[leftBracket-1])
			k := leftBracket - 2
			for k >= 0 {
				if stack[k] < 48 || stack[k] > 57 {
					break
				}
				digit = string(stack[k]) + digit
				k--
			}

			times, _ := strconv.Atoi(digit)

			tmp := string(stack[leftBracket+1:])
			stack = stack[:k+1]
			//fmt.Println("left", leftBracket, "times", times, "stack", string(stack), "tmp ", tmp)
			for j := 0; j < times; j++ {
				stack = append(stack, tmp...)
			}
			//fmt.Println("After insert: ", string(stack))

			continue
		}

		stack = append(stack, s[i])
		if s[i] == '[' {
			cache = append(cache, len(stack)-1)
		}

	}

	return string(stack)
}
