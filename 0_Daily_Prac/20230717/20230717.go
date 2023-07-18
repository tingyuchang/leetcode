package _0230717

import (
	__Daily_Prac "leetcode/0_Daily_Prac"
	"leetcode/LinkedList"
	"leetcode/Tree"
	"math"
	"sort"
	"strconv"
	"strings"
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
	result := make([][]int, len(text1)+1)
	for i := range result {
		result[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				result[i][j] = result[i-1][j-1] + 1
			} else {
				result[i][j] = __Daily_Prac.Max(result[i-1][j], result[i][j-1])
			}
		}
	}
	return result[len(text1)][len(text2)]
}

// 5. Longest Palindromic Substring
/*
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
*/
func LongestPalindrome(s string) string {
	start, end := 0, -1

	for i := 0; i < len(s); i++ {
		ans1 := expand(s, i, i)
		ans2 := expand(s, i, i+1)
		ans := __Daily_Prac.Max(ans1, ans2)

		if ans > end-start+1 {
			start = i - (ans-1)/2
			end = i + ans/2
		}
	}
	return s[start : end+1]
}

func expand(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start -= 1
		end += 1
	}

	return end - start + 1 - 2
}

// 516. Longest Palindromic Subsequence
/*
Input: s = "bbbab"
Output: 4
Explanation: One possible longest palindromic subsequence is "bbbb".
*/
func LongestPalindromeSubseq(s string) int {
	result := make([]int, len(s))
	resultPre := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		result[i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				result[j] = resultPre[j-1] + 2
			} else {
				result[j] = __Daily_Prac.Max(result[j-1], resultPre[j])
			}
		}

		copy(resultPre, result)
	}
	return result[len(s)-1]
}

// 322. Coin Change
/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
func CoinChange(coins []int, amount int) int {
	result := make([]int, amount+1)
	sort.Ints(coins)

	// i means money's val
	for i := 1; i <= amount; i++ {
		// mark init val to inf
		result[i] = math.MaxInt
		for _, c := range coins {
			// if current coin > val, break looop
			if i < c {
				break
			}

			if result[i-c] != math.MaxInt {
				result[i] = __Daily_Prac.Min(result[i], result[i-c]+1)
			}
		}
	}

	if result[amount] != math.MaxInt {
		return result[amount]
	}
	return -1
}

// 1416. Restore The Array
/*
Input: s = "1317", k = 2000
Output: 8
Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
*/
func NumberOfArrays(s string, k int) int {
	result := make([]int, len(s)+1)
	result[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}

		for j := i; j < len(s); j++ {
			x, _ := strconv.Atoi(s[i : j+1])
			if x > k {
				break
			}
			result[j+1] = (result[j+1] + result[i]) % __Daily_Prac.MOD
		}
	}
	return result[len(s)]
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
	result := make([][]int, len(word1)+1)
	for i := range result {
		result[i] = make([]int, len(word2)+1)

		if i == 0 {
			for j := 0; j < len(result[i]); j++ {
				result[i][j] = j
			}
		} else {
			result[i][0] = i
		}
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				result[i][j] = result[i-1][j-1]
			} else {
				result[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(result[i-1][j], result[i][j-1]), result[i-1][j-1]) + 1
			}
		}
	}

	return result[len(word1)][len(word2)]
}

