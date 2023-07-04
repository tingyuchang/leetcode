package _0230628

import (
	__Daily_Prac "leetcode/0_Daily_Prac"
	"leetcode/LinkedList"
	"leetcode/Tree"
	"sort"
)

type Name struct {
}

// 1143. Longest Common Subsequence
/*
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
*/
func LongestCommonSubsequence(text1 string, text2 string) int {
	return 0
}

// 5. Longest Palindromic Substring
/*
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
*/
func LongestPalindrome(s string) string {
	return ""
}

// 516. Longest Palindromic Subsequence
/*
Input: s = "bbbab"
Output: 4
Explanation: One possible longest palindromic subsequence is "bbbb".
*/
func LongestPalindromeSubseq(s string) int {
	return 0
}

// 322. Coin Change
/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
func CoinChange(coins []int, amount int) int {
	return -1
}

// 1416. Restore The Array
/*
Input: s = "1317", k = 2000
Output: 8
Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
*/
func NumberOfArrays(s string, k int) int {
	return 0
}

// 72. Edit Distance
/*
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
*/
func MinDistance(word1 string, word2 string) int {
	return 0
}

// 139. Word Break
/*
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func WordBreak(s string, wordDict []string) bool {
	return false
}

// 127. Word Ladder
/*
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
*/
func LadderLength(beginWord string, endWord string, wordList []string) int {
	return 0
}

// 221. Maximal Square
/*
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
*/
func MaximalSquare(matrix [][]byte) int {
	return 0
}

// 300. Longest Increasing Subsequence
/*
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
func LengthOfLIS(nums []int) int {
	return 1
}

// 122. Best Time to Buy and Sell Stock II
/*
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.
*/
func MaxProfit(prices []int) int {
	return 0
}

// 45. Jump Game II
/*
Return the minimum number of jumps to reach nums[n - 1].
The test cases are generated such that you can reach nums[n - 1].

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
Jump 1 step from index 0 to 1, then 3 steps to the last index.
*/

func JumpII(nums []int) int {
	return 0
}

// 134. Gas Station
/*
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
*/

func CanCompleteCircuit(gas []int, cost []int) int {
	return -1
}

// 135. Candy
/*
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
*/
func Candy(ratings []int) int {
	return 0
}

// 42. Trapping Rain Water
/*
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
*/

func Trap(height []int) int {
	return 0
}

// 209. Minimum Size Subarray Sum
/*
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
*/
func MinSubArrayLen(target int, nums []int) int {
	return 0
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
/*
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *Tree.TreeNode {
	return nil
}

// 117. Populating Next Right Pointers in Each Node II
/*
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer
to point to its next right node, just like in Figure B. The serialized output is in level order
as connected by the next pointers, with '#' signifying the end of each level.
*/

func ConnectTreeNode(root *Tree.NodeN) *Tree.NodeN {
	return nil
}

// 1498. Number of Subsequences That Satisfy the Given Sum Condition
/*
You are given an array of integers nums and an integer target.

Return the number of non-empty subsequences of nums such that the sum of the minimum and maximum element
on it is less or equal to target. Since the answer may be too large, return it modulo 109 + 7.

Input: nums = [3,5,6,7], target = 9
Output: 4
Explanation: There are 4 subsequences that satisfy the condition.
[3] -> Min value + max value <= target (3 + 3 <= 9)
[3,5] -> (3 + 5 <= 9)
[3,5,6] -> (3 + 6 <= 9)
[3,6] -> (3 + 6 <= 9)

*/

func NumSubseq(nums []int, target int) int {
	return 0
}

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/description/

func SwapPairs(head *LinkedList.ListNode) *LinkedList.ListNode {

	return nil
}

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&id=top-interview-150

func RemoveNthFromEnd(head *LinkedList.ListNode, n int) *LinkedList.ListNode {
	return nil
}

// 92. Reverse Linked List II

func ReverseBetween(head *LinkedList.ListNode, left int, right int) *LinkedList.ListNode {
	return nil
}

// 215. Kth Largest Element in an Array

func FindKthLargest(nums []int, k int) int {
	return 0
}

// 22. Generate Parentheses

func GenerateParenthesis(n int) []string {
	return nil
}

// 56. Merge Intervals

func MergeIntervals(intervals [][]int) [][]int {
	return nil
}

// 15. 3Sum
func ThreeSum(nums []int) [][]int {
	return nil
}

// 151. Reverse Words in a String

func ReverseWords(s string) string {
	return ""
}

// 91. Decode Ways

func NumDecodings(s string) int {
	return 0
}

// 394. Decode String

func DecodeString(s string) string {
	return ""
}

// 78. Subsets
func Subsets(nums []int) [][]int {
	return nil
}

// 133. Clone Graph

func CloneGraph(node *LinkedList.Node) *LinkedList.Node {
	return nil
}

// 10. Regular Expression Matching

func IsMatch(text string, pattern string) bool {
	return true
}

// word count engine
/*
input:  document = "Practice makes perfect. you'll only
                    get Perfect by practice. just practice!"

output: [ ["practice", "3"], ["perfect", "2"],
          ["makes", "1"], ["youll", "1"], ["only", "1"],
          ["get", "1"], ["by", "1"], ["just", "1"] ]

*/

func WordCountEngine(document string) [][]string {
	return nil
}

/*

 */

func MinimizeDistanceToFarthestPoint(blocks [][]bool, requires int) int {
	return 0
}

/*
37. Sudoku Solver

https://leetcode.com/problems/sudoku-solver/
*/
func SolveSudoku(board [][]byte) {
}

/*
33. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
func SearchRotated(nums []int, target int) int {
	if nums[0] < nums[len(nums)-1] || len(nums) == 1 {
		return binarySearch(nums, target)
	}

	l := 0
	r := len(nums) - 1
	pivot := 0

	for r >= l {
		mid := (l + r) / 2

		if nums[mid] > nums[mid+1] {
			pivot = mid + 1
			break
		}

		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
	}

	if nums[len(nums)-1] >= target {
		result := binarySearch(nums[pivot:], target)

		if result == -1 {
			return -1
		} else {
			return pivot + result
		}
	} else {
		return binarySearch(nums[:pivot], target)
	}

	return -1
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := (l + r) / 2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

/*
73. Set Matrix Zeroes
https://leetcode.com/problems/set-matrix-zeroes/
*/
func SetZeroes(matrix [][]int) {
	needSetToZero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			needSetToZero = true
		}

		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 1; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if needSetToZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}

