package __Daily_Prac

import (
	"container/list"
	"fmt"
	"leetcode/Tree"
	"sort"
	"strconv"
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
func GDeleteNode(node *ListNode) {
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

func GCopyRandomList(head *NodeR) *NodeR {
	/*
		approach: iterating linked list to creating new list
		and using hashMap to mapping olf nodes and new nodes
		and leaving random as empty
		the 2nd run iterating and using hashMap to assign random node in new list
	*/

	prehead := &NodeR{}
	node := &NodeR{Val: head.Val}
	prehead.Next = node

	nodesMap := make(map[*NodeR]*NodeR)
	nodesMap[head] = node
	current := head.Next

	for current != nil {
		temp := &NodeR{Val: current.Val}
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
func GReverseList(head *ListNode) *ListNode {
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
15. 3Sum
https://leetcode.com/problems/3sum/
*/

func GThreeSum(nums []int) [][]int {
	/*
		approach:
		1. soring array
		2. iterating array and using  lower & upper to find arr[i] + arr[lower] + arr[upper] = 0
			- if  arr[i] + arr[lower] + arr[upper] > 0 lower += 1
			- if arr[i] + arr[lower] + arr[upper] < 0 upper -= 1

		TC O(n^3)
	*/

	ans := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for r > l {
			sum := nums[i] + nums[l] + nums[r]

			if sum > 0 {
				r -= 1
			} else if sum < 0 {
				l += 1
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return ans
}

/*
160. Intersection of Two Linked Lists
https://leetcode.com/problems/intersection-of-two-linked-lists/
*/
func GGetIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	for nodeB != nodeA {
		if nodeA == nodeB {
			return nodeA
		}

		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}

		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nil
}

/*
283. Move Zeroes
https://leetcode.com/problems/move-zeroes/
*/
func GMoveZeroes(nums []int) {
	/*
			approach: using nonzero index to address most left zero index
			if nums[i] != 0, swap nums[nonzero] & nums[i], nonzero += 1
			e.g. V
				 0,1,0,3,12
					V
				1, 0 0 3, 12
				 	 V
				1 3  0  0 12
			TC O(n)
			SC O(1)

		 approach2 create new array and fill in 0 in default

		TC O(n)
		SC O(n)
	*/

	nonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZero], nums[i] = nums[i], nums[nonZero]
			nonZero += 1
		}
	}
}

/*
21. Merge Two Sorted Lists

https://leetcode.com/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	var node = &ListNode{}
	preHead.Next = node

	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	if list1 == nil {
		node.Next = list2
	} else {
		node.Next = list1
	}
	return preHead.Next.Next
}

/*
114. Flatten Binary Tree to Linked List

https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
*/
func flatten(root *Tree.TreeNode) {

	/*

		approach1: recursive
		1. flatten(root.Right)
		2. flatten(root.Left)
		3. root.Right = root.Left
		4. root.Left = nil
		5. make node = root.Right and  node = node.Right until  root.Right = nil
		6. node.Right = temp

		TC O(n)

		approach 2: using an array to store node in preorder sequence

		then iterating array
		1. preorder[i].Right = preorder[i+1]
		2. preorder[i].Left = nil

		TC O(n)
		SC O(n)
	*/

	if root == nil {
		return
	}
	var temp *Tree.TreeNode

	if root.Right != nil {
		flatten(root.Right)
		temp = root.Right
	}

	if root.Left != nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil

		current := root.Right

		for current != nil {
			if current.Right == nil {
				break
			}
			current = current.Right
		}

		current.Right = temp
	}
}

/*
151. Reverse Words in a String
https://leetcode.com/problems/reverse-words-in-a-string/
*/

func GReverseWords(s string) string {

	/*
		approach 1 If we are allowed to use split function
		1. Split string with delimiter " " to an array
		2. reverse array with " "

		TC O(n)
		SC O(n)

		approach 2 write split function from end of array and export valid string to result
	*/
	ans := ""
	l := len(s) - 1
	r := len(s)

	for l >= 0 {
		if s[l] == ' ' {
			// export
			if len(s[l+1:r]) >= 1 && s[l+1:r] != " " {
				ans += s[l+1:r] + " "
			}
			r = l
		}
		l -= 1
	}
	if r != 0 {
		ans += s[:r]
	}
	if len(ans) > 0 && ans[len(ans)-1] == ' ' {
		ans = ans[:len(ans)-1]
	}
	return ans
}

/*
91. Decode Ways

https://leetcode.com/problems/decode-ways/
*/

func GNumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]

			if i < len(s)-1 && (s[i] == '1' || s[i] == '2' && s[i+1] < '7') {
				dp[i] += dp[i+2]
			}
		}

	}

	return dp[0]
}