// 139. Word Break
/*
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func WordBreak(s string, wordDict []string) bool {
	result := make([]bool, len(s)+1)
	result[0] = true

	wordMap := make(map[string]struct{})

	for _, v := range wordDict {
		wordMap[v] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordMap[s[j:i]]; ok && result[j] {
				result[i] = true
				break
			}
		}
	}

	return result[len(s)]

}

// 127. Word Ladder
/*
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
*/
func LadderLength(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	visitedWords := make(map[string]bool)

	for _, v := range wordList {
		visitedWords[v] = false
	}
	visitedWords[beginWord] = true
	step := 0

	for len(queue) != 0 {
		currentSize := len(queue)
		step += 1
		for i := 0; i < currentSize; i++ {
			currentWord := queue[i]
			for word, visited := range visitedWords {
				if !visited {
					count := 0
					for j := 0; j < len(word); j++ {
						if word[j] != currentWord[j] {
							count += 1
						}

						if count > 1 {
							break
						}
					}

					if count == 1 {
						if word == endWord {
							return step + 1
						}

						queue = append(queue, word)
						visitedWords[word] = true
					}
				}
			}
		}

		queue = queue[currentSize:]
	}

	return 0
}

// 221. Maximal Square
/*
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
*/
func MaximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, m+1)
	for i := range result {
		result[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				result[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(result[i-1][j], result[i][j-1]), result[i-1][j-1]) + 1
				if ans < result[i][j] {
					ans = result[i][j]
				}
			}
		}
	}

	return ans * ans
}

// 300. Longest Increasing Subsequence
/*
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
func LengthOfLIS(nums []int) int {
	subArray := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(subArray) == 0 || nums[i] > subArray[len(subArray)-1] {
			subArray = append(subArray, nums[i])
		} else {
			l := 0
			r := len(subArray) - 1

			for r >= l {
				mid := int(uint(l+r) >> 1)
				if subArray[mid] < nums[i] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			subArray[l] = nums[i]
		}

	}

	return len(subArray)

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
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
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
	maxDistance := 0
	currentEnd := 0
	step := 0
	for i := 0; i < len(nums); i++ {
		maxDistance = __Daily_Prac.Max(maxDistance, i+nums[i])

		if currentEnd == i {
			step += 1
			currentEnd = maxDistance
			if currentEnd >= len(nums)-1 {
				return step
			}
		}
	}
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
	totalTank := 0
	currentTank := 0
	station := 0

	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]
		if currentTank < 0 {
			currentTank = 0
			station = i + 1
		}
	}

	if totalTank >= 0 {
		return station
	}
	return -1
}

// 135. Candy
/*
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
*/
func Candy(ratings []int) int {
	candies := make([]int, len(ratings))
	candies[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	sum := candies[len(candies)-1]

	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = candies[i+1] + 1
		}
		sum += candies[i]
	}

	return sum
}

// 42. Trapping Rain Water
/*
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
*/

func Trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}

		if height[n-i-1] > rightMax[n-i] {
			rightMax[n-i-1] = height[n-i-1]
		} else {
			rightMax[n-i-1] = rightMax[n-i]
		}
	}

	water := 0

	for i := 1; i < n-1; i++ {
		water += __Daily_Prac.Min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

// 209. Minimum Size Subarray Sum
/*
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
*/
func MinSubArrayLen(target int, nums []int) int {
	l, r, sum, ans := 0, 0, 0, 0

	for r < len(nums) {
		sum += nums[r]

		for sum >= target {
			if ans == 0 || r-l+1 < ans {
				ans = r - l + 1
			}

			sum -= nums[l]
			l += 1
		}

		r += 1
	}

	return ans
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
/*
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *Tree.TreeNode {
	inordrMap := make(map[int]int)

	for i, v := range inorder {
		inordrMap[v] = i
	}
	index := 0
	return constructTree(preorder, 0, len(preorder)-1, &index, inordrMap)
}

func constructTree(preorder []int, start, end int, index *int, inorderMap map[int]int) *Tree.TreeNode {
	if *index >= len(preorder) {
		return nil
	}

	rootVal := preorder[*index]
	*index += 1
	root := &Tree.TreeNode{Val: rootVal}
	pivot := inorderMap[rootVal]

	if pivot-start > 0 {
		root.Left = constructTree(preorder, start, pivot-1, index, inorderMap)
	}
	if end-pivot > 0 {
		root.Right = constructTree(preorder, pivot+1, end, index, inorderMap)
	}

	return root
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
	current := root
	var head, next *Tree.NodeN
	for current != nil {

		for current != nil {
			if current.Left != nil {
				if next != nil {
					next.Next = current.Left
				} else {
					head = current.Left
				}

				next = current.Left
			}

			if current.Right != nil {
				if next != nil {
					next.Next = current.Right
				} else {
					head = current.Right
				}

				next = current.Right
			}

			current = current.Next
		}

		current = head

		head, next = nil, nil
	}

	return root
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

	power := make([]int, len(nums))
	power[0] = 1

	for i := 1; i < len(power); i++ {
		power[i] = (power[i-1] * 2) % __Daily_Prac.MOD
	}

	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		l := 0
		r := len(nums) - 1

		for r >= l {
			mid := int(uint(l+r) >> 1)

			if nums[mid] <= target-nums[i] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		if r >= i {
			ans = (ans + power[r-i]) % __Daily_Prac.MOD
		}
	}

	return ans
}

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/description/

func SwapPairs(head *LinkedList.ListNode) *LinkedList.ListNode {
	preHead := &LinkedList.ListNode{}
	preHead.Next = head
	pre := preHead

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		pre.Next = next
		pre = head
		head = head.Next
	}

	return preHead.Next
}

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&id=top-interview-150

func RemoveNthFromEnd(head *LinkedList.ListNode, n int) *LinkedList.ListNode {
	preHead := &LinkedList.ListNode{}
	preHead.Next = head
	pre := preHead
	move := n - 1
	// 0 1 2 3 4 5, n = 2
	// move = 1
	// head := 2

	for move > 0 {
		head = head.Next
		move -= 1
	}
	// h: 2, p: 0
	// h: 3, p: 1
	// h: 4, p: 2
	// h: 5, p: 3

	for head.Next != nil {
		head = head.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next

	return preHead.Next
}

// 92. Reverse Linked List II

func ReverseBetween(head *LinkedList.ListNode, left int, right int) *LinkedList.ListNode {
	preFirst := &LinkedList.ListNode{}
	preFirst.Next = head
	move := left - 1

	first := head
	for move > 0 {
		first = first.Next
		preFirst = preFirst.Next
		move -= 1
	}

	var start, next *LinkedList.ListNode
	end := first
	move = right - left + 1
	for move > 0 {
		next = first.Next
		first.Next = start
		start = first
		first = next
		move -= 1
	}
	end.Next = next
	preFirst.Next = start
	if left == 1 {
		return preFirst
	}
	return head
}

// 215. Kth Largest Element in an Array
/*
1. sort and interating reverse => nlogn
2. random select -> n
*/

func FindKthLargest(nums []int, k int) int {
	return randomSelect(nums, 0, len(nums)-1, len(nums)-k+1)
}

func randomSelect(nums []int, p, r, target int) int {
	if p == r {
		return nums[p]
	}
	q := partition(nums, p, r)

	k := q - p + 1

	if k == target {
		return nums[q]
	} else if k < target {
		return randomSelect(nums, q+1, r, target-k)
	} else {
		return randomSelect(nums, p, q-1, target)
	}
}

func partition(nums []int, p, r int) int {
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

// 22. Generate Parentheses

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	GenerateParenthesisHelper("", 0, 0, n, &result)
	return result
}

func GenerateParenthesisHelper(s string, leftCount, rightCount, n int, result *[]string) {
	if len(s) == 2*n {
		*result = append(*result, s)
		return
	}

	if leftCount < n {
		GenerateParenthesisHelper(s+"(", leftCount+1, rightCount, n, result)
	}
	if rightCount < leftCount {
		GenerateParenthesisHelper(s+")", leftCount, rightCount+1, n, result)
	}
}

// 56. Merge Intervals

func MergeIntervals(intervals [][]int) [][]int {
	starts := make([]int, 0)
	endMap := make(map[int]int)

	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		if existing, ok := endMap[start]; ok {
			if existing < end {
				endMap[start] = end
			}
		} else {
			endMap[start] = end
			starts = append(starts, start)
		}
	}

	sort.Ints(starts)

	current := starts[0]
	result := make([][]int, 0)
	for i := 1; i < len(starts); i++ {
		next := starts[i]
		if endMap[current] >= next && endMap[current] < endMap[next] {
			endMap[current] = endMap[next]
		} else if endMap[current] < next {
			result = append(result, []int{current, endMap[current]})
			current = next
		}
	}
	result = append(result, []int{current, endMap[current]})

	return result
}

// 15. 3Sum
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for r > l {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l += 1

				for l < r && nums[l] == nums[l-1] {
					l += 1
				}
			}
		}
	}
	return result
}

