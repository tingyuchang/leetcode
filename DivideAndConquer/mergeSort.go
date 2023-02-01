package DivideAndConquer

func MergeSort(unsortedList []int) []int {
	n := len(unsortedList)

	if n == 1 {
		return unsortedList
	}

	listA := unsortedList[0 : n/2]
	listB := unsortedList[n/2:]

	listA = MergeSort(listA)
	listB = MergeSort(listB)

	return merge(listA, listB)
}

func merge(a, b []int) []int {
	var c []int
	m, n := 0, 0

	for m < len(a) && n < len(b) {
		if a[m] > b[n] {
			c = append(c, b[n])
			n++
		} else {
			c = append(c, a[m])
			m++
		}
	}

	if m < len(a) {
		c = append(c, a[m:]...)
	}

	if n < len(b) {
		c = append(c, b[n:]...)
	}
	return c
}
