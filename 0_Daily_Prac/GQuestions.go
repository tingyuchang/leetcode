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
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	/*

		approach1: binary search + expand T: logn+k, s: k

		approach2: 2 point narrow down T: k s: 1

		lower <-----> upper T: n
		if abs(arr[lower] - x ) >  abs(arr[upper] - x)
		lower ++
		else
		upper--
	*/

	/*
		if x <= arr[0] {
			return arr[:k]
		}

		if x >= arr[len(arr)-1] {
			return arr[len(arr)-k:]
		}

		// binary search

		l := 0
		r := len(arr) - 1

		for r >= l {
			mid := int(uint(l+r) >> 1)

			if arr[mid] <= x {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		l, r = r, l

		ans := make([]int, 0)
		for k > 0 {
			if l < 0 {
				ans = append(ans, arr[r:r+k]...)
				break
			}

			if r == len(arr) {
				ans = append(arr[l-k+1:l+1], ans...)
				break
			}

			if abs(arr[l]-x) < abs(arr[r]-x) || abs(arr[l]-x) == abs(arr[r]-x) {
				ans = append([]int{arr[l]}, ans...)
				l -= 1
			} else {
				ans = append(ans, arr[r])
				r += 1
			}
			k--
		}

		return ans
	*/
	l, r := 0, len(arr)-1

	for r-l >= k {
		if abs(arr[l]-x) > abs(arr[r]-x) {
			l += 1
		} else {
			r -= 1
		}
	}

	return arr[l : r+1]

}

/*
237. Delete Node in a Linked List
https://leetcode.com/problems/delete-node-in-a-linked-list/
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
*/
func GDeleteNode(node *LinkedList.ListNode) {
	/*
		appraoch:  copy next node's val to current and remove last node from linked list
	*/

	for node != nil {
		node.Val = node.Next.Val

		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next

	}
}

/*
138. Copy List with Random Pointer
https://leetcode.com/problems/copy-list-with-random-pointer/
*/

func GCopyRandomList(head *LinkedList.Node) *LinkedList.Node {
	/*
		approach: iterating linked list to creating new list
		and using hashMap to mapping olf nodes and new nodes
		and leaving random as empty
		the 2nd run iterating and using hashMap to assign random node in new list
	*/

	prehead := &LinkedList.Node{}
	node := &LinkedList.Node{Val: head.Val}
	prehead.Next = node

	nodesMap := make(map[*LinkedList.Node]*LinkedList.Node)
	nodesMap[head] = node
	current := head.Next

	for current != nil {
		temp := &LinkedList.Node{Val: current.Val}
		node.Next = temp
		nodesMap[current] = temp
		node = node.Next
		current = current.Next
	}

	current = head
	node = prehead.Next

	for current != nil {
		node.Random = nodesMap[current.Random]
		current = current.Next
		node = node.Next
	}

	return prehead.Next
}

/*
226. Invert Binary Tree

https://leetcode.com/problems/invert-binary-tree/

*/

func GInvertTree(root *Tree.TreeNode) *Tree.TreeNode {
	if root == nil {
		return nil
	}

	temp := GInvertTree(root.Left)
	root.Left = GInvertTree(root.Right)
	root.Right = temp

	return root

}

/*
100. Same Tree
https://leetcode.com/problems/same-tree/
*/
func GIsSameTree(p *Tree.TreeNode, q *Tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return GIsSameTree(p.Left, q.Left) && GIsSameTree(p.Right, q.Right)
}

/*
139. Word Break
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func GWordBreak(s string, wordDict []string) bool {
	/*
		approach: DP

		using an array to store each letters for position j to i is existing in wordDict

		dp[i] = true when s[j:i] in wordDict && dp[j] == true
		j from 0 to i

		TC O(n^2)
		SC O(n)
	*/

	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordMap := make(map[string]struct{})

	for _, v := range wordDict {
		wordMap[v] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordMap[s[j:i]]; ok && dp[j] {
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
	ans := 0

	for i := 0; i < len(s); i++ {
		ans += GCountSubstringsCheck(s, i, i)
		ans += GCountSubstringsCheck(s, i, i+1)
	}

	return ans
}

func GCountSubstringsCheck(s string, start, end int) int {
	count := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start -= 1
		end += 1
		count += 1
	}

	return count
}

/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/
*/