// 151. Reverse Words in a String

func ReverseWords(s string) string {
	result := ""
	l, r := len(s)-1, len(s)-1

	for l >= 0 {
		if s[l] == ' ' {
			if r > l {
				if len(result) != 0 {
					result += " "
				}
				result += s[l+1 : r+1]
			}
			r = l - 1
		} else if l == 0 {
			if s[l:r+1] != " " {
				if len(result) != 0 {
					result += " "
				}
				result += s[l : r+1]
			}
		}

		l -= 1
	}
	return result
}

// 91. Decode Ways

func NumDecodings(s string) int {
	result := make([]int, len(s)+1)
	result[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			result[i] += result[i+1]

			if i < len(s)-1 && (s[i] == '1' || s[i] == '2' && s[i+1] < '7') {
				result[i] += result[i+2]
			}

		}
	}
	return result[0]
}

// 394. Decode String

func DecodeString(s string) string {
	stack := make([]byte, 0)
	numStack := make([]int, 0)
	openBracketStack := make([]int, 0)
	currentNum := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x, _ := strconv.Atoi(string(s[i]))
			currentNum = currentNum*10 + x
		} else if s[i] == '[' {
			numStack = append(numStack, currentNum)
			currentNum = 0
			openBracketStack = append(openBracketStack, len(stack))
		} else if s[i] == ']' {
			duplicatedCount := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			strIndex := openBracketStack[len(openBracketStack)-1]
			openBracketStack = openBracketStack[:len(openBracketStack)-1]

			currentStr := string(stack[strIndex:])

			for j := 0; j < duplicatedCount-1; j++ {
				stack = append(stack, currentStr...)
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

// 78. Subsets
func Subsets(nums []int) [][]int {
	result := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		currentSize := len(result)
		for j := 0; j < currentSize; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			temp = append(temp, nums[i])
			result = append(result, temp)
		}
	}
	return result
}

// 133. Clone Graph

func CloneGraph(node *LinkedList.Node) *LinkedList.Node {
	return nil
}

// 10. Regular Expression Matching
type Result struct {
	val bool
}

func IsMatch(text string, pattern string) bool {
	memo := make([][]*Result, len(text)+1)
	for i := range memo {
		memo[i] = make([]*Result, len(pattern)+1)
	}
	return IsMatchHelper(text, pattern, 0, 0, memo)
}

func IsMatchHelper(text, pattern string, i, j int, memo [][]*Result) bool {
	if memo[i][j] != nil {
		return memo[i][j].val
	}

	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		firstMatch := i < len(text) && (text[i] == pattern[j] || pattern[j] == '.')

		if j < len(pattern)-1 && pattern[j+1] == '*' {
			ans = firstMatch && IsMatchHelper(text, pattern, i+1, j, memo) || IsMatchHelper(text, pattern, i, j+2, memo)
		} else {
			ans = firstMatch && IsMatchHelper(text, pattern, i+1, j+1, memo)
		}
	}

	memo[i][j] = &Result{val: ans}
	return memo[i][j].val
}

