package __Daily_Prac

import (
	"container/list"
	"leetcode/LinkedList"
	"leetcode/Tree"
	"sort"
)

// Google's questions
// Find the Kth largest element in a number stream
// https://leetcode.com/problems/kth-largest-element-in-an-array/
/*
Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
You must solve it in O(n) time complexity.
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
*/

type heap struct {
	data []int
}

func maxHeapTopDown(nums []int, n int) {
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
		maxHeapTopDown(nums, largest)
	}
}

func maxHeapBottomUp(nums []int, n int) {
	current := n
	parent := (current - 1) >> 1

	for current != 0 && nums[parent] < nums[current] {
		nums[parent], nums[current] = nums[current], nums[parent]
		current = parent
		parent = (current - 1) >> 1
	}
}

func GFindKthLargest(nums []int, k int) int {
	/*
		approach1: sorting array -> T: n^2 S: n
		approach2: heap -> T: nlogn S:1
		approach3: quick select -> T: n  S: 1
	*/
	/*  heap
	for i := len(nums) - 1; i >= len(nums)/2; i-- {
		maxHeapTopDown(nums, i)
	}

	for k != 1 {
		last := len(nums) - 1
		nums[last], nums[0] = nums[0], nums[last]
		nums = nums[:last]
		maxHeapTopDown(nums, 0)
		k -= 1
	}
	return nums[0]
	*/
	// randomised select
	return GFindKthLargestRandomizedSelection(nums, 0, len(nums)-1, len(nums)-k+1)
}

func GFindKthLargestRandomizedSelection(nums []int, p, r, target int) int {
	if p == r {
		return nums[p]
	}

	q := GFindKthLargestPartition(nums, p, r)

	k := q - p + 1

	if k == target {
		return nums[q]
	}

	if k > target {
		return GFindKthLargestRandomizedSelection(nums, p, q-1, target)
	} else {
		return GFindKthLargestRandomizedSelection(nums, q+1, r, target-k)
	}
}

func GFindKthLargestPartition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1

	for j := p; j < r; j++ {
		if nums[j] < x {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	i += 1
	nums[i], nums[r] = nums[r], nums[i]

	return i
}

/*
658. Find K Closest Elements
https://leetcode.com/problems/find-k-closest-elements/
Given a sorted integer array arr, two integers k and x,
return the k closest integers to x in the array. The result should also be sorted in ascending order.
An integer a is closer to x than an integer b if:
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b

Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]
Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
*/
func GFindClosestElements(arr []int, k int, x int) []int {
	return []int{}
}

/*
237. Delete Node in a Linked List
https://leetcode.com/problems/delete-node-in-a-linked-list/
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
*/
func GDeleteNode(node *LinkedList.ListNode) {

}

/*
138. Copy List with Random Pointer
https://leetcode.com/problems/copy-list-with-random-pointer/
*/

func GCopyRandomList(head *LinkedList.Node) *LinkedList.Node {
	return nil
}

/*
226. Invert Binary Tree

https://leetcode.com/problems/invert-binary-tree/

*/

func GInvertTree(root *Tree.TreeNode) *Tree.TreeNode {

	return root

}

/*
100. Same Tree
https://leetcode.com/problems/same-tree/
*/
func GIsSameTree(p *Tree.TreeNode, q *Tree.TreeNode) bool {
	return true
}

/*
139. Word Break
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func GWordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordMaps := make(map[string]struct{})

	for _, v := range wordDict {
		wordMaps[v] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordMaps[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

/*
647. Palindromic Substrings
https://leetcode.com/problems/palindromic-substrings/
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/
func GCountSubstrings(s string) int {
	return 0
}

/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/
*/

func GMaxSubArray(nums []int) int {
	current := 0
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := range nums {
		v := nums[i]
		current = max(v, current+v)
		ans = max(ans, current)
	}

	return ans

}

/*
65. Valid Number
https://leetcode.com/problems/valid-number/

*/

func GIsNumber(s string) bool {
	return true
}

/*
22. Generate Parentheses
https://leetcode.com/problems/generate-parentheses/
*/
func GGenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	left, right := n-1, n
	current := []byte{'('}

	res := []string{}

	dfsGenerateParenthesis(left, right, &res, current)

	return res
}

func dfsGenerateParenthesis(left, right int, res *[]string, current []byte) {

	if left == 0 && right == 0 {
		temp := current[:len(current)]
		*res = append(*res, string(temp))
	}

	if left > 0 {
		dfsGenerateParenthesis(left-1, right, res, append(current, '('))
	}

	if right > left {
		dfsGenerateParenthesis(left, right-1, res, append(current, ')'))
	}
}

