package _0230625

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
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = __Daily_Prac.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
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
	dp := make([]int, len(s))
	dpPre := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = 1

		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[j] = dpPre[j-1] + 2
			} else {
				dp[j] = __Daily_Prac.Max(dp[j-1], dpPre[j])
			}
		}

		copy(dpPre, dp)
	}

	return dp[len(s)-1]
}

// 322. Coin Change
/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if c > i {
				break
			}

			if dp[i-c] != math.MaxInt {
				dp[i] = __Daily_Prac.Min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] != math.MaxInt {
		return dp[amount]
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
	// memo[i] : s[i+1:] has how many results,
	memo := make([]int, len(s)+1)

	return dfsNumberOfArrays(s, 0, k, memo)
}

func dfsNumberOfArrays(s string, start, k int, memo []int) int {
	if start == len(s) {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	if memo[start] != 0 {
		return memo[start]
	}

	ans := 0
	for end := start; end < len(s); end++ {
		x, _ := strconv.Atoi(s[start : end+1])
		if x > k {
			break
		}

		ans = (ans + dfsNumberOfArrays(s, end+1, k, memo)) % __Daily_Prac.MOD
	}

	memo[start] = ans

	return memo[start]
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
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)

		if i == 0 {
			for j := 0; j < len(dp[i]); j++ {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// 139. Word Break
/*
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.
*/

func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	// dp[i] = s[:i] is valid or not
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
						if endWord == word {
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
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = __Daily_Prac.Min(__Daily_Prac.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
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
		if len(subArray) == 0 || subArray[len(subArray)-1] < nums[i] {
			subArray = append(subArray, nums[i])
		} else {
			l := 0
			r := len(subArray) - 1

			for r >= l {
				mid := (l + r) / 2

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
	totalCost := 0
	currentTank := 0
	station := 0

	for i := 0; i < len(gas); i++ {
		totalCost += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]

		if currentTank < 0 {
			currentTank = 0
			station = i + 1
		}
	}

	if totalCost >= 0 {
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

	for i := len(ratings) - 2; i >= 0; i-- {
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
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(rightMax)-1] = height[len(height)-1]
	n := len(height) - 1

	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}

		if height[n-i] > rightMax[n-i+1] {
			rightMax[n-i] = height[n-i]
		} else {
			rightMax[n-i] = rightMax[n-i+1]
		}

	}

	water := 0

	for i := 1; i < n; i++ {
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
	l, r, sum := 0, 0, 0
	ans := math.MaxInt

	for r < len(nums) {
		sum += nums[r]

		for sum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}

			sum -= nums[l]
			l += 1

		}
		r += 1
	}

	if ans != math.MaxInt {
		return ans
	}
	return 0
}

// 105. Construct Binary Tree from Preorder and Inorder Traversal
/*
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *Tree.TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	index := 0
	return constructTree(preorder, 0, len(preorder)-1, &index, inorderMap)
}

func constructTree(preorder []int, start, end int, index *int, inorderMap map[int]int) *Tree.TreeNode {
	if *index == len(preorder) {
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
		head = nil
		next = nil
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
			mid := (l + r) / 2

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
	fast := head
	move := n - 1
	for move > 0 {
		move -= 1
		fast = fast.Next
	}
	preHead := &LinkedList.ListNode{}
	preHead.Next = head
	pre := preHead

	for fast.Next != nil {
		fast = fast.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return preHead.Next
}

// 92. Reverse Linked List II

func ReverseBetween(head *LinkedList.ListNode, left int, right int) *LinkedList.ListNode {
	first := head

	preFirst := &LinkedList.ListNode{}
	preFirst.Next = head
	move := left - 1

	for move > 0 {
		first = first.Next
		preFirst = preFirst.Next
		move -= 1
	}

	var reverseStart, next *LinkedList.ListNode
	reverseEnd := first
	move = right - left + 1

	for move > 0 && first != nil {
		next = first.Next
		first.Next = reverseStart
		reverseStart = first
		first = next
		move -= 1
	}

	reverseEnd.Next = next
	preFirst.Next = reverseStart

	if left == 1 {
		return reverseStart
	}
	return head
}

// 215. Kth Largest Element in an Array

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
	} else if k > target {
		return randomSelect(nums, p, q-1, target)
	} else {
		return randomSelect(nums, q+1, r, target-k)
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
	ans := make([]string, 0)
	recursiveGenerateParenthesis("", 0, 0, n, &ans)

	return ans
}

func recursiveGenerateParenthesis(s string, openCount, closeCount, n int, ans *[]string) {
	if len(s) == n*2 {
		*ans = append(*ans, s)
		return
	}

	if openCount < n {
		recursiveGenerateParenthesis(s+"(", openCount+1, closeCount, n, ans)
	}

	if closeCount < openCount {
		recursiveGenerateParenthesis(s+")", openCount, closeCount+1, n, ans)
	}
}

// 56. Merge Intervals

func MergeIntervals(intervals [][]int) [][]int {
	endsMap := make(map[int]int)
	starts := make([]int, 0)

	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]

		if existing, ok := endsMap[start]; ok {
			if existing < end {
				endsMap[start] = end
			}
		} else {
			endsMap[start] = end
			starts = append(starts, start)
		}
	}

	sort.Ints(starts)
	ans := make([][]int, 0)

	currentStart := starts[0]

	for i := 1; i < len(starts); i++ {
		nextStart := starts[i]
		if nextStart <= endsMap[currentStart] && endsMap[currentStart] < endsMap[nextStart] {
			endsMap[currentStart] = endsMap[nextStart]
		} else if nextStart > endsMap[currentStart] {
			ans = append(ans, []int{currentStart, endsMap[currentStart]})
			currentStart = nextStart
		}
	}
	ans = append(ans, []int{currentStart, endsMap[currentStart]})
	return ans
}

// 15. 3Sum
func ThreeSum(nums []int) [][]int {
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
				l += 1

				for l < r && nums[l] == nums[l-1] {
					l += 1
				}
			}
		}
	}

	return ans
}

// 151. Reverse Words in a String

func ReverseWords(s string) string {
	ans := ""
	l, r := len(s)-1, len(s)-1

	for l >= 0 {
		if s[l] == ' ' {
			if r > l {
				if len(ans) == 0 {
					ans = s[l+1 : r+1]
				} else {
					ans += " " + s[l+1:r+1]
				}
			}
			r = l - 1
		}
		l -= 1
	}
	if r > l {
		ans += " " + s[l+1:r+1]
	}

	return ans
}

// 91. Decode Ways

func NumDecodings(s string) int {
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

// 394. Decode String

func DecodeString(s string) string {
	stack := make([]byte, 0)
	openBracketStack := make([]int, 0)
	numStack := make([]int, 0)

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
			openBracket := openBracketStack[len(openBracketStack)-1]
			openBracketStack = openBracketStack[:len(openBracketStack)-1]

			currentStr := string(stack[openBracket:])
			stack = stack[:openBracket]
			for j := 0; j < duplicatedCount; j++ {
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
	ans := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		currentSize := len(ans)

		for j := 0; j < currentSize; j++ {
			temp := make([]int, len(ans[j]))
			copy(temp, ans[j])
			temp = append(temp, nums[i])
			ans = append(ans, temp)
		}
	}

	return ans
}

// 133. Clone Graph

func CloneGraph(node *LinkedList.NodeWithNeighbors) *LinkedList.NodeWithNeighbors {
	return recursiveClone(node, make(map[int]*LinkedList.NodeWithNeighbors))
}

func recursiveClone(node *LinkedList.NodeWithNeighbors, nodesMap map[int]*LinkedList.NodeWithNeighbors) *LinkedList.NodeWithNeighbors {
	if nodesMap[node.Val] != nil {
		return nodesMap[node.Val]
	}

	copyNode := &LinkedList.NodeWithNeighbors{
		Val:       node.Val,
		Neighbors: make([]*LinkedList.NodeWithNeighbors, 0),
	}

	for _, neighbor := range node.Neighbors {
		if nodesMap[neighbor.Val] != nil {
			copyNode.Neighbors = append(copyNode.Neighbors, nodesMap[neighbor.Val])
		} else {
			copyNode.Neighbors = append(copyNode.Neighbors, recursiveClone(neighbor, nodesMap))
		}
	}

	return copyNode
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
	return IsMatchTopDown(text, pattern, 0, 0, memo)
}

func IsMatchTopDown(text, pattern string, i, j int, memo [][]*Result) bool {
	if memo[i][j] != nil {
		return memo[i][j].val
	}

	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		firstMatch := i < len(text) && (text[i] == pattern[j] || pattern[j] == '.')

		if j < len(pattern)-1 && pattern[j+1] == '*' {
			ans = IsMatchTopDown(text, pattern, i, j+2, memo) || (firstMatch && IsMatchTopDown(text, pattern, i+1, j, memo))
		} else {
			ans = firstMatch && IsMatchTopDown(text, pattern, i+1, j+1, memo)
		}
	}
	memo[i][j] = &Result{ans}
	return ans
}

// word count engine
/*
input:  document = "Practice makes perfect. you'll only
                    get Perfect by practice. just practice!"

output: [ ["practice", "3"], ["perfect", "2"],
          ["makes", "1"], ["youll", "1"], ["only", "1"],
          ["get", "1"], ["by", "1"], ["just", "1"] ]

*/
type Wordcount struct {
	word  string
	count int
	index int
}

func WordCountEngine(document string) [][]string {
	wordsMap := make(map[string]*Wordcount)
	l, r := 0, 0

	for r < len(document) {
		if document[r] == ' ' {
			if r > l {
				insertWord(document, l, r, &wordsMap)
			}
			l = r + 1
		}

		r += 1
	}

	if r > l {
		insertWord(document, l, r, &wordsMap)
	}

	temp := make([]*Wordcount, 0)
	for _, v := range wordsMap {
		temp = append(temp, v)
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].count == temp[j].count {
			return temp[i].index < temp[j].index
		}
		return temp[i].count > temp[j].count
	})
	ans := make([][]string, len(temp))

	for i, v := range temp {
		ans[i] = []string{
			v.word,
			strconv.Itoa(v.count),
		}
	}

	return ans
}

func insertWord(document string, l, r int, wordsMap *map[string]*Wordcount) {
	word := strings.ToLower(document[l:r])
	onlyAlphabet := make([]byte, 0)

	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			onlyAlphabet = append(onlyAlphabet, word[i])
		}
	}

	if existing, ok := (*wordsMap)[string(onlyAlphabet)]; ok {
		existing.count += 1
	} else {
		(*wordsMap)[string(onlyAlphabet)] = &Wordcount{word: string(onlyAlphabet), count: 1, index: l}
	}
}

