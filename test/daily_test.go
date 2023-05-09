package test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	_0230323 "leetcode/0_Daily_Prac/20230323"
	_0230403 "leetcode/0_Daily_Prac/20230403"
	_0230509 "leetcode/0_Daily_Prac/20230509"
	"leetcode/Tree"
	"reflect"
	"regexp"
	"testing"
)

var TreeHelper = Tree.Codec{}

var name = _0230509.Name{}

var LongestCommonSubsequence = _0230509.LongestCommonSubsequence
var LongestPalindrome = _0230509.LongestPalindrome
var LongestPalindromeSubseq = _0230509.LongestPalindromeSubseq
var CoinChange = _0230509.CoinChange
var NumberOfArrays = _0230509.NumberOfArrays
var MinDistance = _0230509.MinDistance
var WordBreak = _0230509.WordBreak
var LadderLength = _0230509.LadderLength
var MaximalSquare = _0230509.MaximalSquare
var LengthOfLIS = _0230509.LengthOfLIS
var MaxProfit = _0230509.MaxProfit
var JumpII = _0230509.JumpII
var CanCompleteCircuit = _0230509.CanCompleteCircuit
var Candy = _0230509.Candy
var Trap = _0230509.Trap
var MinSubArrayLen = _0230509.MinSubArrayLen
var BuildTree = _0230509.BuildTree
var ConnectTreeNode = _0230509.ConnectTreeNode
var NumSubseq = _0230509.NumSubseq

