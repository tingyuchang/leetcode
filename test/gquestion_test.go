package test

import (
	"github.com/magiconair/properties/assert"
	__Daily_Prac "leetcode/0_Daily_Prac"
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
		exp int
	}{}

	for _, td := range testData {
		result := GFindClosestElements(td.arr, td.k, td.x)
		assert.Equal(t, result, td.exp)
	}
}
func TestGDeleteNode(t *testing.T) {

}
func TestGCopyRandomList(t *testing.T) {

}
func TestGInvertTree(t *testing.T) {

}
func TestGIsSameTree(t *testing.T) {

}
func TestGWordBreak(t *testing.T) {

}
func TestGCountSubstrings(t *testing.T) {

}
func TestGMaxSubArray(t *testing.T) {

}
func TestGIsNumber(t *testing.T) {

}
func TestGGenerateParenthesis(t *testing.T) {

}
func TestGSearchRange(t *testing.T) {

}
func TestGMergeIntervals(t *testing.T) {

}
func TestGHasPathSum(t *testing.T) {

}
func TestGMissingNumber(t *testing.T) {

}
func TestGReverseList(t *testing.T) {

}