/*
39. Combination Sum
https://leetcode.com/problems/combination-sum/
*/
func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	dfsCombinationSum(candidates, target, 0, 0, &[]int{}, &result)
	return result
}

func dfsCombinationSum(candidates []int, target, start, currentSum int, currentResult *[]int, result *[][]int) {
	if currentSum == target {
		temp := make([]int, len(*currentResult))
		copy(temp, *currentResult)
		*result = append(*result, temp)
	}

	if currentSum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		*currentResult = append(*currentResult, candidates[i])
		currentSum += candidates[i]
		dfsCombinationSum(candidates, target, i, currentSum, currentResult, result)
		*currentResult = (*currentResult)[:len(*currentResult)-1]
		currentSum -= candidates[i]
	}
}

/*
116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
*/

func ConnectAllSiblings(root *Tree.NodeN) *Tree.NodeN {
	if root == nil {
		return nil
	}

	current := root
	for current != nil {
		node := current
		for node != nil {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right = node.Next.Left
			}

			node = node.Next
		}
		current = current.Left
	}
	return root
}

/*
296. Best Meeting Point

A group of two or more people wants to meet and minimize the total travel distance.
You are given a 2D grid of values 0 or 1, where each 1 marks the home of someone in the group.
The distance is calculated using Manhattan Distance,
where distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|.

input:

	1 0 0 0 1
	0 0 0 0 0
	0 0 1 0 0

output: 6

[0,2] is ideal meeting point, so 2 + 2 + 2 = 6
*/
func BestMeetingPoint(grid [][]int) int {
	cols := make([]int, 0)
	rows := make([]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	sort.Ints(rows)
	sort.Ints(cols)

	targetRow := rows[len(rows)/2]
	targetCol := cols[len(cols)/2]

	distance := 0

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 0; i < len(rows); i++ {
		distance += abs(rows[i]-targetRow) + abs(cols[i]-targetCol)
	}

	return distance
}

/*
74. Search a 2D Matrix
https://leetcode.com/problems/search-a-2d-matrix/
*/
func SearchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := len(matrix) - 1

	for r >= l {
		mid := (l + r) / 2

		if matrix[mid][0] < target {
			l = mid + 1
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			return true
		}
	}

	k := r

	if k < 0 {
		k = 0
	}

	l = 0
	r = len(matrix[k]) - 1

	for r >= l {
		mid := (l + r) / 2

		if matrix[k][mid] < target {
			l = mid + 1
		} else if matrix[k][mid] > target {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}

/*
77. Combinations
https://leetcode.com/problems/combinations/
*/
func Combine(n int, k int) [][]int {
	result := make([][]int, 0)
	dfsCombine(n, k, 1, &[]int{}, &result)
	return result
}

func dfsCombine(n, k, start int, currentResult *[]int, result *[][]int) {
	if len(*currentResult) == k {
		temp := make([]int, len(*currentResult))
		copy(temp, *currentResult)
		*result = append(*result, temp)
		return
	}

	for i := start; i <= n; i++ {
		*currentResult = append(*currentResult, i)
		dfsCombine(n, k, i+1, currentResult, result)
		*currentResult = (*currentResult)[:len(*currentResult)-1]
	}

}

/*
31. Next Permutation
https://leetcode.com/problems/next-permutation/
*/
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 1

	reverse := func(nums []int, n int) {
		l := n
		r := len(nums) - 1

		for r > l {
			nums[l], nums[r] = nums[r], nums[l]
			l += 1
			r -= 1
		}
	}

	for nums[i-1] > nums[i] {
		i -= 1
		if i == 0 {
			reverse(nums, 0)
			return
		}
	}
	j := len(nums) - 1

	for j < i && nums[j] < nums[i-1] {
		j -= 1
	}

	nums[i-1], nums[j] = nums[j], nums[i-1]
	reverse(nums, i)
}

/*
347. Top K Frequent Elements

https://leetcode.com/problems/top-k-frequent-elements/

1. sort
2. hepp
3. bucket sort
4. random select
*/
func TopKFrequent(nums []int, k int) []int {
	buckets := make([][]int, 0)
	frequentMap := make(map[int]int)

	for _, v := range nums {
		frequentMap[v] += 1
		buckets = append(buckets, []int{})
	}

	for val, frequent := range frequentMap {
		buckets[frequent-1] = append(buckets[frequent-1], val)
	}

	ans := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		ans = append(ans, buckets[i]...)
		if len(ans) >= k {
			break
		}
	}

	return ans
}

