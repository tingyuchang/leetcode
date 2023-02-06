package Sort

// priority queue 其實是一個 abstract class -- 它不像二元樹或 doubly linked list 等等那麼清楚地有明確的結構, 它只是一個抽象的介面 interface

/*
- insert(p): Inserts a new element with priority p.
- extractMax(): Extracts an element with maximum priority.
- remove(i): Removes an element pointed by an iterator i.
- getMax(): Returns an element with maximum priority.
- changePriority(i, p): Changes the priority of an element pointed by i to p.
*/

func GetMax(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	return nums[0]
}

func ExtractMax(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	max := nums[0]
	nums[0] = nums[n-1]
	nums = nums[:n-1]
	maxHeap(nums, 0)
	return max
}

func parent(n int) int {
	return (n - 1) / 2
}

//Insert a new element x with priority k.
// this is a problem between x & k,
func InsertKey(nums []int, x, k int) {
	i := k
	for i > 1 && nums[parent(i)] < nums[i] {
		nums[parent(i)], nums[i] = nums[i], nums[parent(i)]
		i = parent(i)
	}
}
