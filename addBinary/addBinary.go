package addBinary

import (
	"strconv"
)

func AddBinary(a string, b string) string {
	var c string
	carry := 0
	n := 0

	for n < len(a) || n < len(b) {
		var sum int
		var x, y int

		if len(a)-n-1 >= 0 {
			x, _ = strconv.Atoi(string(a[len(a)-n-1]))
		}
		if len(b)-n-1 >= 0 {
			y, _ = strconv.Atoi(string(b[len(b)-n-1]))
		}

		sum = x + y + carry
		switch sum {
		case 0:
			carry = 0
		case 1:
			carry = 0
		default:
			carry = 1
			sum -= 2
		}
		c = strconv.Itoa(sum) + c
		n++
	}

	if carry == 1 {
		c = "1" + c
	}

	return c
}
