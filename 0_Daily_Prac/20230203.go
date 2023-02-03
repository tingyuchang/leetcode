package Daily_Prac

import "fmt"

func MergeSort(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	listA := nums[0 : n/2]
	listB := nums[n/2:]

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

func HeapSort(nums []int) []int {
	n := len(nums)
	nums = buildHeap(nums)
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		nums = nums[:i]
		maxHeap(nums, 0)
	}
	return nums[:n]
}

func maxHeap(nums []int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var largest int
	if l < len(nums) && nums[l] > nums[n] {
		largest = l
	} else {
		largest = n
	}

	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != n {
		nums[n], nums[largest] = nums[largest], nums[n]
		maxHeap(nums, largest)
	}
}

func buildHeap(nums []int) []int {
	n := len(nums)

	for i := n / 2; i >= 0; i-- {
		maxHeap(nums, i)
	}
	fmt.Println(nums)
	return nums
}