func TestDaily(t *testing.T) {
	re := regexp.MustCompile(`\d{8}`)
	fmt.Printf("%q\n", re.Find([]byte(reflect.TypeOf(name).PkgPath())))

	LongestCommonSubsequenceTestData := []struct {
		text1    string
		text2    string
		expected int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"ezupkr", "ubmrapg", 2},
		{"bsbininm", "jmjkbkjkv", 1},
		{"oxcpqrsvwf", "shmtulqrypy", 2},
	}
	fmt.Printf("Start test\tLongestCommonSubsequenceTestData\n")
	for _, td := range LongestCommonSubsequenceTestData {
		result := LongestCommonSubsequence(td.text1, td.text2)
		assert.Equal(t, result, td.expected)
	}
	fmt.Printf("End test\tLongestCommonSubsequenceTestData\n")

	LongestPalindromeTestData := []struct {
		input    string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"acbca", "acbca"},
		{"abcda", "a"},
		{"", ""},
		{"babadada", "adada"},
		{"1baccab2", "baccab"},
	}
	fmt.Printf("Start test\tLongestPalindrome\n")
	for _, td := range LongestPalindromeTestData {
		actual := LongestPalindrome(td.input)
		assert.Equal(t, actual, td.expected)
	}
	fmt.Printf("End test\tLongestPalindrome\n")

	LongestPalindromeSubseqTestData := []struct {
		s        string
		expected int
	}{
		{"bbbab", 4},
		{"cbbd", 2},
		{"euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159},
	}
	fmt.Printf("Start test\tLongestPalindromeSubseq\n")
	for _, td := range LongestPalindromeSubseqTestData {
		result := LongestPalindromeSubseq(td.s)
		assert.Equal(t, result, td.expected)
	}
	fmt.Printf("End test\tLongestPalindromeSubseq\n")

	CoinChangeTestData := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	fmt.Printf("Start test\tCoinChange\n")
	for _, td := range CoinChangeTestData {
		result := CoinChange(td.coins, td.amount)
		assert.Equal(t, result, td.expected)
	}
	fmt.Printf("End test\tCoinChange\n")

	NumberOfArraysTestData := []struct {
		s   string
		k   int
		exp int
	}{
		{
			"1000", 10000, 1,
		},
		{"1000", 10, 0},
		{"1317", 2000, 8},
		{"1317", 100, 5},
	}
	fmt.Printf("Start test\tNumberOfArrays\n")
	for _, td := range NumberOfArraysTestData {
		result := NumberOfArrays(td.s, td.k)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tNumberOfArrays\n")

	MinDistanceTestData := []struct {
		word1    string
		word2    string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	fmt.Printf("Start test\tMinDistance\n")
	for _, td := range MinDistanceTestData {
		result := MinDistance(td.word1, td.word2)
		assert.Equal(t, result, td.expected)
	}
	fmt.Printf("End test\tMinDistance\n")

	WordBreakTestData := []struct {
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
	fmt.Printf("Start test\tWordBreak\n")
	for _, td := range WordBreakTestData {
		result := WordBreak(td.s, td.wordDict)
		assert.Equal(t, result, td.excepted)
	}
	fmt.Printf("End test\tWordBreak\n")

	LadderLengthTestData := []struct {
		beginWord string
		endWord   string
		wordList  []string
		exp       int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	}

	fmt.Printf("Start test\tLadderLength\n")
	for _, td := range LadderLengthTestData {
		result := LadderLength(td.beginWord, td.endWord, td.wordList)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tLadderLength\n")

	MaximalSquareTestData := []struct {
		matrix [][]byte
		exp    int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			1,
		},
		{
			[][]byte{{'0'}}, 0,
		},
		{[][]byte{{'0', '1'}}, 1},
	}
	fmt.Printf("Start test\tMaximalSquare\n")
	for _, td := range MaximalSquareTestData {
		result := MaximalSquare(td.matrix)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tMaximalSquare\n")

	LengthOfLISTestData := []struct {
		nums []int
		exp  int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	fmt.Printf("Start test\tLengthOfLIS\n")
	for _, td := range LengthOfLISTestData {
		result := LengthOfLIS(td.nums)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tLengthOfLIS\n")

	MaxProfitTestData := []struct {
		prices []int
		exp    int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	// MaxProfit
	fmt.Printf("Start test\tMaxProfit\n")
	for _, td := range MaxProfitTestData {
		result := MaxProfit(td.prices)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tMaxProfit\n")

	JumpIITestData := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{1, 3, 2}, 2},
		{[]int{4, 3, 2, 1}, 1},
		{[]int{2, 1, 1, 1, 1}, 3},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}, 2},
	}
	fmt.Printf("Start test\tJumpII\n")
	for _, td := range JumpIITestData {
		result := JumpII(td.input)
		assert.Equal(t, result, td.expected)
	}
	fmt.Printf("End test\tJumpII\n")

	CanCompleteCircuitTestData := []struct {
		gas  []int
		cost []int
		exp  int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}
	fmt.Printf("Start test\tCanCompleteCircuit\n")
	for _, td := range CanCompleteCircuitTestData {
		result := CanCompleteCircuit(td.gas, td.cost)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tCanCompleteCircuit\n")

	CandyTestData := []struct {
		ratings []int
		exp     int
	}{
		{
			[]int{1, 0, 2}, 5,
		},
		{[]int{1, 2, 2}, 4},
	}

	fmt.Printf("Start test\tCandy\n")
	for _, td := range CandyTestData {
		result := Candy(td.ratings)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tCandy\n")

	TrapTestData := []struct {
		height []int
		exp    int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6,
		},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	fmt.Printf("Start test\tTrap\n")
	for _, td := range TrapTestData {
		result := Trap(td.height)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tTrap\n")

	// MinSubArrayLen

	MinSubArrayLenTestData := []struct {
		target int
		nums   []int
		exp    int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	fmt.Printf("Start test\tMinSubArrayLen\n")
	for _, td := range MinSubArrayLenTestData {
		result := MinSubArrayLen(td.target, td.nums)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tMinSubArrayLen\n")

	BuildTreeTestData := []struct {
		preorder []int
		inorder  []int
		exp      string
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, "3,9,20,null,null,15,7"},
	}

	fmt.Printf("Start test\tBuildTree\n")
	for _, td := range BuildTreeTestData {
		result := BuildTree(td.preorder, td.inorder)
		assert.Equal(t, TreeHelper.Serialize(result), td.exp)
	}
	fmt.Printf("End test\tBuildTree\n")

	NumSubseqTestData := []struct {
		nums   []int
		target int
		exp    int
	}{
		{
			[]int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26, 9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21, 20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22},
			31,
			688052206,
		},
	}
	fmt.Printf("Start test\tNumSubseq\n")
	for _, td := range NumSubseqTestData {
		result := NumSubseq(td.nums, td.target)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tNumSubseq\n")
}

var mergeSort = _0230403.MergeSort
var heapSort = _0230403.HeapSort
var insertionSort = _0230403.InsertionSort
var quickSort = _0230403.QuickSort
var binarySearch = _0230403.BinarySearch
var longestChar = _0230403.LongestCharatersInReplacement
var minWindow = _0230403.MinWindow
var combination = _0230403.Combination
var coinChange = _0230403.CoinChange
var longestCommonSubsequence = _0230403.LongestCommonSubsequence
var medianFinder = _0230323.Constructor()

func testDailyOld(t *testing.T) {
	re := regexp.MustCompile(`\d{8}`)
	fmt.Printf("%q\n", re.Find([]byte(reflect.TypeOf(name).PkgPath())))

	mergeSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	mergeSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	mergeSortAns := mergeSort(mergeSortInput)
	fmt.Printf("MergeSort")
	assert.Equal(t, mergeSortAns, mergeSortExpected)

	heapSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heapSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	heapSortAns := heapSort(heapSortInput)
	fmt.Printf("\nHeapSort")
	assert.Equal(t, heapSortAns, heapSortExpected)

	insertionSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	insertionSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	insertionSortAns := insertionSort(insertionSortInput)
	fmt.Printf("\nInsertionSort")
	assert.Equal(t, insertionSortAns, insertionSortExpected)

	quickSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	quickSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	quickSortAns := quickSort(quickSortInput, 0, len(quickSortInput)-1)
	fmt.Printf("\nQuickSort")
	assert.Equal(t, quickSortAns, quickSortExpected)

	binarySearchInput := []int{3, 5, 7, 9, 11}
	binarySearchTarget1 := 7
	binarySearchTarget2 := 11
	binarySearchTarget3 := 6
	binarySearchExpected1 := 2
	binarySearchExpected2 := 4
	binarySearchExpected3 := -1

	binarySearchAns1 := binarySearch(binarySearchInput, binarySearchTarget1)
	binarySearchAns2 := binarySearch(binarySearchInput, binarySearchTarget2)
	binarySearchAns3 := binarySearch(binarySearchInput, binarySearchTarget3)
	fmt.Printf("\nBinarySearch")
	assert.Equal(t, binarySearchAns1, binarySearchExpected1)
	assert.Equal(t, binarySearchAns2, binarySearchExpected2)
	assert.Equal(t, binarySearchAns3, binarySearchExpected3)

	coinChangeTestData := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	fmt.Printf("\nCoin Change")
	for _, td := range coinChangeTestData {
		result := coinChange(td.coins, td.amount)
		assert.Equal(t, result, td.expected)
	}

	longestCommonSubsequenceTestData := []struct {
		text1    string
		text2    string
		expected int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"ezupkr", "ubmrapg", 2},
		{"bsbininm", "jmjkbkjkv", 1},
		{"oxcpqrsvwf", "shmtulqrypy", 2},
	}

	fmt.Printf("\nlongest Common Subsequence")
	for _, td := range longestCommonSubsequenceTestData {
		result := longestCommonSubsequence(td.text1, td.text2)
		assert.Equal(t, result, td.expected)
	}

	LongestCharatersTestData := []struct {
		input       string
		replacement int
		expected    int
	}{
		{"ABCDDCABAA", 2, 5},
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"KRSCDCSONAJNHLBMDQGIFCPEKPOHQIHLTDIQGEKLRLCQNBOHNDQGHJPNDQPERNFSSSRDEQLFPCCCARFMDLHADJADAGNNSBNCJQOF", 4, 7},
		{"IMNJJTRMJEGMSOLSCCQICIHLQIOGBJAEHQOCRAJQMBIBATGLJDTBNCPIFRDLRIJHRABBJGQAOLIKRLHDRIGERENNMJSDSSMESSTR", 2, 6},
	}
	fmt.Printf("\nLongest Charaters")
	for _, td := range LongestCharatersTestData {
		result := longestChar(td.input, td.replacement)
		assert.Equal(t, result, td.expected)
	}

	minWindowTestData := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}
	fmt.Printf("\nmin window substring")
	for _, td := range minWindowTestData {
		result := minWindow(td.s, td.t)
		assert.Equal(t, result, td.expected)
	}

	fmt.Printf("\ncombination")
	combinationTestData := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}

	for _, td := range combinationTestData {
		result := combination(td.candidates, td.target)
		assert.Equal(t, result, td.expected)
	}

	fmt.Printf("\nMedian Finder")

	medianFinder.AddNum(6)
	assert.Equal(t, medianFinder.FindMedian(), float64(6))
	medianFinder.AddNum(10)
	assert.Equal(t, medianFinder.FindMedian(), float64(8))
	medianFinder.AddNum(2)
	assert.Equal(t, medianFinder.FindMedian(), float64(6))
	medianFinder.AddNum(6)
	assert.Equal(t, medianFinder.FindMedian(), float64(6))
	medianFinder.AddNum(5)
	assert.Equal(t, medianFinder.FindMedian(), float64(6))
	medianFinder.AddNum(0)
	assert.Equal(t, medianFinder.FindMedian(), 5.5)
	medianFinder.AddNum(6)
	assert.Equal(t, medianFinder.FindMedian(), float64(6))
	medianFinder.AddNum(3)
	assert.Equal(t, medianFinder.FindMedian(), 5.5)
	medianFinder.AddNum(1)
	assert.Equal(t, medianFinder.FindMedian(), float64(5))
	medianFinder.AddNum(0)
	assert.Equal(t, medianFinder.FindMedian(), float64(4))
	medianFinder.AddNum(0)
	assert.Equal(t, medianFinder.FindMedian(), float64(3))
}