// word count engine
/*
input:  document = "Practice makes perfect. you'll only
                    get Perfect by practice. just practice!"

output: [ ["practice", "3"], ["perfect", "2"],
          ["makes", "1"], ["youll", "1"], ["only", "1"],
          ["get", "1"], ["by", "1"], ["just", "1"] ]

*/

type WordCount struct {
	word  string
	index int
	count int
}

func WordCountEngine(document string) [][]string {
	unprocessedWord := strings.Split(document, " ")
	countMap := make(map[string]*WordCount)

	for i, unprocessed := range unprocessedWord {
		lowerCase := strings.ToLower(unprocessed)
		word := make([]byte, 0)

		for j := 0; j < len(lowerCase); j++ {
			if lowerCase[j] >= 'a' && lowerCase[j] <= 'z' {
				word = append(word, lowerCase[j])
			}
		}

		if _, ok := countMap[string(word)]; ok {
			countMap[string(word)].count += 1
		} else {
			countMap[string(word)] = &WordCount{
				word:  string(word),
				index: i,
				count: 1,
			}
		}
	}

	temp := make([]*WordCount, len(countMap))
	index := 0

	for _, v := range countMap {
		temp[index] = v
		index += 1
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].count == temp[j].count {
			return temp[i].index < temp[j].index
		}
		return temp[i].count > temp[j].count
	})

	result := make([][]string, len(temp))
	for i, wordCount := range temp {
		result[i] = []string{wordCount.word, strconv.Itoa(wordCount.count)}
	}

	return result
}

/*

 */

func MinimizeDistanceToFarthestPoint(blocks [][]bool, requires int) int {
	distances := make([][]int, len(blocks))

	for i := range distances {
		// last element is maxnium
		distances[i] = make([]int, requires+1)

		for j := 0; j < requires; j++ {
			distances[i][j] = math.MaxInt
		}
	}

	for i := 0; i < len(blocks); i++ {
		for j := 0; j < requires; j++ {
			if blocks[i][j] {
				distances[i][j] = 0
			} else if i > 0 && distances[i-1][j] != math.MaxInt {
				distances[i][j] = __Daily_Prac.Min(distances[i][j], distances[i-1][j]+1)
			}
			if distances[i][j] != math.MaxInt && distances[i][j] > distances[i][requires] {
				distances[i][requires] = distances[i][j]
			}
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		for j := 0; j < requires; j++ {
			if blocks[i][j] {
				distances[i][j] = 0
			} else if i < len(blocks)-1 && distances[i+1][j] != math.MaxInt {
				distances[i][j] = __Daily_Prac.Min(distances[i][j], distances[i+1][j]+1)
			}
			if distances[i][j] != math.MaxInt && distances[i][j] > distances[i][requires] {
				distances[i][requires] = distances[i][j]
			}
		}
	}

	ans := 0
	maxDistance := math.MaxInt

	for i := 0; i < len(distances); i++ {
		if distances[i][requires] < maxDistance {
			maxDistance = distances[i][requires]
			ans = i
		}
	}

	return ans
}

/*
37. Sudoku Solver

https://leetcode.com/problems/sudoku-solver/
*/
func SolveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := 1; k <= 9; k++ {
					x := strconv.Itoa(k)[0]
					if validate(board, i, j, x) {
						board[i][j] = x

						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}

	return true
}

func validate(board [][]byte, x, y int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[x][i] == c {
			return false
		}

		if board[i][y] == c {
			return false
		}

		if board[3*(x/3)+i/3][3*(y/3)+i%3] == c {
			return false
		}
	}
	return true
}

