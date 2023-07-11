package __Daily_Prac

// COPY FROM HERE
import (
	"leetcode/LinkedList"
	"leetcode/Tree"
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
	return -1
}

/*
73. Set Matrix Zeroes
https://leetcode.com/problems/set-matrix-zeroes/
*/
func SetZeroes(matrix [][]int) {

}

/*
39. Combination Sum
https://leetcode.com/problems/combination-sum/
*/
func CombinationSum(candidates []int, target int) [][]int {
	return nil
}

/*
116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
*/

func ConnectAllSiblings(root *Tree.NodeN) *Tree.NodeN {
	return nil
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
	return 0
}

/*
74. Search a 2D Matrix
https://leetcode.com/problems/search-a-2d-matrix/
*/
func SearchMatrix(matrix [][]int, target int) bool {
	return false
}

/*
77. Combinations
https://leetcode.com/problems/combinations/
*/
func Combine(n int, k int) [][]int {
	return nil
}

/*
31. Next Permutation
https://leetcode.com/problems/next-permutation/
*/
func NextPermutation(nums []int) {

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
	return nil
}

/*
402. Remove K Digits
https://leetcode.com/problems/remove-k-digits/

*/

func RemoveKdigits(num string, k int) string {
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
	return nil
}

/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/
*/

func MaxSubArray(nums []int) int {
	return 0
}

/*
918. Maximum Sum Circular Subarray
https://leetcode.com/problems/maximum-sum-circular-subarray/description/
*/
func MaxSubarraySumCircular(nums []int) int {
	return 0
}

/*
978. Longest Turbulent Subarray
https://leetcode.com/problems/longest-turbulent-subarray/description/
*/
func MaxTurbulenceSize(arr []int) int {
	return 1
}

/*
28. Find the Index of the First Occurrence in a String
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
please using KMP to resolve this question
*/

func StrStr(haystack string, needle string) int {
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
	return 0
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
	return nil
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
	return 0
}
