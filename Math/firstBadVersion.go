package Math

func firstBadVersion(n int) int {
	l := 1
	r := n

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if isBadVersion(mid) == false && isBadVersion(mid) == true {
			return mid + 1
		}

		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return l
}

func isBadVersion(n int) bool {
	return true
}