/*
33. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
func SearchRotated(nums []int, target int) int {
	n := len(nums)
	if n <= 1 || nums[0] < nums[n-1] {
		return binarySearch(nums, target)
	}

	l := 0
	r := len(nums) - 1
	pivot := 0

	for r >= l {
		mid := int(uint(l+r) >> 1)
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

	if target <= nums[n-1] {
		result := binarySearch(nums[pivot:], target)
		if result == -1 {
			return result
		}

		return result + pivot
	} else {
		return binarySearch(nums[:pivot], target)
	}
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)
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
	setToZero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			setToZero = true
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
			if matrix[i][j] != 0 && matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 1; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if setToZero {
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
	CombinationSumHelper(candidates, 0, target, 0, &[]int{}, &result)
	return result
}

func CombinationSumHelper(candidate []int, start, target, currentSum int, currentResult *[]int, result *[][]int) {
	if currentSum == target {
		temp := make([]int, len(*currentResult))
		copy(temp, *currentResult)
		*result = append(*result, temp)
		return
	}

	if currentSum > target {
		return
	}

	for i := start; i < len(candidate); i++ {
		currentSum += candidate[i]
		*currentResult = append(*currentResult, candidate[i])
		CombinationSumHelper(candidate, i, target, currentSum, currentResult, result)
		currentSum -= candidate[i]
		*currentResult = (*currentResult)[:len(*currentResult)-1]
	}
}

/*
116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
*/

