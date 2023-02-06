package _0230206

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Show() string {
	var dummyNode *ListNode
	var s string
	dummyNode.Next = n
	for dummyNode.Next != nil {
		s = s + strconv.Itoa(dummyNode.Next.Val)
		dummyNode = dummyNode.Next
	}

	s = s + strconv.Itoa(dummyNode.Next.Val)

	return s
}

func MergeSort(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	listA := nums[:n/2]
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
	buildHeap(nums)
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
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

func buildHeap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		maxHeap(nums, i)
	}
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums)

	for r >= l {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}

	}

	return l
}

func RemoveDuplicatesFromSortedArray(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}
	node := head
	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			if node.Next.Next == nil {
				node.Next = nil
				return head
			} else {
				node.Next = node.Next.Next
				continue
			}
		}

		node = node.Next
	}

	return head
}