/*
402. Remove K Digits
https://leetcode.com/problems/remove-k-digits/

*/

func RemoveKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	stack := make([]byte, 0)

	for i := 0; i < len(num); i++ {

		for len(stack) != 0 && k > 0 && stack[len(stack)-1] > num[i] {
			k -= 1
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 || num[i] != '0' {
			stack = append(stack, num[i])
		}
	}

	if k > 0 {
		if k > len(stack) {
			return "0"
		} else {
			stack = stack[:len(stack)-k]
		}
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}

func EmployeeFreeTime(schedule [][][]int) [][]int {
	intervals := make([][]int, 0)
	for i := 0; i < len(schedule); i++ {
		for j := 0; j < len(schedule[i]); j++ {
			intervals = append(intervals, schedule[i][j])
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if current[1] < next[0] {
			result = append(result, []int{current[1], next[0]})
			current = next
		} else {
			if next[1] > current[1] {
				current = next
			}
		}

	}
	return result
}

/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/
*/

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		currentMax = __Daily_Prac.Max(nums[i], currentMax+nums[i])
		ans = __Daily_Prac.Max(ans, currentMax)
	}
	return ans
}

/*
918. Maximum Sum Circular Subarray
https://leetcode.com/problems/maximum-sum-circular-subarray/description/
*/
func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)
	rightMax := make([]int, len(nums))
	rightMax[n-1] = nums[n-1]
	suffixSum := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		suffixSum += nums[i]
		rightMax[i] = __Daily_Prac.Max(suffixSum, rightMax[i+1])
	}

	current, ans := 0, nums[0]
	specialSum := nums[0]
	prefixSum := 0

	for i := 0; i < n; i++ {
		current = __Daily_Prac.Max(current, 0) + nums[i]
		ans = __Daily_Prac.Max(ans, current)

		prefixSum += nums[i]
		if i < n-1 {
			specialSum = __Daily_Prac.Max(specialSum, prefixSum+rightMax[i+1])
		}
	}

	return __Daily_Prac.Max(specialSum, ans)
}

/*
978. Longest Turbulent Subarray
https://leetcode.com/problems/longest-turbulent-subarray/description/
*/
func MaxTurbulenceSize(arr []int) int {
	compare := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		} else {
			return 0
		}
	}

	l := 0
	ans := 1

	for i := 1; i < len(arr); i++ {
		c := compare(arr[i-1], arr[i])

		if c == 0 {
			l = i
		} else if i == len(arr)-1 || c*compare(arr[i], arr[i+1]) != -1 {
			// if 0 or +1 than we should record temp ans here
			ans = __Daily_Prac.Max(ans, i-l+1)
			l = i
		}
	}

	return ans
}
