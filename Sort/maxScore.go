package Sort

import (
	"sort"
)

func MaxScore(nums1 []int, nums2 []int, k int) int64 {
	return maxScore(nums1, nums2, k)
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	ans := 0

	type node struct {
		a int
		b int
	}

	nodes := make([]node, len(nums1))

	for i := 0; i < len(nodes); i++ {
		nodes[i] = node{
			a: nums1[i],
			b: nums2[i],
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].b > nodes[j].b
	})

	minHeap := MinHeap{}
	tempSum := 0
	for i := 0; i < k; i++ {
		tempSum += nodes[i].a
		minHeap.Add(nodes[i].a)
	}

	ans = tempSum * nodes[k-1].b

	for i := k; i < len(nodes); i++ {

		popVal := minHeap.replace(nodes[i].a)

		tempSum = tempSum - popVal + nodes[i].a

		if tempSum*nodes[i].b > ans {
			ans = tempSum * nodes[i].b
		}

	}

	return int64(ans)
}

type MinHeap struct {
	data []int
}

func (mh *MinHeap) Add(n int) {
	if mh.data == nil {
		mh.data = make([]int, 0)
	}
	mh.data = append(mh.data, n)
	minHeapBottomUp(mh.data, len(mh.data)-1)
}

func (mh *MinHeap) replace(n int) int {
	if mh.data == nil {
		mh.data = make([]int, 0)
	}
	if len(mh.data) == 0 {

		mh.data = append(mh.data, n)
		return 0
	}
	oldVal := mh.data[0]
	mh.data[0] = n
	minHeapTopDown(mh.data, 0)
	return oldVal
}

func minHeapBottomUp(nums []int, n int) {
	current := n
	parentNode := (n - 1) >> 1
	for parentNode >= 0 && nums[parentNode] > nums[current] {
		nums[current], nums[parentNode] = nums[parentNode], nums[current]
		current = parentNode
		parentNode = (current - 1) >> 1
	}
}

func minHeapTopDown(nums []int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1

	var smallest int

	if l < len(nums) && nums[l] < nums[n] {
		smallest = l
	} else {
		smallest = n
	}

	if r < len(nums) && nums[r] < nums[smallest] {
		smallest = r
	}

	if smallest != n {
		nums[n], nums[smallest] = nums[smallest], nums[n]
		minHeapTopDown(nums, smallest)
	}

}