func MinimizeDistanceToFarthestPoint(blocks [][]bool, requires int) int {
	distances := make([][]int, len(blocks))

	for i := range distances {
		distances[i] = make([]int, requires)
		for j := 0; j < len(distances[i]); j++ {
			distances[i][j] = math.MaxInt
		}
	}

	/*
		[1, 0, 4]
		[0, 1, 3]
		[0, 0, 2]
		[1, 0, 1]
		[2, 0, 0]

		BFS
	*/

	for i := 0; i < len(distances); i++ {
		for j := 0; j < len(distances[i]); j++ {
			if blocks[i][j] {
				distances[i][j] = 0
			} else {
				if i > 0 && distances[i-1][j] != math.MaxInt {
					distances[i][j] = distances[i-1][j] + 1
				}
			}
		}
	}

	for i := len(distances) - 1; i >= 0; i-- {
		for j := 0; j < len(distances[i]); j++ {
			if blocks[i][j] {
				distances[i][j] = 0
			} else {
				if i < len(distances)-1 && distances[i+1][j] != math.MaxInt {
					distances[i][j] = __Daily_Prac.Min(distances[i+1][j]+1, distances[i][j])
				}
			}
		}
	}
	minDistance := -1
	ans := -1
	for i := 0; i < len(distances); i++ {
		temp := distances[i][0]
		for j := 1; j < len(distances[i]); j++ {
			if distances[i][j] > temp {
				temp = distances[i][j]
			}
		}

		if minDistance == -1 || temp < minDistance {
			minDistance = temp
			ans = i
		}
	}

	return ans
}
