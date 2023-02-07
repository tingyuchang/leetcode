package _0230207

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

// using new idea to do merge, avoid using append to reduce alloc
func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	i, m, n := 0, 0, 0

	for m < len(a) && n < len(b) {
		if a[m] > b[n] {
			c[i] = b[n]
			n++
		} else {
			c[i] = a[m]
			m++
		}
		i++
	}
	for m < len(a) {
		c[i] = a[m]
		i++
		m++
	}

	for n < len(b) {
		c[i] = b[n]
		i++
		n++
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

func QuickSort(nums []int, p, r int) []int {
	if p < r {
		q := partition(nums, p, r)
		QuickSort(nums, p, q-1)
		QuickSort(nums, q+1, r)
	}
	return nums
}

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] > x {
			continue
		} else {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

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

func removeDuplicatesFromSortedArray(head *ListNode) *ListNode {
	var node *ListNode
	node = head

	for node != nil {
		if node.Next != nil && node.Val == node.Next.Val {
			if node.Next == nil {
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