/*
394. Decode String
https://leetcode.com/problems/decode-string/

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Input: s = "3[a2[c]]"
Output: "accaccacc"
*/

func GDecodeString(s string) string {
	/*
		approach 1: using stack to store temp value
		1. put element into stack
		2. when we meet close bracket
	*/
	stack := make([]byte, 0)
	indexOfOpenBrackets := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			openBraket := indexOfOpenBrackets[len(indexOfOpenBrackets)-1]
			indexOfOpenBrackets = indexOfOpenBrackets[:len(indexOfOpenBrackets)-1]

			k := openBraket - 2
			for k >= 0 {
				if stack[k] > '9' || stack[k] < '0' {
					break
				}
				k -= 1
			}

			repeatNumber, _ := strconv.Atoi(string(stack[k+1 : openBraket]))
			currentString := string(stack[openBraket+1:])

			fmt.Println(repeatNumber, currentString, k, openBraket, string(stack))
			stack = stack[:k+1]

			for j := 0; j < repeatNumber; j++ {
				stack = append(stack, currentString...)
			}
			fmt.Println(string(stack))
		} else {
			stack = append(stack, s[i])
			if s[i] == '[' {
				indexOfOpenBrackets = append(indexOfOpenBrackets, len(stack)-1)
			}
		}
	}

	return string(stack)
}

/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/
func GMaxProfit(prices []int) int {
	/*
		approach:
		[7,1,5,3,6,4]

		minVal: 7
		ans: 0

		if nums[i] > minVal => max(ans, nums[i] - minVal)
		if nums[i] < minVal => minVal = nums[i]

	*/

	ans := 0
	minVal := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		} else {
			profit := prices[i] - minVal
			if profit > ans {
				ans = profit
			}
		}
	}

	return ans
}

/*
50. Pow(x, n)

https://leetcode.com/problems/powx-n/
*/

