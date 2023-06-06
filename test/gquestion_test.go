package test

import (
	"github.com/stretchr/testify/assert"
	__Daily_Prac "leetcode/0_Daily_Prac"
	"leetcode/LinkedList"
	"leetcode/Math"
	"leetcode/Tree"
	"testing"
)

var GFindKthLargest = __Daily_Prac.GFindKthLargest
var GFindClosestElements = __Daily_Prac.GFindClosestElements
var GDeleteNode = __Daily_Prac.GDeleteNode
var GCopyRandomList = __Daily_Prac.GCopyRandomList
var GInvertTree = __Daily_Prac.GInvertTree
var GIsSameTree = __Daily_Prac.GIsSameTree
var GWordBreak = __Daily_Prac.GWordBreak
var GCountSubstrings = __Daily_Prac.GCountSubstrings
var GMaxSubArray = __Daily_Prac.GMaxSubArray
var GIsNumber = __Daily_Prac.GIsNumber
var GGenerateParenthesis = __Daily_Prac.GGenerateParenthesis

// LRU?
var GSearchRange = __Daily_Prac.GSearchRange
var GMergeIntervals = __Daily_Prac.GMergeIntervals
var GHasPathSum = __Daily_Prac.GHasPathSum
var GMissingNumber = __Daily_Prac.GMissingNumber
var GReverseList = __Daily_Prac.GReverseList

func TestGFindKthLargest(t *testing.T) {
	testData := []struct {
		nums []int
		k    int
		exp  int
	}{
		{
			[]int{3, 2, 1, 5, 6, 4},
			2, 5,
		},
		{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4, 4,
		},
	}

	for _, td := range testData {
		result := GFindKthLargest(td.nums, td.k)
		assert.Equal(t, result, td.exp)
	}
}
func TestGFindClosestElements(t *testing.T) {
	testData := []struct {
		arr []int
		k   int
		x   int
		exp []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{1}, 1, 1, []int{1}},
		{[]int{-2, -1, 1, 2, 3, 4, 5}, 7, 3, []int{-2, -1, 1, 2, 3, 4, 5}},
	}

	for _, td := range testData {
		result := GFindClosestElements(td.arr, td.k, td.x)
		assert.Equal(t, result, td.exp)
	}
}
func TestGDeleteNode(t *testing.T) {
	head := LinkedList.GenerateNodeFromArray([]int{4, 5, 1, 9})
	node := head.Next
	exp := LinkedList.GenerateNodeFromArray([]int{4, 1, 9})
	GDeleteNode(node)
	assert.Equal(t, head.String(), exp.String())
}
func TestGCopyRandomList(t *testing.T) {
	// hard to test
}
func TestGInvertTree(t *testing.T) {
	testData := []struct {
		root *Tree.TreeNode
		exp  *Tree.TreeNode
	}{
		{TreeHelper.Deserialize("4,2,7,1,3,6,9"), TreeHelper.Deserialize("4,7,2,9,6,3,1")},
	}

	for _, td := range testData {
		result := GInvertTree(td.root)
		assert.Equal(t, td.exp, result)
	}
}
func TestGIsSameTree(t *testing.T) {
	testData := []struct {
		p   *Tree.TreeNode
		q   *Tree.TreeNode
		exp bool
	}{
		{TreeHelper.Deserialize("1,2,3"), TreeHelper.Deserialize("1,2,3"), true},
		{TreeHelper.Deserialize("1,2,1"), TreeHelper.Deserialize("1,1,2"), false},
	}

	for _, td := range testData {
		result := GIsSameTree(td.p, td.q)
		assert.Equal(t, td.exp, result)
	}
}
func TestGWordBreak(t *testing.T) {
	testData := []struct {
		s        string
		wordDict []string
		excepted bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"catcats", []string{"cat", "cats"}, true},
		{"catsand", []string{"cat", "cats", "and"}, true},
	}

	for _, td := range testData {
		result := GWordBreak(td.s, td.wordDict)
		assert.Equal(t, result, td.excepted)
	}
}
func TestGCountSubstrings(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{
			"abc",
			3,
		},
		{
			"aaa",
			6,
		},
		{
			"addaccadbabdbdbdbcabdcbcadacccbdddcbddacdaacbbdcbdbccdaaddadcaacdacbaaddbcaadcdab",
			126,
		},
	}

	for _, td := range testData {
		result := GCountSubstrings(td.input)
		assert.Equal(t, result, td.expected)
	}
}
func TestGMaxSubArray(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}

	for _, td := range testData {
		result := GMaxSubArray(td.nums)
		assert.Equal(t, result, td.expected)
	}
}
func TestGIsNumber(t *testing.T) {

}
func TestGGenerateParenthesis(t *testing.T) {
	testData := []struct {
		n   int
		exp []string
	}{
		{3, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
		{1, []string{"()"}},
	}

	for _, td := range testData {
		result := GGenerateParenthesis(td.n)
		assert.ElementsMatch(t, result, td.exp)
	}
}
func TestGSearchRange(t *testing.T) {
	testData := []struct {
		num    []int
		target int
		exp    []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}

	for _, td := range testData {
		result := GSearchRange(td.num, td.target)
		assert.Equal(t, result, td.exp)
	}
}
func TestGMergeIntervals(t *testing.T) {
	testData := []struct {
		intervals [][]int
		exp       [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}

	for _, td := range testData {
		result := GMergeIntervals(td.intervals)
		assert.Equal(t, result, td.exp)
	}
}
func TestGHasPathSum(t *testing.T) {
	testData := []struct {
		root      *Tree.TreeNode
		targetSum int
		exp       bool
	}{
		{TreeHelper.Deserialize("5,4,8,11,null,13,4,7,2,null,null,null,1"), 22, true},
		{TreeHelper.Deserialize("1,2,3"), 5, false},
	}

	for _, td := range testData {
		result := GHasPathSum(td.root, td.targetSum)
		assert.Equal(t, td.exp, result)
	}
}
func TestGMissingNumber(t *testing.T) {
	testData := []struct {
		nums []int
		exp  int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 7, 6, 5, 4, 3, 2, 1, 0}, 8},
	}

	for _, td := range testData {
		result := Math.MissingNumbers(td.nums)
		assert.Equal(t, td.exp, result)
	}
}
func TestGReverseList(t *testing.T) {

}
