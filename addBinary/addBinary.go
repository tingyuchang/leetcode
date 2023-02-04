package addBinary

import "strconv"

func AddBinary(a string, b string) string {
	var c string
	carry := 0
	n := 0

	for n < len(a) || n < len(b) {
		var sum int

		i := len(a) - n - 1
		j := len(b) - n - 1

		sum = int(a[i]) + int(a[j]) + carry

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