func GMyPow(x float64, n int) float64 {
	/*

		if x == 1  {

		}

		 ans := 1
		 for n > 0 {
			ans = ans * x
		}

		if n < 0 {
			ans = 1/ans
		}

	*/

	if x == 1.0 {
		return 1.0
	}

	if x == -1 {
		if n%2 == 0 {
			return 1.0
		} else {
			return -1.0
		}
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	ans := 1.0
	for n != 0 {
		if n%2 == 1 {
			ans = ans * x
			n -= 1
		}

		x = x * x
		n = n / 2
	}
	return ans
}

/*
78. Subsets
https://leetcode.com/problems/subsets/
*/
func GSubsets(nums []int) [][]int {

	/*
		approach1: add new element to current ans to generate new ans
		initial: []
		1. add 1
		[1]  ([]. [1])
		2, add 2
		[2], [1,2]  ([], [1], [2], [1,2])
		3 add 3
		[3] , [1,3], [2,3], [1,2,3]
	*/
	result := [][]int{{}}

	for i := 0; i < len(nums); i++ {

		currentSize := len(result)
		for j := 0; j < currentSize; j++ {
			temp := make([]int, len(result[j]))
			copy(temp[:], result[j][:])
			temp = append(temp, nums[i])
			result = append(result, temp)
		}

	}

	return result

	//result := make([][]int, 0)
	//currentSet := make([]int, 0)
	//dfsGSubsets(nums, 0, &currentSet, &result)
	//return result
}

func dfsGSubsets(nums []int, index int, currentSet *[]int, result *[][]int) {
	temp := make([]int, len(*currentSet))
	copy(temp[:], (*currentSet)[:])
	*result = append(*result, temp)
	for i := index; i < len(nums); i++ {
		*currentSet = append(*currentSet, nums[i])
		dfsGSubsets(nums, i+1, currentSet, result)
		*currentSet = (*currentSet)[:len(*currentSet)-1]
	}
}

/*
133. Clone Graph
https://leetcode.com/problems/clone-graph/
*/
func GCloneGraph(node *Node) *Node {
	/*
		approach1 BFS

		approach 2 recursive
	*/

	if node == nil {
		return nil
	}
	//nodeMap := make(map[int]*Node)
	//queue := []*Node{node}
	//
	//for len(queue) != 0 {
	//	currentNode := queue[0]
	//	queue = queue[1:]
	//
	//	var copyNode *Node
	//	if node, ok := nodeMap[currentNode.Val]; ok {
	//		copyNode = node
	//	} else {
	//		copyNode = &Node{Val: currentNode.Val, Neighbors: make([]*Node, 0)}
	//		nodeMap[currentNode.Val] = copyNode
	//	}
	//
	//	if len(currentNode.Neighbors) != len(copyNode.Neighbors) {
	//		queue = append(queue, currentNode.Neighbors...)
	//		for _, v := range currentNode.Neighbors {
	//			queue = append(queue, v)
	//			var neighbor *Node
	//			if node, ok := nodeMap[v.Val]; ok {
	//				neighbor = node
	//			} else {
	//				neighbor = &Node{Val: v.Val, Neighbors: make([]*Node, 0)}
	//				nodeMap[v.Val] = neighbor
	//			}
	//
	//			copyNode.Neighbors = append(copyNode.Neighbors, neighbor)
	//		}
	//	}
	//}
	//
	//return nodeMap[node.Val]
	return recursiveGCloneGraph(node, map[int]*Node{})
}

func recursiveGCloneGraph(node *Node, nodeMap map[int]*Node) *Node {
	if res, ok := nodeMap[node.Val]; ok {
		return res
	}

	copyNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	nodeMap[node.Val] = copyNode

	for _, neighbor := range node.Neighbors {
		copyNode.Neighbors = append(copyNode.Neighbors, recursiveGCloneGraph(neighbor, nodeMap))
	}

	return copyNode
}

func FlattenDictionary(dict map[string]interface{}) map[string]string {
	return flattenDict(dict, "")
}

func flattenDict(dict interface{}, prefix string) map[string]string {
	result := make(map[string]string)
	d, ok := dict.(map[string]interface{})

	if !ok {
		return nil
	}

	for k, v := range d {
		s, ok2 := v.(string)

		if ok2 {
			if prefix == "" {
				result[k] = s
			} else if k == "" {
				result[prefix] = s
			} else {
				result[prefix+"."+k] = s
			}
		} else {
			var next string
			if prefix == "" {
				next = k
			} else if k == "" {
				next = prefix
			} else {
				next = prefix + "." + k
			}

			for k2, v2 := range flattenDict(v, next) {
				result[k2] = v2
			}
		}

	}

	return result
}

/*
10. Regular Expression Matching
https://leetcode.com/problems/regular-expression-matching/description/
*/
type Result struct {
	value bool
}

func GIsMatch(text string, pattern string) bool {
	/* recursive

	if len(pattern) == 0 {
		return len(text) == 0
	}
	firstMatch := len(text) != 0 && (pattern[0] == text[0] || pattern[0] == '.')
	if len(pattern) >= 2 && pattern[1] == '*' {
		return (GIsMatch(text, pattern[2:])) || (firstMatch && GIsMatch(text[1:], pattern))
	} else {
		return firstMatch && GIsMatch(text[1:], pattern[1:])
	}
	*/

	/*
			DP:
			dp[i][j] = text[i:] match pattern[j:]


		memo := make([][]*Result, len(text)+1)
		for i := range memo {
			memo[i] = make([]*Result, len(pattern)+1)
		}

		return DPTopDownGIsMatch(0, 0, text, pattern, memo)

	*/

	/*
			DP: bottom up
		  a *  a  b a
		a
		a
		a
		b
		a f  f f  f  t
		  f  f f  f  f t
	*/

	dp := make([][]bool, len(text)+1)

	for i := range dp {
		dp[i] = make([]bool, len(pattern)+1)
	}

	dp[len(text)][len(pattern)] = true
	for i := len(text); i >= 0; i-- {
		for j := len(pattern) - 1; j >= 0; j-- {
			firstMatch := i < len(text) && (pattern[j] == text[i] || pattern[j] == '.')
			if j+1 < len(pattern) && pattern[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

func DPTopDownGIsMatch(i, j int, text, pattern string, memo [][]*Result) bool {
	if memo[i][j] != nil {
		return memo[i][j].value
	}

	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		firstMatch := i < len(text) && (text[i] == pattern[j] || pattern[j] == '.')

		if j+1 < len(pattern) && pattern[j+1] == '*' {
			ans = DPTopDownGIsMatch(i, j+2, text, pattern, memo) || (firstMatch && DPTopDownGIsMatch(i+1, j, text, pattern, memo))
		} else {
			ans = firstMatch && DPTopDownGIsMatch(i+1, j+1, text, pattern, memo)
		}
	}

	memo[i][j] = &Result{value: ans}
	return ans
}

/*

Serialize / deserialize binary tree

Search rotated array

Set columns and rows as zeros

Connect all siblings

Find all sum combinations

Clone a directed graph

Closest meeting point

Search for the given key in a 2d matrix

combination

permutation

*/