func ConnectAllSiblings(root *Tree.NodeN) *Tree.NodeN {
	current := root

	for current != nil {

		for current != nil {
			current.Left.Next = current.Right
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			current = current.Next
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
	rows := make([]int, 0)
	cols := make([]int, 0)

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

	ans := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := 0; i < len(rows); i++ {
		ans += abs(rows[i]-targetRow) + abs(cols[i]-targetCol)
	}

	return ans
}

/*
74. Search a 2D Matrix
https://leetcode.com/problems/search-a-2d-matrix/
*/
func SearchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := len(matrix) - 1
	slot := 0

	for r >= l {
		mid := int(uint(l+r) >> 1)
		if matrix[mid][0] < target {
			l = mid + 1
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			return true
		}
	}

	if r > 0 {
		slot = r
	}

	l = 0
	r = len(matrix[slot])
	for r >= l {
		mid := int(uint(l+r) >> 1)

		if matrix[slot][mid] < target {
			l = mid + 1
		} else if matrix[slot][mid] > target {
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
	CombinationHelper(1, k, n, &[]int{}, &result)
	return result
}

func CombinationHelper(start, k, n int, currentResult *[]int, result *[][]int) {
	if len(*currentResult) == k {
		temp := make([]int, len(*currentResult))
		copy(temp, *currentResult)
		*result = append(*result, temp)
		return
	}

	if len(*currentResult) > k {
		return
	}

	for i := start; i <= n; i++ {
		*currentResult = append(*currentResult, i)
		CombinationHelper(i+1, k, n, currentResult, result)
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
	reverse := func(n int) {
		l := n
		r := len(nums) - 1
		for r > l {
			nums[l], nums[r] = nums[r], nums[l]
			l += 1
			r -= 1
		}
	}

	i := len(nums) - 1
	for nums[i-1] > nums[i] {
		i -= 1
		if i == 0 {
			reverse(0)
			return
		}
	}

	j := len(nums) - 1
	for j < i && nums[j] <= nums[i-1] {
		j -= 1
	}
	nums[i-1], nums[j] = nums[j], nums[i-1]
	reverse(i)
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
	bucket := make([][]int, 1)
	freqMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := freqMap[nums[i]]; ok {
			freqMap[nums[i]] += 1
		} else {
			freqMap[nums[i]] = 1
			bucket = append(bucket, []int{})
		}
	}

	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	result := make([]int, 0)

	for i := len(bucket) - 1; i >= 0; i-- {
		result = append(result, bucket[i]...)
		if len(result) == k {
			break
		}
	}

	return result
}

/*
402. Remove K Digits
https://leetcode.com/problems/remove-k-digits/

*/

func RemoveKdigits(num string, k int) string {
	stack := make([]byte, 0)

	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			k -= 1
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 || num[i] != '0' {
			stack = append(stack, num[i])
		}
	}

	if k > 0 {
		if k >= len(stack) {
			return "0"
		}

		stack = stack[:len(stack)-k]
	}

	if len(stack) != 0 {
		return string(stack)
	}

	return "0"
}

/*
Employee Free Time
Input:
 schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]

Output:
 [[3,4]]

Explanation:

There are a total of three employees, and all common
free time intervals would be [-inf, 1], [3, 4], [10, inf].
We discard any intervals that contain inf as they aren't finite.
*/

func EmployeeFreeTime(schedule [][][]int) [][]int {
	intervals := make([][]int, 0)

	for _, person := range schedule {
		for _, interval := range person {
			intervals = append(intervals, interval)
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
		} else if current[1] < next[1] {
			current = next
		}
	}
	return result
}

/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/
*/

func MaxSubArray(nums []int) int {
	current := 0
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		current = __Daily_Prac.Max(0, current) + nums[i]
		ans = __Daily_Prac.Max(ans, current)
	}

	return ans
}

/*
918. Maximum Sum Circular Subarray
https://leetcode.com/problems/maximum-sum-circular-subarray/description/
*/
func MaxSubarraySumCircular(nums []int) int {
	rightMax := make([]int, len(nums))
	suffix := nums[len(nums)-1]
	rightMax[len(rightMax)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		suffix += nums[i]
		rightMax[i] = __Daily_Prac.Max(rightMax[i+1], suffix)
	}

	current := 0
	normalSum := nums[0]
	specialSum := 0
	prefix := 0

	for i := 0; i < len(nums); i++ {
		current = __Daily_Prac.Max(0, current) + nums[i]
		normalSum = __Daily_Prac.Max(normalSum, current)

		prefix += nums[i]

		if i < len(nums)-1 {
			specialSum = __Daily_Prac.Max(specialSum, prefix+rightMax[i+1])
		}
	}

	return __Daily_Prac.Max(normalSum, specialSum)
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

	anchor := 0
	ans := 1

	for i := 1; i < len(arr); i++ {
		c := compare(arr[i-1], arr[i])

		if c == 0 {
			anchor = i
		} else if i == len(arr)-1 || c*compare(arr[i], arr[i+1]) != -1 {
			if i-anchor+1 > ans {
				ans = i - anchor + 1
			}
			anchor = i
		}
	}

	return ans
}

/*
28. Find the Index of the First Occurrence in a String
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
please using KMP to resolve this question
*/

func StrStr(haystack string, needle string) int {
	lps := make([]int, len(needle))

	prefix := 0
	i := 1
	for i < len(needle) {
		if needle[i] == needle[prefix] {
			prefix += 1
			lps[i] = prefix
			i += 1
		} else if prefix > 0 {
			prefix = lps[prefix-1]
		} else {
			i += 1
		}
	}

	i = 0
	j := 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i += 1
			j += 1
		} else {
			if j == 0 {
				i += 1
			} else {
				j = lps[j-1]
			}
		}

		if j == len(needle) {
			return i - len(needle)
		}
	}

	return -1
}

/*
200. Number of Islands
https://leetcode.com/problems/number-of-islands/
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
*/

func NumIslands(grid [][]byte) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				bfsNumIslands(grid, i, j)
				result += 1
			}
		}
	}
	return result
}

func bfsNumIslands(grid [][]byte, x, y int) {
	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	queue := [][]int{
		{x, y},
	}

	grid[x][y] = '0'

	for len(queue) != 0 {
		currentSize := len(queue)
		for i := 0; i < currentSize; i++ {
			for _, direction := range directions {
				newX := queue[i][0] + direction[0]
				newY := queue[i][1] + direction[1]

				if newX >= 0 &&
					newX < len(grid) &&
					newY >= 0 &&
					newY < len(grid[0]) &&
					grid[newX][newY] == '1' {
					grid[newX][newY] = '0'
					queue = append(queue, []int{newX, newY})

				}
			}
		}

		queue = queue[currentSize:]
	}

}

