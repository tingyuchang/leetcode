package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	_0230710 "leetcode/0_Daily_Prac/20230710"
	"leetcode/LinkedList"
	"leetcode/Tree"
	"reflect"
	"regexp"
	"testing"
)

var TreeHelper = Tree.Codec{}

var name = _0230710.Name{}

var LongestCommonSubsequence = _0230710.LongestCommonSubsequence
var LongestPalindrome = _0230710.LongestPalindrome
var LongestPalindromeSubseq = _0230710.LongestPalindromeSubseq
var CoinChange = _0230710.CoinChange
var NumberOfArrays = _0230710.NumberOfArrays
var MinDistance = _0230710.MinDistance
var WordBreak = _0230710.WordBreak
var LadderLength = _0230710.LadderLength
var MaximalSquare = _0230710.MaximalSquare
var LengthOfLIS = _0230710.LengthOfLIS
var MaxProfit = _0230710.MaxProfit
var JumpII = _0230710.JumpII
var CanCompleteCircuit = _0230710.CanCompleteCircuit
var Candy = _0230710.Candy
var Trap = _0230710.Trap
var MinSubArrayLen = _0230710.MinSubArrayLen
var BuildTree = _0230710.BuildTree
var ConnectTreeNode = _0230710.ConnectTreeNode
var NumSubseq = _0230710.NumSubseq
var SwapPairs = _0230710.SwapPairs
var RemoveNthFromEnd = _0230710.RemoveNthFromEnd
var ReverseBetween = _0230710.ReverseBetween
var FindKthLargest = _0230710.FindKthLargest
var GenerateParenthesis = _0230710.GenerateParenthesis
var MergeIntervals = _0230710.MergeIntervals
var ThreeSum = _0230710.ThreeSum
var ReverseWords = _0230710.ReverseWords
var NumDecodings = _0230710.NumDecodings
var DecodeString = _0230710.DecodeString
var Subsets = _0230710.Subsets
var IsMatch = _0230710.IsMatch
var WordCountEngine = _0230710.WordCountEngine
var MinimizeDistanceToFarthestPoint = _0230710.MinimizeDistanceToFarthestPoint
var SolveSudoku = _0230710.SolveSudoku
var SearchRotated = _0230710.SearchRotated
var SetZeroes = _0230710.SetZeroes
var CombinationSum = _0230710.CombinationSum
var BestMeetingPoint = _0230710.BestMeetingPoint
var SearchMatrix = _0230710.SearchMatrix
var Combine = _0230710.Combine
var NextPermutation = _0230710.NextPermutation
var TopKFrequent = _0230710.TopKFrequent
var RemoveKdigits = _0230710.RemoveKdigits
var EmployeeFreeTime = _0230710.EmployeeFreeTime
var MaxSubArray = _0230710.MaxSubArray
var MaxSubarraySumCircular = _0230710.MaxSubarraySumCircular
var MaxTurbulenceSize = _0230710.MaxTurbulenceSize
var StrStr = _0230710.StrStr

