package _0230630

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

	for i := 1; i <= amount; i++ {
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
	memo := make([]int, len(s)+1)
	return NumberOfArraysTopDown(s, 0, k, memo)
}

func NumberOfArraysTopDown(s string, start, k int, memo []int) int {
	if start == len(s) {
		return 1
	}

	if memo[start] != 0 {
		return memo[start]
	}

	if s[start] == '0' {
		return 0
	}

	ans := 0

	for end := start; end < len(s); end++ {
		x, _ := strconv.Atoi(s[start : end+1])

		if x > k {
			break
		}

		ans = (ans + NumberOfArraysTopDown(s, end+1, k, memo)) % __Daily_Prac.MOD
	}
	memo[start] = ans

	return ans
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
	dp[0] = true
	wordsMap := make(map[string]struct{})

	for _, v := range wordDict {
		wordsMap[v] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordsMap[s[j:i]]; ok && dp[j] {
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
	queue := make([]string, 0)
	queue = append(queue, beginWord)
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
						if currentWord[j] != word[j] {
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
		if len(subArray) == 0 || nums[i] > subArray[len(subArray)-1] {
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
	n := len(height)
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

sum is greater than or equal to target
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
	currentIndex := 0
	return constructTree(preorder, 0, len(preorder)-1, &currentIndex, inorderMap)
}

func constructTree(preorder []int, start, end int, currentIndex *int, inorderMap map[int]int) *Tree.TreeNode {
	if *currentIndex >= len(preorder) {
		return nil
	}

	rootVal := preorder[*currentIndex]
	*currentIndex += 1
	node := &Tree.TreeNode{Val: rootVal}
	pivot := inorderMap[rootVal]

	if pivot-start > 0 {
		node.Left = constructTree(preorder, start, pivot-1, currentIndex, inorderMap)
	}

	if end-pivot > 0 {
		node.Right = constructTree(preorder, pivot+1, end, currentIndex, inorderMap)
	}
	return node

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
	return RandomSelect(nums, 0, len(nums)-1, len(nums)-k+1)
}

func RandomSelect(nums []int, p, r, target int) int {
	if p == r {
		return nums[p]
	}

	q := partition(nums, p, r)
	k := q - p + 1

	if k == target {
		return nums[q]
	} else if k > target {
		return RandomSelect(nums, p, q-1, target)
	} else {
		return RandomSelect(nums, q+1, r, target-k)
	}
}

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1

	for j := p; j < r; j++ {
		if nums[j] < x {
			i += 1
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	i += 1
	nums[r], nums[i] = nums[i], nums[r]
	return i
}

// 22. Generate Parentheses

func GenerateParenthesis(n int) []string {
	result := []string{}
	recursiveGenerateParenthesis("", n, 0, 0, &result)
	return result
}

func recursiveGenerateParenthesis(s string, n, leftCount, rightCount int, result *[]string) {
	if len(s) == n*2 {
		*result = append(*result, s)
		return
	}

	if leftCount < n {
		recursiveGenerateParenthesis(s+"(", n, leftCount+1, rightCount, result)
	}

	if rightCount < leftCount {
		recursiveGenerateParenthesis(s+")", n, leftCount, rightCount+1, result)
	}

}

// 56. Merge Intervals

func MergeIntervals(intervals [][]int) [][]int {
	endMaps := make(map[int]int)
	starts := make([]int, 0)

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]

		if existing, ok := endMaps[start]; ok {
			if existing < end {
				endMaps[start] = end
			}
		} else {
			endMaps[start] = end
			starts = append(starts, start)
		}
	}

	sort.Ints(starts)
	current := starts[0]
	result := make([][]int, 0)

	for i := 1; i < len(starts); i++ {
		next := starts[i]
		if endMaps[current] >= next && endMaps[current] < endMaps[next] {
			endMaps[current] = endMaps[next]
		} else if endMaps[current] < next {
			result = append(result, []int{current, endMaps[current]})
			current = next
		}
	}

	result = append(result, []int{current, endMaps[current]})

	return result
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

			if sum < 0 {
				l += 1
			} else if sum > 0 {
				r -= 1
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l += 1

				for i < r && nums[l] == nums[l-1] {
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
			if l < r {
				if len(ans) == 0 {
					ans = string(s[l+1 : r+1])
				} else {
					ans += " " + string(s[l+1:r+1])
				}
			}
			r = l - 1
		}

		l -= 1
	}

	if l < r {
		ans += " " + string(s[l+1:r+1])
	}

	return ans
}

// 91. Decode Ways

func NumDecodings(s string) int {
	results := make([]int, len(s)+1)
	results[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			results[i] += results[i+1]

			if i < len(s)-1 && (s[i] == '1' || s[i] == '2' && s[i+1] < '7') {
				results[i] += results[i+2]
			}
		}
	}

	return results[0]
}

// 394. Decode String

func DecodeString(s string) string {
	stack := make([]byte, 0)
	openBracketStack := make([]int, 0)
	numbersStack := make([]int, 0)
	currentNumber := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			x, _ := strconv.Atoi(string(s[i]))
			currentNumber = currentNumber*10 + x
		} else if s[i] == '[' {
			numbersStack = append(numbersStack, currentNumber)
			currentNumber = 0
			openBracketStack = append(openBracketStack, len(stack))
		} else if s[i] == ']' {
			duplicatedNumber := numbersStack[len(numbersStack)-1]
			numbersStack = numbersStack[:len(numbersStack)-1]
			position := openBracketStack[len(openBracketStack)-1]
			openBracketStack = openBracketStack[:len(openBracketStack)-1]
			duplicatedStr := string(stack[position:])
			stack = stack[:position]

			for j := 0; j < duplicatedNumber; j++ {
				stack = append(stack, duplicatedStr...)
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
	return isMatchTopDown(text, pattern, 0, 0, memo)
}

func isMatchTopDown(text, pattern string, i, j int, memo [][]*Result) bool {
	if memo[i][j] != nil {
		return memo[i][j].val
	}

	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		firstMatch := i < len(text) && (text[i] == pattern[j] || pattern[j] == '.')

		if j < len(pattern)-1 && pattern[j+1] == '*' {
			ans = firstMatch && isMatchTopDown(text, pattern, i+1, j, memo) || isMatchTopDown(text, pattern, i, j+2, memo)
		} else {
			ans = firstMatch && isMatchTopDown(text, pattern, i+1, j+1, memo)
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

type WordCount struct {
	word  string
	index int
	count int
}

func WordCountEngine(document string) [][]string {
	words := strings.Split(document, " ")

	wordFreq := make(map[string]*WordCount)

	for i, word := range words {
		lowerCase := strings.ToLower(word)
		if len(lowerCase) == 0 || lowerCase == " " {
			continue
		}
		newWord := make([]byte, 0)

		for i := 0; i < len(lowerCase); i++ {
			if lowerCase[i] >= 'a' && lowerCase[i] <= 'z' {
				newWord = append(newWord, lowerCase[i])
			}
		}

		if existing, ok := wordFreq[string(newWord)]; ok {
			existing.count += 1
		} else {
			wordFreq[string(newWord)] = &WordCount{
				word:  string(newWord),
				index: i,
				count: 1,
			}
		}
	}

	temp := make([]*WordCount, 0)

	for _, v := range wordFreq {
		temp = append(temp, v)
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].count == temp[j].count {
			return temp[i].index < temp[j].index
		}
		return temp[i].count > temp[j].count
	})

	result := make([][]string, len(temp))

	for i := 0; i < len(temp); i++ {
		result[i] = []string{temp[i].word, strconv.Itoa(temp[i].count)}
	}

	return result
}

/*

 */

func MinimizeDistanceToFarthestPoint(blocks [][]bool, requires int) int {
	result := make([][]int, len(blocks))
	for i := range result {
		result[i] = make([]int, requires+1)
		for j := 0; j < len(result[i])-1; j++ {
			result[i][j] = math.MaxInt
		}
	}

	for i := 0; i < len(blocks); i++ {

		for j := 0; j < requires; j++ {
			if blocks[i][j] {
				result[i][j] = 0
			} else if i > 0 && result[i-1][j] != math.MaxInt {
				result[i][j] = result[i-1][j] + 1
			}

			if result[i][j] != math.MaxInt && result[i][j] > result[i][requires] {
				result[i][requires] = result[i][j]
			}
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		for j := 0; j < requires; j++ {
			if blocks[i][j] {
				result[i][j] = 0
			} else if i < len(blocks)-1 && result[i+1][j] != math.MaxInt {
				result[i][j] = __Daily_Prac.Min(result[i+1][j]+1, result[i][j])
			}

			if result[i][j] != math.MaxInt && result[i][j] > result[i][requires] {
				result[i][requires] = result[i][j]
			}
		}
	}

	smallestDistance := result[0][requires]
	ans := 0

	for i := 1; i < len(result); i++ {
		if result[i][requires] < smallestDistance {
			smallestDistance = result[i][requires]
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
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				for k := 1; k <= 9; k++ {
					x := strconv.Itoa(k)[0]
					if validateSudoku(board, i, j, x) {
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

func validateSudoku(board [][]byte, x, y int, target byte) bool {

	for i := 0; i < len(board); i++ {
		if board[x][i] == target {
			return false
		}

		if board[i][y] == target {
			return false
		}

		if board[x/3+i/3][y/3+i%3] == target {
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
	r := n - 1
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

	if nums[n-1] >= target {
		result := binarySearch(nums[pivot:], target)
		if result == -1 {
			return -1
		} else {
			return result + pivot
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

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

/*
73. Set Matrix Zeroes
https://leetcode.com/problems/set-matrix-zeroes/
*/
func SetZeroes(matrix [][]int) {
	needSetZero := false

	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			needSetZero = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if needSetZero {
		for i := 0; i < m; i++ {
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
	dfsCombinationSum(candidates, 0, target, 0, &[]int{}, &result)
	return result
}

func dfsCombinationSum(candidates []int, start, target, currentSum int, currentResult *[]int, result *[][]int) {
	if currentSum > target {
		return
	}

	if currentSum == target {
		temp := make([]int, len(*currentResult))
		copy(temp, *currentResult)
		*result = append(*result, temp)
	}

	for i := start; i < len(candidates); i++ {
		currentSum += candidates[i]
		*currentResult = append(*currentResult, candidates[i])
		dfsCombinationSum(candidates, i, target, currentSum, currentResult, result)
		currentSum -= candidates[i]
		*currentResult = (*currentResult)[:len(*currentResult)-1]

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

		for node != nil && node.Left != nil {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
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
	rows := make([]int, 0)
	cols := make([]int, 0)

	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cols = append(cols, j)
				rows = append(rows, i)
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
	r := len(matrix)

	for r >= l {
		mid := (l + r) / 2

		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	slot := r

	if slot < 0 {
		slot = 0
	}

	l = 0
	r = len(matrix[slot])
	for r >= l {
		mid := (l + r) / 2

		if matrix[slot][mid] == target {
			return true
		} else if matrix[slot][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
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

	nums[j], nums[i-1] = nums[i-1], nums[j]
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
	if k == len(nums) {
		return nums
	}
	freqMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]] += 1
	}
	myHeap := make([]int, 0)

	for val, count := range freqMap {
		if len(myHeap) == k {
			if freqMap[myHeap[0]] < count {
				myHeap[0] = val
				minHeapTopDown(myHeap, 0, freqMap)
			}
		} else {
			myHeap = append([]int{val}, myHeap...)
			//minHeapBottomUp(myHeap, len(myHeap)-1, freqMap)
			minHeapTopDown(myHeap, 0, freqMap)
		}

	}

	return myHeap
}
func minHeapBottomUp(nums []int, n int, freqMap map[int]int) {
	current := n
	parent := (n - 1) / 2

	for current != 0 && freqMap[nums[parent]] > freqMap[nums[current]] {
		nums[parent], nums[current] = nums[current], nums[parent]
		current = parent
		parent = (current - 1) / 2
	}

}
func minHeapTopDown(nums []int, n int, freqMap map[int]int) {
	l := ((n + 1) * 2) - 1
	r := (n + 1) * 2

	var smallest int
	if l < len(nums) && freqMap[nums[l]] < freqMap[nums[n]] {
		smallest = l
	} else {
		smallest = n
	}

	if r < len(nums) && freqMap[nums[r]] < freqMap[nums[smallest]] {
		smallest = r
	}
	if smallest != n {
		nums[smallest], nums[n] = nums[n], nums[smallest]
		minHeapTopDown(nums, smallest, freqMap)
	}
}

/*
402. Remove K Digits
https://leetcode.com/problems/remove-k-digits/

*/

func RemoveKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}

	if k == 0 {
		return num
	}

	stack := make([]byte, 0)

	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k -= 1
		}
		if len(stack) > 0 || num[i] != '0' {
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

	current := intervals[0]
	result := make([][]int, 0)

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if current[1] < next[0] {
			result = append(result, []int{current[1], next[0]})
			current = next
		} else {
			if current[1] < next[1] {
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
	current := 0
	ans := nums[0]

	for i := 0; i < len(nums); i++ {
		current = __Daily_Prac.Max(current, 0) + nums[i]
		ans = __Daily_Prac.Max(ans, current)
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
	suffix := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		suffix += nums[i]
		rightMax[i] = __Daily_Prac.Max(rightMax[i+1], suffix)
	}

	specialSum := 0
	current := 0
	prefix := 0
	normalSum := nums[0]

	for i := 0; i < len(nums); i++ {
		current = __Daily_Prac.Max(current, 0) + nums[i]
		normalSum = __Daily_Prac.Max(current, normalSum)
		prefix += nums[i]
		if i < len(nums)-1 {
			specialSum = __Daily_Prac.Max(specialSum, prefix+rightMax[i+1])
		}

	}
	return __Daily_Prac.Max(specialSum, normalSum)
}

/*
978. Longest Turbulent Subarray
https://leetcode.com/problems/longest-turbulent-subarray/description/
*/
func MaxTurbulenceSize(arr []int) int {
	compare := func(a, b int) int {
		v := a - b
		if v > 0 {
			return -1
		} else if v < 0 {
			return 1
		} else {
			return 0
		}
	}

	ans := 1
	anchor := 0

	for i := 1; i < len(arr); i++ {
		c := compare(arr[i], arr[i-1])

		if c == 0 {
			anchor = i
		} else if i == len(arr)-1 || c*compare(arr[i], arr[i+1]) != 1 {
			ans = __Daily_Prac.Max(ans, i-anchor+1)
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
	preLPS, i := 0, 1

	for i < len(needle) {
		if needle[i] == needle[preLPS] {
			preLPS += 1
			lps[i] = preLPS
			i += 1
		} else if preLPS == 0 {
			i += 1
		} else {
			preLPS -= 1
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