func GMaxSubArray(nums []int) int {
	ans := 0
	current := 0

	for i := 0; i < len(nums); i++ {
		current = Max(nums[i], current+nums[i])
		ans = Max(ans, current)
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
	/*
		approach1: brute force
		generate all possible strings and check it is valid or not

		using queue to do the bfs
		1 ['(', ')']
		2 ['()'. '((', ')(', '))']
		3 ....
		when string length == 2 * n
		check validation

		TC O(2^n*n)
		SC O(2^n)

		queue := []string{"(", ")"}
		ans := make([]string, 0)

		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			if len(current) == 2*n {
				if isValidateParenthesis(current) {
					ans = append(ans, current)
				}
				continue
			}
			next1 := current + "("
			next2 := current + ")"

			queue = append(queue, next1, next2)

		}

		return ans
	*/
	/*
			approach 2: backtrack, only build valid string

			like ")" should put it into the queue, we can reduce half size
			we need 2 vars to help us
			leftCount: if leftCount < n, newStr = current + '(' then leftCount += 1
			rightCount: only add ")" when leftCount > rightCount,  newStr = current + ')' then rightCount += 1

		TC O(4^n/(n^(3/2)))

	*/
	ans := make([]string, 0)
	backtrackingParenthesis("", 0, 0, n, &ans)
	return ans
	/*
		approach3 divide & conquer

		F(n) = (F(0)F(n-1) + F(1)F(n-2) + .... + F(n-1)F(0))

		F(0)F(n-1) =
		F(1)F(n-2)
		F(n-1)F(0)
	*/

}

func isValidateParenthesis(s string) bool {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count += 1
		} else {
			count -= 1
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

func backtrackingParenthesis(s string, leftCount, rightCount, n int, ans *[]string) {
	if len(s) == 2*n {
		*ans = append(*ans, s)
		return
	}
	if leftCount < n {
		backtrackingParenthesis(s+"(", leftCount+1, rightCount, n, ans)
	}

	if rightCount < leftCount {
		backtrackingParenthesis(s+")", leftCount, rightCount+1, n, ans)
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
	/*
		approach binary search find target and expand from target to get start and end position
		of the same elements
	*/

	l := 0
	r := len(nums) - 1
	start, end := -1, -1

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			// find target
			start, end = mid, mid
			for {
				if start != 0 && nums[start-1] == target {
					start -= 1
				}
				if end < len(nums)-1 && nums[end+1] == target {
					end += 1
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
	/*
	   	approach:

	   	sort intervals by start time
	   	1  2 3  6  8 10 15 18
	   	[--------]
	   	          [----]
	    					[---]
	*/

	ans := make([][]int, 0)
	// start: end
	intervalMap := make(map[int]int)

	for _, interval := range intervals {
		if end, ok := intervalMap[interval[0]]; ok {
			if end < interval[1] {
				intervalMap[interval[0]] = interval[1]
			}
		} else {
			intervalMap[interval[0]] = interval[1]
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	currentStart := intervals[0][0]

	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		currentEnd := intervalMap[currentStart]
		if start <= currentEnd && end > currentEnd {
			intervalMap[currentStart] = end
		} else if start > currentEnd {
			ans = append(ans, []int{currentStart, currentEnd})
			currentStart = start
		}
	}
	ans = append(ans, []int{currentStart, intervalMap[currentStart]})

	return ans

	/*
		Interval search trees
		1. build BST and using left endpoint as BST key
		2. store max end point in  subtree rooted at node
		insert
		1. if intersects(root.interval, target.interval) update root.interval, min, max
		2. if insert.min < root.min, root = root.left
		3  else  root = root.right
		4. if root == nil, root.left or right = interval

		search
		1. if intersects(root.interval, target.interval) return root.interval
		2. else if root.left == nil, root = root.right
		3. else if root.left.max < target.min, root = root.right
		4. root = root.left
	*/
}

/*
112. Path Sum

https://leetcode.com/problems/path-sum/
*/

func GHasPathSum(root *Tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := []*Tree.TreeNode{root}

	for len(queue) > 0 {
		currentSize := len(queue)

		for i := 0; i < currentSize; i++ {
			currentNode := queue[i]

			if currentNode.Val == targetSum && currentNode.Left == nil && currentNode.Right == nil {
				return true
			}

			// move on

			if currentNode.Left != nil && targetSum-currentNode.Val > 0 {
				currentNode.Left.Val += currentNode.Val
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil && targetSum-currentNode.Val > 0 {
				currentNode.Right.Val += currentNode.Val
				queue = append(queue, currentNode.Right)
			}
		}

		queue = queue[currentSize:]
	}

	return false
}

/*
268. Missing Number
https://leetcode.com/problems/missing-number/
*/
func GMissingNumber(nums []int) int {
	/*
				approach1: hasMap to store current  value and using for-loop for 0 to len(nums)
				TC O(n)
				SC O(n)

		valMap := make(map[int]struct{})
			for _, v := range nums {
				valMap[v] = struct{}{}
			}

			for i := 0; i <= len(nums); i++ {
				if _, ok := valMap[i]; !ok {
					return i
				}
			}

			return len(nums)

			approach2: sum from 1 to len(nums)  only work on  0 <= nums[i] <= n
			and decrease by each val in nums
			TC O(n)
			SC O(1)
	*/

	sum := 0

	for i := 1; i <= len(nums); i++ {
		sum += i
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
	current := head
	next := head.Next

	for next != nil {
		temp := next.Next
		next.Next = current
		current = next
		next = temp
	}
	head.Next = nil

	return current
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
