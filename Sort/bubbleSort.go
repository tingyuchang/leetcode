package Sort

func BubbleSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := len(n) - 1; j > i; j-- {
			if n[j] < n[j-1] {
				n[j], n[j-1] = n[j-1], n[j]
			}
		}
	}
	return n
}