func TestDaily(t *testing.T) {
	re := regexp.MustCompile(`\d{8}`)
	fmt.Printf("%q\n", re.Find([]byte(reflect.TypeOf(name).PkgPath())))

	LongestCommonSubsequenceTestData := []struct {
		text1 string
		text2 string
		exp   int
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
		assert.Equal(t, td.exp, result)
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
		assert.Equal(t, td.exp, result)
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
		assert.Equal(t, td.exp, result)
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

	FindKthLargesTestData := []struct {
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
	fmt.Printf("Start test\tFindKthLargest\n")
	for _, td := range FindKthLargesTestData {
		result := FindKthLargest(td.nums, td.k)
		assert.Equal(t, result, td.exp)
	}
	fmt.Printf("End test\tFindKthLargest\n")

	GenerateParenthesisTestData := []struct {
		n   int
		exp []string
	}{
		{3, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
		{1, []string{"()"}},
	}
	fmt.Printf("Start test\tGenerateParenthesis\n")
	for _, td := range GenerateParenthesisTestData {
		result := GenerateParenthesis(td.n)
		assert.ElementsMatch(t, result, td.exp)
	}
	fmt.Printf("End test\tGenerateParenthesis\n")

	MergeIntervalsTestData := []struct {
		intervals [][]int
		exp       [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}
	fmt.Printf("Start test\tMergeIntervals\n")
	for _, td := range MergeIntervalsTestData {
		result := MergeIntervals(td.intervals)
		assert.ElementsMatch(t, result, td.exp)
	}
	fmt.Printf("End test\tMergeIntervals\n")

	ThreeSumTestData := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{-1, 0, 1, 0}, [][]int{[]int{-1, 0, 1}}},
	}
	fmt.Printf("Start test\tThreeSum\n")
	for _, tt := range ThreeSumTestData {
		result := ThreeSum(tt.input)
		assert.ElementsMatch(t, result, tt.expected)
	}
	fmt.Printf("End test\tThreeSum\n")

	ReverseWordsTestData := []struct {
		s   string
		exp string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"", ""},
	}
	fmt.Printf("Start test\tReverseWords\n")
	for _, td := range ReverseWordsTestData {
		result := ReverseWords(td.s)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tReverseWords\n")

	NumDecodingsTestData := []struct {
		s   string
		exp int
	}{
		{"12", 2},
		{"226", 3},
	}
	fmt.Printf("Start test\tNumDecodings\n")
	for _, td := range NumDecodingsTestData {
		result := NumDecodings(td.s)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tNumDecodings\n")

	DecodeStringTestData := []struct {
		s   string
		exp string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		//{"3[a2[c]]", "accaccacc"},
		//{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		//{"3[z]2[2[y]pq4[2[jk]e1[f]]]ef", "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
		//{"100[leetcode]", "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode"},
	}
	fmt.Printf("Start test\tDecodeString\n")
	for _, td := range DecodeStringTestData {
		result := DecodeString(td.s)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tDecodeString\n")

	SubsetsTestData := []struct {
		nums []int
		exp  [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{}, {1},
				{2}, {1, 2},
				{3}, {1, 3},
				{2, 3}, {1, 2, 3},
			},
		},
	}
	fmt.Printf("Start test\tSubsets\n")
	for _, td := range SubsetsTestData {
		result := Subsets(td.nums)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("End test\tSubsets\n")

	IsMatchTestData := []struct {
		text    string
		pattern string
		exp     bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
	}
	fmt.Printf("Start test\tIsMatch\n")
	for _, td := range IsMatchTestData {
		result := IsMatch(td.text, td.pattern)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tIsMatch\n")

	WordCountEngineTestData := []struct {
		document string
		exp      [][]string
	}{
		{
			"Practice makes perfect, you'll get perfecT by practice. just practice! just just just!!",
			[][]string{[]string{"just", "4"}, []string{"practice", "3"}, []string{"perfect", "2"}, []string{"makes", "1"}, []string{"youll", "1"}, []string{"get", "1"}, []string{"by", "1"}},
		},
		{
			"To be, or not to be, that is the question:",
			[][]string{[]string{"to", "2"}, []string{"be", "2"}, []string{"or", "1"}, []string{"not", "1"}, []string{"that", "1"}, []string{"is", "1"}, []string{"the", "1"}, []string{"question", "1"}},
		},
		{
			"Every book is a quotation; and every house is a quotation out of all forests, and mines, and stone quarries; and every man is a quotation from all his ancestors.",
			[][]string{[]string{"and", "4"}, []string{"every", "3"}, []string{"is", "3"}, []string{"a", "3"}, []string{"quotation", "3"}, []string{"all", "2"}, []string{"book", "1"}, []string{"house", "1"}, []string{"out", "1"}, []string{"of", "1"}, []string{"forests", "1"}, []string{"mines", "1"}, []string{"stone", "1"}, []string{"quarries", "1"}, []string{"man", "1"}, []string{"from", "1"}, []string{"his", "1"}, []string{"ancestors", "1"}},
		},
	}

	fmt.Printf("Start test\tWordCountEngine\n")
	for _, td := range WordCountEngineTestData {
		result := WordCountEngine(td.document)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("End test\tWordCountEngine\n")

	SwapPairsTestData := []struct {
		head *LinkedList.ListNode
		exp  *LinkedList.ListNode
	}{
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4}),
			LinkedList.GenerateNodeFromArray([]int{2, 1, 4, 3}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{}),
			LinkedList.GenerateNodeFromArray([]int{}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{1}),
			LinkedList.GenerateNodeFromArray([]int{1}),
		},
	}

	fmt.Printf("Start test\tSwapPairs\n")
	for _, td := range SwapPairsTestData {
		result := SwapPairs(td.head)
		assert.EqualValues(t, result.String(), td.exp.String())
	}
	fmt.Printf("End test\tSwapPairs\n")

	RemoveNthFromEndTestData := []struct {
		head *LinkedList.ListNode
		n    int
		exp  *LinkedList.ListNode
	}{
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}),
			2,
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 5}),
		},
	}

	fmt.Printf("Start test\tRemoveNthFromEnd\n")
	for _, td := range RemoveNthFromEndTestData {
		result := RemoveNthFromEnd(td.head, td.n)
		assert.EqualValues(t, result.String(), td.exp.String())
	}
	fmt.Printf("End test\tRemoveNthFromEnd\n")

	ReverseBetweenTestData := []struct {
		head  *LinkedList.ListNode
		left  int
		right int
		exp   *LinkedList.ListNode
	}{
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}),
			2,
			4,
			LinkedList.GenerateNodeFromArray([]int{1, 4, 3, 2, 5}),
		},
	}

	fmt.Printf("Start test\tReverseBetween\n")
	for _, td := range ReverseBetweenTestData {
		result := ReverseBetween(td.head, td.left, td.right)
		assert.EqualValues(t, result.String(), td.exp.String())
	}
	fmt.Printf("End test\tReverseBetween\n")

	MinimizeDistanceToFarthestPointTestData := []struct {
		blocks   [][]bool
		requires int
		exp      int
	}{
		{
			[][]bool{
				{false, true, false},
				{true, false, false},
				{true, true, false},
				{false, true, false},
				{false, true, true},
			},
			3,
			3,
		},
	}

	fmt.Printf("Start test\tMinimizeDistanceToFarthestPoint\n")
	for _, td := range MinimizeDistanceToFarthestPointTestData {
		result := MinimizeDistanceToFarthestPoint(td.blocks, td.requires)
		assert.EqualValues(t, td.exp, result)
	}
	fmt.Printf("End test\tMinimizeDistanceToFarthestPoint\n")

	SolveSudokutestData := []struct {
		board [][]byte
		exp   [][]byte
	}{
		{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			[][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}

	fmt.Printf("Start test\tSolveSudoku\n")
	for _, td := range SolveSudokutestData {
		SolveSudoku(td.board)
		assert.Equal(t, td.exp, td.board)
	}
	fmt.Printf("End test\tSolveSudoku\n")

	SearchRotatedTestData := []struct {
		input  []int
		target int
		exp    int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 9, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 2, -1},
		{[]int{3, 1}, 3, 0},
	}
	fmt.Printf("Start test\tSearchRotated\n")
	for _, td := range SearchRotatedTestData {
		result := SearchRotated(td.input, td.target)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tSearchRotated\n")

	SetZeroesTestData := []struct {
		matrix [][]int
		exp    [][]int
	}{
		{[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},

		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			[][]int{
				{-4, -2147483648, 6, -7, 0},
				{-8, 6, -8, -6, 0},
				{2147483647, 2, -9, -6, -10},
			},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{2147483647, 2, -9, -6, 0},
			},
		},
	}

	fmt.Printf("Start test\tSetZeroes\n")
	for _, td := range SetZeroesTestData {
		SetZeroes(td.matrix)
		assert.ElementsMatch(t, td.matrix, td.exp)
	}
	fmt.Printf("End test\tSetZeroes\n")

	CombinationSumTestData := []struct {
		candidates []int
		target     int
		exp        [][]int
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
	fmt.Printf("Start test\tCombinationSum\n")
	for _, td := range CombinationSumTestData {
		result := CombinationSum(td.candidates, td.target)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("End test\tCombinationSum\n")

	BestMeetingPointTestData := []struct {
		grid [][]int
		exp  int
	}{
		{
			[][]int{
				{1, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
			},
			6,
		},
	}

	fmt.Printf("Start test\tBestMeetingPoint\n")
	for _, td := range BestMeetingPointTestData {
		result := BestMeetingPoint(td.grid)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tBestMeetingPoint\n")

	SearchMatrixTestData := []struct {
		matrix [][]int
		target int
		exp    bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1},
			},
			0,
			false,
		},
	}

	fmt.Printf("Start test\tSearchMatrix\n")
	for _, td := range SearchMatrixTestData {
		result := SearchMatrix(td.matrix, td.target)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("End test\tSearchMatrix\n")

	CombineTestData := []struct {
		n   int
		k   int
		exp [][]int
	}{
		{
			4, 2,
			[][]int{
				{1, 2}, {1, 3}, {1, 4},
				{2, 3}, {2, 4}, {3, 4},
			},
		},
	}

	fmt.Printf("Start test\tCombine\n")
	for _, td := range CombineTestData {
		result := Combine(td.n, td.k)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("End test\tCombine\n")

	NextPermutationTestData := []struct {
		input []int
		exp   []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 6, 5}},
		{[]int{1, 2, 3, 4, 6, 5}, []int{1, 2, 3, 5, 4, 6}},
		{[]int{1, 2, 3, 5, 4, 6}, []int{1, 2, 3, 5, 6, 4}},
	}

	fmt.Printf("Start test\tNextPermutation\n")
	for _, td := range NextPermutationTestData {
		NextPermutation(td.input)
		assert.Equal(t, td.exp, td.input)
	}
	fmt.Printf("Start test\tNextPermutation\n")

	TopKFrequentTestData := []struct {
		nums []int
		k    int
		exp  []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{3, 0, 1, 0}, 1, []int{0}},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3, []int{1, 2, 3}},
		{[]int{5, 3, 1, 1, 1, 3, 73, 1}, 2, []int{3, 1}},
		{[]int{5, 3, 1, 1, 1, 3, 5, 73, 1}, 3, []int{3, 5, 1}},
		{[]int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
		{[]int{5, 2, 5, 3, 5, 3, 1, 1, 3}, 2, []int{3, 5}},
	}
	fmt.Printf("Start test\tTopKFrequent\n")
	for _, td := range TopKFrequentTestData {
		result := TopKFrequent(td.nums, td.k)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("Start test\tTopKFrequent\n")

	RemoveKdigitsTestData := []struct {
		num string
		k   int
		exp string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
		{"9", 1, "0"},
		{"112", 1, "11"},
		{"10001", 4, "0"},
	}

	fmt.Printf("Start test\tRemoveKdigits\n")
	for _, td := range RemoveKdigitsTestData {
		result := RemoveKdigits(td.num, td.k)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("Start test\tRemoveKdigits\n")

	EmployeeFreeTimeTestData := []struct {
		schedule [][][]int
		exp      [][]int
	}{
		{
			[][][]int{
				{
					{1, 2},
					{5, 6},
				},
				{
					{1, 3},
				},
				{
					{4, 10},
				},
			},
			[][]int{
				{3, 4},
			},
		},
		{
			[][][]int{
				{
					{1, 3},
					{6, 7},
				},
				{
					{2, 4},
				},
				{
					{2, 5},
					{9, 12},
				},
			},
			[][]int{
				{5, 6},
				{7, 9},
			},
		},
	}

	fmt.Printf("Start test\tEmployeeFreeTime\n")
	for _, td := range EmployeeFreeTimeTestData {
		result := EmployeeFreeTime(td.schedule)
		assert.ElementsMatch(t, td.exp, result)
	}
	fmt.Printf("Start test\tEmployeeFreeTime\n")

	MaxSubArrayTestData := []struct {
		nums []int
		exp  int
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
	fmt.Printf("Start test\tMaxSubArray\n")
	for _, td := range MaxSubArrayTestData {
		result := MaxSubArray(td.nums)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("Start test\tMaxSubArray\n")

	MaxSubarraySumCircularTestData := []struct {
		nums []int
		exp  int
	}{
		{
			[]int{1, -2, 3, -2},
			3,
		},
		{
			[]int{5, -3, 5},
			10,
		},
		{
			[]int{1, 2, -3, -4, 5, 6},
			14,
		},
	}
	fmt.Printf("Start test\tMaxSubarraySumCircular\n")
	for _, td := range MaxSubarraySumCircularTestData {
		result := MaxSubarraySumCircular(td.nums)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("Start test\tMaxSubarraySumCircular\n")

	MaxTurbulenceSizeTestData := []struct {
		arr []int
		exp int
	}{
		{
			[]int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			5,
		},
		{
			[]int{4, 8, 12, 16},
			2,
		},
		{
			[]int{100},
			1,
		},
	}
	fmt.Printf("Start test\tMaxTurbulenceSize\n")
	for _, td := range MaxTurbulenceSizeTestData {
		result := MaxTurbulenceSize(td.arr)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("Start test\tMaxTurbulenceSize\n")

	StrStrTestData := []struct {
		haystack string
		needle   string
		exp      int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}
	fmt.Printf("Start test\tStrStr\n")
	for _, td := range StrStrTestData {
		result := StrStr(td.haystack, td.needle)
		assert.Equal(t, td.exp, result)
	}
	fmt.Printf("Start test\tStrStr\n")

}