/*
118. Pascal's Triangle
https://leetcode.com/problems/pascals-triangle/
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 		[1]
 	   [1,1]
	  [1,2,1]
	 [1,3,3,1]
	[1,4,6,4,1]
*/

func GeneratePascalTriangle(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result[i] = []int{1}
			continue
		}

		if i == 1 {
			result[i] = []int{1, 1}
			continue
		}

		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 0; j < len(result[i-1])-1; j++ {
			temp[j+1] = result[i-1][j] + result[i-1][j+1]
		}
		result[i] = temp
	}
	return result
}

/*
159. Longest substring which contains 2 unique characters
Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.

Example 1:

Input: "eceba"
Output: 3
Explanation: "ece" which its length is 3.
Example 2:

Input: "ccaabbb"
Output: 5
Explanation: "aabbb" which its length is 5.

*/

func LengthOfLongestSubstringTwoDistinct(s string) int {
	result := 0
	l, r, currentChar := 0, 0, 0
	indexMap := make(map[byte]int)
	// ccaabbb
	for r < len(s) {
		if _, ok := indexMap[s[r]]; ok {
			indexMap[s[r]] = r
		} else {
			if currentChar == 2 {
				temp := s[l]
				l = indexMap[s[l]] + 1
				delete(indexMap, temp)
				currentChar -= 1
			}
			indexMap[s[r]] = r
			currentChar += 1
		}

		if r-l+1 > result {
			result = r - l + 1
		}
		r += 1
	}

	return result
}

/*
214. Shortest Palindrome
https://leetcode.com/problems/shortest-palindrome/description/
aacecaaa
aaacecaa
aaacecaaa
*/

func ShortestPalindrome(s string) string {
	reverse := func(s string) string {
		result := make([]byte, len(s))
		l := 0
		r := len(s) - 1
		for r > l {
			result[l], result[r] = s[r], s[l]
			l += 1
			r -= 1
			if r == l {
				result[l] = s[l]
			}
		}
		return string(result)
	}

	reverseStr := reverse(s)
	// brute force
	//for i := len(s); i >= 0; i-- {
	//	if s[:i] == reverseStr[len(s)-i:] {
	//		return reverseStr[:len(s)-i] + s
	//	}
	//}

	// kmp
	newStr := s + "#" + reverseStr
	lps := make([]int, len(newStr))
	prefix := 0
	i := 1

	for i < len(newStr) {
		if newStr[i] == newStr[prefix] {
			prefix += 1
			lps[i] = prefix
			i += 1
		} else if prefix > 0 {
			prefix = lps[prefix-1]
		} else {
			i += 1
		}
	}

	return reverseStr[:len(reverseStr)-lps[len(lps)-1]] + s
}

/*
210. Course Schedule II
https://leetcode.com/problems/course-schedule-ii/description/
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

*/

func FindCourseScheduleOrder(numCourses int, prerequisites [][]int) []int {
	// i means course, val is how many requires courses
	requires := make([]int, numCourses)
	// i means course, val is which course need i as prerrquires
	adjs := make([][]int, numCourses)

	for _, pre := range prerequisites {
		requires[pre[0]] += 1
		adjs[pre[1]] = append(adjs[pre[1]], pre[0])
	}

	queue := make([]int, 0)

	for i, v := range requires {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, numCourses)
	index := 0
	for len(queue) != 0 {
		currentSize := len(queue)

		for i := 0; i < currentSize; i++ {
			result[index] = queue[i]
			index += 1
			for _, course := range adjs[queue[i]] {
				requires[course] -= 1
				if requires[course] == 0 {
					queue = append(queue, course)
				}
			}
		}

		queue = queue[currentSize:]

	}

	if index != numCourses {
		return []int{}
	}

	return result
}