/*
146. LRU Cache
https://leetcode.com/problems/lru-cache/
*/

type Key int

type entry struct {
	key Key
	val int
}

type LRUCache struct {
	cap      int
	data     map[Key]*list.Element
	recently *list.List
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, data: make(map[Key]*list.Element), recently: list.New()}
}

// Get
// return value if key exist otherwise -1
// update key to the end of recently

func (this *LRUCache) Get(key int) int {
	if this.data == nil {
		return -1
	}

	if ele, ok := this.data[Key(key)]; ok {
		this.recently.MoveToFront(ele)
		return ele.Value.(*entry).val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if this.data == nil {
		this.data = make(map[Key]*list.Element)
		this.recently = list.New()
	}

	// update existing value
	if ele, ok := this.data[Key(key)]; ok {
		this.recently.MoveToFront(ele)
		ele.Value.(*entry).val = value
		return
	}

	// check cap
	ele := this.recently.PushFront(&entry{key: Key(key), val: value})
	this.data[Key(key)] = ele
	if this.recently.Len() > this.cap {
		this.RemoveOldest()
	}

}

func (this *LRUCache) RemoveOldest() {
	ele := this.recently.Back()
	if ele != nil {
		this.removeElement(ele)
	}
}

func (this *LRUCache) removeElement(e *list.Element) {
	this.recently.Remove(e)

	deleteingEntry := e.Value.(*entry)
	delete(this.data, deleteingEntry.key)
}

/*
34. Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/
func GSearchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	start, end := -1, -1

	for r >= l {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			start, end = mid, mid
			for {
				if start != 0 && nums[start-1] == target {
					start--
				}

				if end < len(nums)-1 && nums[end+1] == target {
					end++
				}
				if (start == 0 || nums[start-1] != target) && (end == len(nums)-1 || nums[end+1] != target) {
					break
				}
			}

			break
		}
	}

	return []int{start, end}
}

/*
56. Merge Intervals
https://leetcode.com/problems/merge-intervals/
*/
func GMergeIntervals(intervals [][]int) [][]int {
	ans := make([][]int, 0)

	intervalMaps := make(map[int]int)
	startTimes := make([]int, 0)
	for _, interval := range intervals {
		end, ok := intervalMaps[interval[0]]

		if ok {
			if end < interval[1] {
				intervalMaps[interval[0]] = interval[1]
			}
		} else {
			intervalMaps[interval[0]] = interval[1]
			startTimes = append(startTimes, interval[0])
		}
	}

	sort.Ints(startTimes)

	current := startTimes[0]

	for i := 1; i < len(startTimes); i++ {
		start := startTimes[i]
		end := intervalMaps[start]
		currentEnd := intervalMaps[current]
		if start <= currentEnd && end > currentEnd {
			intervalMaps[current] = end
		} else {
			if start > currentEnd {
				ans = append(ans, []int{current, currentEnd})
				current = start
			}
		}
	}

	ans = append(ans, []int{current, intervalMaps[current]})
	return ans
}

/*
112. Path Sum

https://leetcode.com/problems/path-sum/
*/

func GHasPathSum(root *Tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return GHasPathSum(root.Left, targetSum-root.Val) || GHasPathSum(root.Right, targetSum-root.Val)
}

/*
268. Missing Number
https://leetcode.com/problems/missing-number/
*/
func GMissingNumber(nums []int) int {
	sum := len(nums)
	if (1+len(nums))%2 != 0 {
		sum = sum / 2 * (1 + len(nums))
	} else {
		sum *= (1 + len(nums)) / 2
	}
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}

	return sum
}

/*
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/
*/
func GReverseList(head *LinkedList.ListNode) *LinkedList.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node, next := head, head.Next

	for next != nil {
		temp := next.Next
		next.Next = node
		node = next
		next = temp
	}

	head.Next = nil

	return node
}

/*
Determine if the sum of three integers is equal to the given value

Intersection point of two linked lists

Move zeros to the left

Add two integers

Merge two sorted linked lists

Convert binary tree to doubly linked list

level order traversal of binary tree

Reverse words in a sentence

Find maximum single sell profit

Calculate the power of a number

Find all possible subsets

Clone a directed graph

Serialize / deserialize binary tree

Search rotated array

Set columns and rows as zeros

Connect all siblings

Find all sum combinations

Clone a directed graph

Closest meeting point

Search for the given key in a 2d matrix
*/
