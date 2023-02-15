package Math

func AddToArrayForm(num []int, k int) []int {
	// ks is reverse
	var ks []int
	for k != 0 {
		s := k % 10
		ks = append(ks, s)
		k = k / 10
	}
	result := make([]int, max(len(num), len(ks))+1)
	var carry bool

	m, n, l := len(num)-1, 0, len(result)-1

	for m >= 0 && n < len(ks) {
		sum := num[m] + ks[n]

		if carry {
			sum++
			carry = false
		}

		if sum >= 10 {
			sum = sum % 10
			carry = true
		}

		result[l] = sum

		l--
		m--
		n++
	}

	for l >= 0 {

		if l == 0 {
			if carry {
				result[0] = 1
				carry = false
			}
			break
		}
		var sum int
		if m >= 0 {
			sum = num[m]
			m--
		} else if n < len(ks) {
			sum = ks[n]
			n++
		}
		if carry {
			sum++
			carry = false
		}
		if sum >= 10 {
			sum = sum % 10
			carry = true
		}
		result[l] = sum

		l--
	}

	if result[0] == 0 {
		return result[1:]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
