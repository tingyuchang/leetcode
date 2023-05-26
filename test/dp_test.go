package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/DP"
	"leetcode/Strings"
	"testing"
)

func TestStoneGameII(t *testing.T) {
	testData := []struct {
		piles []int
		exp   int
	}{
		{[]int{2, 7, 9, 4, 4}, 10},
		{[]int{1, 2, 3, 4, 5, 100}, 104},
	}

	for _, td := range testData {
		result := DP.StoneGameII(td.piles)
		assert.Equal(t, result, td.exp)
	}
}

func TestStoneGame(t *testing.T) {
	testData := []struct {
		piles []int
		exp   bool
	}{
		{[]int{5, 3, 4, 5}, true},
		{[]int{3, 7, 2, 3}, true},
		{[]int{3, 7, 3, 2, 5, 1, 6, 3, 10, 7}, true},
	}

	for _, td := range testData {
		result := DP.StoneGame(td.piles)
		assert.Equal(t, result, td.exp)
	}
}

func TestNew21Game(t *testing.T) {
	i := []struct {
		n      int
		k      int
		maxPts int
		exp    float64
	}{
		{0, 0, 1, 1},
		{10, 1, 10, 1.0},
		{6, 1, 10, 0.6},
		{21, 17, 10, 0.73278},
		{185, 183, 2, 1},
	}
	testData := i

	for _, td := range testData {
		result := DP.New21Game(td.n, td.k, td.maxPts)
		assert.Equal(t, result, td.exp)
	}
}

func TestCountGoodStrings(t *testing.T) {
	testData := []struct {
		low  int
		high int
		zero int
		one  int
		exp  int
	}{
		{3, 3, 1, 1, 8},
		{2, 3, 1, 2, 5},
		{5, 5, 2, 4, 0},
		{200, 200, 10, 1, 764262396},
		{500, 500, 5, 2, 873327137},
		{50000, 100000, 50000, 50000, 6},
	}

	for _, td := range testData {
		result := DP.CountGoodStrings(td.low, td.high, td.zero, td.one)
		assert.Equal(t, result, td.exp)
	}
}

func TestMostPoints(t *testing.T) {
	testData := []struct {
		questions [][]int
		exp       int64
	}{
		{
			[][]int{
				{3, 2},
				{4, 3},
				{4, 4},
				{2, 5},
			}, 5,
		},
		{
			[][]int{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{5, 5},
			}, 7,
		},
	}

	for _, td := range testData {
		result := DP.MostPoints(td.questions)
		assert.Equal(t, result, td.exp)
	}

}

func TestMaxUncrossedLines(t *testing.T) {
	testData := []struct {
		nums1 []int
		nums2 []int
		exp   int
	}{
		//{[]int{1, 4, 2}, []int{1, 2, 4}, 2},
		//{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		//{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
		{[]int{3, 2}, []int{2, 2, 2, 3}, 1},
	}

	for _, td := range testData {
		result := DP.MaxUncrossedLines(td.nums1, td.nums2)
		assert.Equal(t, result, td.exp)
	}

}

func TestNumSubseq(t *testing.T) {
	testData := []struct {
		nums   []int
		target int
		exp    int
	}{
		//{[]int{3, 5, 6, 7}, 9, 4},
		//{[]int{3, 3, 6, 8}, 10, 6},
		//{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
		{
			[]int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26, 9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21, 20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22},
			31,
			688052206,
		},
	}

	for _, td := range testData {
		result := DP.NumSubseq(td.nums, td.target)
		assert.Equal(t, result, td.exp)
	}
}

func TestNumDistinct(t *testing.T) {
	testData := []struct {
		s   string
		t   string
		exp int
	}{
		{"rabbbit", "rabbit", 3},
	}

	for _, td := range testData {
		result := DP.NumDistinct(td.s, td.t)
		assert.Equal(t, result, td.exp)
	}
}
func TestMinimumDeleteSum(t *testing.T) {
	testData := []struct {
		s1  string
		s2  string
		exp int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	}

	for _, td := range testData {
		result := DP.MinimumDeleteSum(td.s1, td.s2)
		assert.Equal(t, result, td.exp)
	}
}

func TestMaximalSquare(t *testing.T) {
	testData := []struct {
		matrix [][]byte
		exp    int
	}{
		{
			[][]byte{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0},
			},
			4,
		},
		{
			[][]byte{
				{0, 1},
				{1, 0},
			},
			1,
		},
		{
			[][]byte{{1}}, 1,
		},
		{[][]byte{{0, 1}}, 1},
	}

	for _, td := range testData {
		result := DP.MaximalSquare(td.matrix)
		assert.Equal(t, result, td.exp)
	}
}

func TestMinFallingPathSum(t *testing.T) {
	testData := []struct {
		matrix [][]int
		exp    int
	}{
		{[][]int{
			{2, 1, 3},
			{6, 5, 4},
			{7, 8, 9},
		}, 13},
		{
			[][]int{
				{-19, 57},
				{-40, -5},
			},
			-59,
		},
	}

	for _, td := range testData {
		result := DP.MinFallingPathSum(td.matrix)
		assert.Equal(t, result, td.exp)
	}
}

func TestUniquePathsWithObstacles(t *testing.T) {
	testData := []struct {
		obstacleGrid [][]int
		exp          int
	}{
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}, 2,
		},
		{
			[][]int{
				{0, 1},
				{0, 0},
			}, 1,
		},
		{
			[][]int{
				{1},
			}, 0,
		},
	}

	for _, td := range testData {
		result := DP.UniquePathsWithObstacles(td.obstacleGrid)
		assert.Equal(t, result, td.exp)
	}
}

func TestDeleteAndEarn(t *testing.T) {
	testData := []struct {
		nums []int
		exp  int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
		{[]int{3, 1}, 4},
		{[]int{8, 7, 3, 8, 1, 4, 10, 10, 10, 2}, 52},
	}

	for _, td := range testData {
		restult := DP.DeleteAndEarn(td.nums)
		assert.Equal(t, restult, td.exp)
	}
}

func TestKanspack(t *testing.T) {
	testData := []struct {
		price     []int
		weight    []int
		maxWeight int
		exp       int
	}{
		{
			[]int{5, 3, 2, 4, 1},
			[]int{1, 2, 5, 4, 3},
			8,
			12,
		},
	}

	for _, td := range testData {
		result := DP.Kanspack(td.price, td.weight, td.maxWeight)
		assert.Equal(t, result, td.exp)
	}
}

func TestProfitableSchemes(t *testing.T) {
	testData := []struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		exp       int
	}{
		{
			5,
			3,
			[]int{2, 2},
			[]int{2, 3},
			2,
		},
		{
			10,
			5,
			[]int{2, 3, 5},
			[]int{6, 7, 8},
			7,
		},
		{
			100,
			10,
			[]int{66, 24, 53, 49, 86, 37, 4, 70, 99, 68, 14, 91, 70, 71, 70, 98, 48, 26, 13, 86, 4, 82, 1, 7, 51, 37, 27, 87, 2, 65, 93, 66, 99, 28, 17, 93, 83, 91, 45, 3, 59, 87, 92, 62, 77, 21, 9, 37, 11, 4, 69, 46, 70, 47, 28, 40, 74, 47, 12, 3, 85, 16, 91, 100, 39, 24, 52, 50, 40, 23, 64, 22, 2, 15, 18, 62, 26, 76, 3, 60, 64, 34, 45, 40, 49, 11, 5, 8, 40, 71, 12, 60, 3, 51, 31, 5, 42, 52, 15, 36},
			[]int{8, 4, 8, 8, 9, 3, 1, 6, 7, 10, 1, 10, 4, 9, 7, 11, 5, 1, 7, 4, 11, 1, 5, 9, 9, 5, 1, 10, 0, 10, 4, 1, 1, 1, 6, 9, 3, 6, 2, 5, 4, 7, 8, 5, 2, 3, 0, 6, 4, 5, 9, 9, 10, 7, 1, 8, 9, 6, 0, 2, 9, 2, 2, 8, 6, 10, 3, 4, 6, 1, 10, 7, 5, 4, 8, 1, 8, 5, 5, 4, 1, 1, 10, 0, 8, 0, 1, 11, 5, 4, 7, 9, 1, 11, 1, 0, 1, 6, 8, 3},
			188883405,
		},
		{
			1,
			1,
			[]int{2, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 2, 2, 2, 1, 2, 1, 1, 2, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 1, 1, 2, 1, 2, 2, 2, 1, 2, 2, 2, 2, 1, 2, 2, 1, 2, 2, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 2, 1, 1, 2, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 2, 2, 1, 1, 2, 2, 2, 2, 2},
			[]int{2, 1, 2, 2, 2, 1, 0, 1, 2, 0, 0, 2, 2, 1, 1, 1, 2, 0, 1, 1, 2, 0, 2, 2, 1, 0, 1, 0, 1, 2, 2, 1, 1, 2, 0, 2, 1, 1, 1, 1, 1, 2, 0, 1, 0, 2, 2, 2, 2, 2, 0, 1, 1, 2, 1, 0, 1, 0, 0, 2, 1, 0, 2, 0, 2, 1, 1, 1, 0, 1, 0, 1, 2, 2, 0, 1, 1, 2, 2, 0, 1, 0, 0, 1, 1, 2, 2, 2, 2, 1, 0, 0, 1, 2, 1, 1, 1, 1, 0, 1},
			31,
		},
	}

	for _, td := range testData {
		result := DP.ProfitableSchemes(td.n, td.minProfit, td.group, td.profit)
		assert.Equal(t, result, td.exp)
	}
}

func BenchmarkCutRodRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DP.MaxCutRodRecursive([]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 10)
	}
}

func BenchmarkCutRodMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DP.MaxCutRodMemo([]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 10)
	}
}

func BenchmarkCutRodButtonUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DP.MaxCutRodButtonUp([]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 10)
	}
}

func TestCutRod(t *testing.T) {
	testData := []struct {
		prices []int
		n      int
		exp    int
	}{
		{
			[]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 4, 10,
		},
		{
			[]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 1, 1,
		},
		{
			[]int{0, 1, 5, 8, 9, 10, 17, 17, 29, 24, 20}, 2, 5,
		},
	}

	for _, td := range testData {
		result := DP.MaxCutRodButtonUp(td.prices, td.n)
		assert.Equal(t, result, td.exp)
	}
}

func TestNumWays(t *testing.T) {
	testData := []struct {
		word   []string
		target string
		exp    int
	}{
		{[]string{"acca", "bbbb", "caca"}, "aba", 6},
		{[]string{"abba", "baab"}, "bab", 4},
	}

	for _, td := range testData {
		result := DP.NumWays(td.word, td.target)
		assert.Equal(t, result, td.exp)
	}
}

func TestMaxValueOfCoins(t *testing.T) {
	testData := []struct {
		piles [][]int
		k     int
		exp   int
	}{
		{[][]int{{1, 100, 3}, {7, 8, 9}}, 2, 101},
		{[][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7, 706},
	}

	for _, td := range testData {
		result := DP.MaxValueOfCoins(td.piles, td.k)
		assert.Equal(t, result, td.exp)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, td := range testData {
		result := Strings.LengthOfLongestSubstring(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestDailyTemperatures(t *testing.T) {
	tesdtData := []struct {
		tempertures []int
		expected    []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{3, 4, 5, 6},
			[]int{1, 1, 1, 0},
		},
		{
			[]int{3, 6, 9},
			[]int{1, 1, 0},
		},
		{
			[]int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78},
			[]int{3, 1, 1, 2, 1, 1, 0, 1, 1, 0},
		},
	}

	for _, td := range tesdtData {
		result := DP.DailyTemperatures(td.tempertures)
		assert.Equal(t, result, td.expected)
	}
}

func TestIsScramble(t *testing.T) {
	testData := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"great", "rgeat", true},
		{"abcde", "caebd", false},
		{"a", "a", true},
		{"acddaaaadbcbdcdaccabdbadccaaa", "adcbacccabbaaddadcdaabddccaaa", false},
	}

	for _, td := range testData {
		result := DP.IsScramble(td.s1, td.s2)
		assert.Equal(t, result, td.expected)
	}
}

func TestMaxSatisfaction(t *testing.T) {
	testData := []struct {
		satisfaction []int
		expected     int
	}{
		{[]int{-1, -8, 0, 5, -9}, 14},
		{[]int{4, 3, 2}, 20},
		{[]int{-1, -4, -5}, 0},
		{[]int{-2, 5, -1., 0, 3, -3}, 35},
	}

	for _, td := range testData {
		result := DP.MaxSatisfaction(td.satisfaction)
		assert.Equal(t, result, td.expected)
	}
}

func TestSubSetWithDup(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 2}, 6},
		{[]int{0}, 2},
	}

	for _, td := range testData {

		result := DP.SubsetsWithDup(td.nums)
		assert.Equal(t, len(result), td.expected)
	}

}

func TestSubSet(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 8},
		{[]int{0}, 2},
	}

	for _, td := range testData {

		result := DP.SubSet(td.nums)
		assert.Equal(t, len(result), td.expected)
	}

}

func TestCombination(t *testing.T) {
	testData := []struct {
		n        int
		k        int
		expected int
	}{
		{4, 2, 6},
		{1, 1, 1},
	}

	for _, td := range testData {

		result := DP.Combination(td.n, td.k)
		assert.Equal(t, len(result), td.expected)
	}

}

func TestMincostTickets(t *testing.T) {
	testData := []struct {
		days     []int
		costs    []int
		expected int
	}{
		{
			[]int{1, 4, 6, 7, 8, 20},
			[]int{2, 7, 15},
			11,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			[]int{2, 7, 15},
			17,
		},
	}

	for _, td := range testData {
		result := DP.MincostTickets(td.days, td.costs)
		assert.Equal(t, result, td.expected)
	}
}

func TestMinPathSum(t *testing.T) {
	testData := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
	}

	for _, td := range testData {
		result := DP.MinPathSum(td.grid)
		assert.Equal(t, result, td.expected)
	}
}

func TestUniquePaths(t *testing.T) {
	testData := []struct {
		m        int
		n        int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}

	for _, td := range testData {
		result := DP.UniquePaths(td.m, td.n)
		assert.Equal(t, result, td.expected)
	}
}

func TestMaxSubarray(t *testing.T) {
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
		result := DP.MaxSubArray(td.nums)
		assert.Equal(t, result, td.expected)
	}
}

func TestRob2(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 1, 1}, 3},
		{[]int{1}, 1},
	}

	for _, td := range testData {
		result := DP.Rob2(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestRob(t *testing.T) {
	testData := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}

	for _, td := range testData {
		result := DP.Rob(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestCombination3(t *testing.T) {
	testData := []struct {
		k        int
		n        int
		expected [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{4, 1, [][]int{}},
	}

	for _, td := range testData {
		result := DP.CombinationSum3(td.k, td.n)
		assert.Equal(t, result, td.expected)
	}
}

func TestCombinationSumII(t *testing.T) {
	testData := []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6}},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, td := range testData {
		result := DP.CombinationSum2(td.candidates, td.target)
		assert.Equal(t, result, td.expected)
	}
}

func TestWordBreakII(t *testing.T) {
	testData := []struct {
		s        string
		wordDict []string
		excepted []string
	}{
		{"catsanddog", []string{"cats", "dog", "sand", "and", "cat"}, []string{"cats and dog", "cat sand dog"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, []string{}},
	}

	for _, td := range testData {
		result := DP.WordBreakII(td.s, td.wordDict)
		assert.Equal(t, result, td.excepted)
	}
}

func TestWordBreak(t *testing.T) {
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
		result := DP.WordBreak(td.s, td.wordDict)
		assert.Equal(t, result, td.excepted)
	}
}

func TestCombinationSum4(t *testing.T) {
	testData := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			[]int{1, 2, 3},
			4,
			7,
		},
		{
			[]int{9},
			3,
			0,
		},
	}

	for _, td := range testData {
		result := DP.CombinationSum4(td.nums, td.target)
		assert.Equal(t, result, td.expected)
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	testData := []struct {
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

	for _, td := range testData {
		result := DP.LongestCommonSubsequence(td.text1, td.text2)
		assert.Equal(t, result, td.expected)
	}
}

func TestCoinChange(t *testing.T) {
	testData := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for _, td := range testData {
		result := DP.CoinChange(td.coins, td.amount)
		assert.Equal(t, result, td.expected)
	}
}

func TestFindMaximizedCapital(t *testing.T) {
	testData := []struct {
		k        int
		w        int
		profits  []int
		capital  []int
		expected int
	}{
		{
			2,
			0,
			[]int{1, 2, 3},
			[]int{0, 1, 1},
			4,
		},
		{
			3,
			0,
			[]int{1, 2, 3},
			[]int{0, 1, 2},
			6,
		},
		{
			1,
			0,
			[]int{1, 2, 3},
			[]int{1, 1, 2},
			0,
		},
		{
			1,
			2,
			[]int{1, 2, 3},
			[]int{1, 1, 2},
			5,
		},
		{
			11,
			11,
			[]int{1, 2, 3},
			[]int{11, 12, 13},
			17,
		},
		{
			111,
			12,
			[]int{319, 178, 37, 756, 152, 763, 481, 1055, 594, 825, 759, 494, 1087, 696, 717, 358, 1091, 1145, 1043, 245, 779, 957, 3, 1060, 1141, 926, 226, 657, 869, 740, 1170, 746, 889, 416, 634, 209, 1189, 1076, 819, 1144, 40, 807, 561, 400, 283, 133, 186, 839, 1043, 491, 936, 559, 70, 9, 854, 1172, 9, 739, 871, 436, 1087, 1029, 66, 561, 798, 1141, 701, 1020, 1072, 92, 636, 509, 392, 77, 84, 202, 843, 563, 329, 906, 101, 997, 1162, 105, 876, 460, 285, 446, 753, 516, 60, 932, 323, 659, 9, 639, 1041, 861, 1199, 343, 792, 1114, 536, 405, 542, 805, 929, 589, 538, 410, 143, 114, 24, 1064, 682, 536, 1016, 1141, 843, 472, 817, 196, 673, 189, 915, 196, 755, 203, 956, 283, 833, 836, 22, 899, 232, 655, 572, 207, 819, 639, 1024, 887, 407, 385, 251, 896, 165, 290, 193, 357, 870, 678, 411, 697, 998, 344, 628, 866, 1004, 894, 547, 574, 36, 847, 691, 601, 896, 363, 21, 1133, 115, 658, 855, 1042, 2, 189, 230, 215, 747, 971, 581, 809, 832, 311, 35, 639, 1127, 701, 167, 672, 763, 997, 1061, 1170, 289, 806, 91, 198, 720, 67, 273, 863, 437, 152, 671, 635, 968, 107, 58, 453, 392, 236, 875, 719, 44, 850, 176, 580, 937, 227, 125, 851, 1188, 915, 747, 1175, 69, 743, 250, 275, 290, 787, 780, 516, 775, 279, 256, 649, 83, 1180, 372, 655, 813, 750, 381, 807, 988, 977, 430, 282, 1067, 551, 726, 756, 372, 17, 1016, 550, 1128, 223, 174, 85, 1039, 1, 27, 178, 822, 243, 440, 79, 772, 211, 640, 454, 1141, 654, 351, 54, 794, 711, 157, 492, 307, 1156, 937, 212, 461, 417, 741, 1104, 642, 1147, 833, 1194, 1066, 1070, 876, 1045, 341, 854, 1014, 787, 518, 404, 1068, 385, 756, 998, 298, 439, 917, 183, 71, 363, 38, 828, 372, 1071, 195, 481, 1153, 871, 712, 803, 849, 145, 1144, 410, 420, 561, 678, 241, 498, 548, 436, 871, 810, 180, 635, 119, 665, 551, 1132, 807, 408, 685, 489, 274, 459, 28, 987, 935, 310, 284, 1163, 924, 812, 267, 1123, 782, 1065, 1075, 1199, 1047, 289, 280, 1044, 931, 1056, 625, 672, 150, 41, 1084, 998, 151, 483, 226, 548, 277, 1187, 1010, 262, 1048, 493, 1084, 845, 1127, 1009, 773, 129, 547, 798, 118, 559, 463, 972, 666, 1155, 519, 381, 405, 819, 201, 790, 729, 1100, 614, 691, 3, 386, 206, 514, 387, 612, 1073, 158, 67, 1116, 164, 496, 1056, 1159, 348, 530, 28, 1177, 774, 1139, 559, 563, 310, 916, 897, 36, 1060, 173, 952, 123, 635, 911, 711, 304, 611, 481, 645, 532, 1033, 385, 402, 797, 117, 307, 969, 1147, 1127, 434, 1099, 1043, 104, 74, 257, 778, 934, 901, 106, 84, 525, 138, 698, 877, 1009, 487, 450, 649, 736, 802, 1032, 456, 757, 10, 2, 720, 1155, 234, 844, 880, 1053, 1134, 821, 1094, 970, 1095, 976, 101, 1163, 104, 38, 198, 350, 896, 345, 867, 1129, 1064, 403, 228, 1103, 416, 561, 722, 915, 1161, 252, 45, 632, 1163, 232, 174, 964, 876, 1126, 479, 432, 984, 815, 544, 1005, 620, 497, 1021, 763, 169, 1028, 706, 1112, 150, 368, 483, 251, 721, 492, 686, 1105, 38, 1140, 1085, 153, 366, 428, 296, 28, 648, 873, 21, 236, 393, 581, 1043, 1086, 747, 402, 201, 1196, 416, 788, 318, 257, 815, 735, 1023, 1143, 1100, 922, 1033, 884, 824, 963, 159, 904, 21, 1168, 511, 723, 400, 239, 338, 89, 1099, 572, 674, 54, 891, 426, 277, 91, 504, 302, 616, 468, 506, 917, 491, 744, 1091, 1051, 594, 465, 850, 338, 417, 320, 1160, 364, 480, 855, 86, 305, 225, 833, 637, 760, 1147, 613, 236, 460, 664, 1183, 38, 806, 100, 654, 852, 975, 816, 506, 1092, 275, 6, 229, 972, 53, 554, 370, 63, 651, 701, 304, 1011, 672, 353, 118, 1111, 448, 549, 151, 1087, 909, 420, 1095, 663, 1119, 1124, 729, 578, 733, 861, 1154, 1195, 387, 1182, 850, 189, 463, 1129, 1185, 541, 546, 1159, 314, 137, 479, 59, 841, 514, 548, 496, 642, 341, 151, 999, 1036, 1186, 704, 550, 1039, 723, 779, 584, 382, 371, 712, 176, 665, 180, 440, 697, 1102, 408, 728, 157, 1050, 558, 692, 336, 384, 107, 839, 477, 204, 862, 345, 429, 1144, 1069, 207, 449, 594, 963, 988, 843, 334, 200, 865, 281, 296, 1040, 914, 995, 743, 1134, 1050, 484, 602, 1001, 570, 1052, 1106, 922, 77, 1099, 681, 360, 955, 1184, 669, 411, 760, 46, 599, 815, 231, 86, 694, 469, 1079, 1161, 1105, 519, 1010, 681, 603, 376, 534, 145, 361, 448, 1006, 91, 397, 671, 597, 238, 119, 467, 587, 723, 162, 804, 638, 568, 195, 460, 610, 868, 806, 995, 1178, 490, 406, 933, 232, 17, 37, 2, 114, 1004, 871, 531, 209, 267, 37, 750, 1196, 940, 89, 725, 377, 898, 6, 685, 210, 1127, 1160, 432, 391, 931, 681, 937, 275, 1190, 137, 743, 652, 430, 566, 180, 1192, 718, 253, 133, 998, 1067, 764, 747, 1159, 827, 143, 506, 641, 327, 468, 731, 50, 15, 569, 80, 310, 1086, 1092, 989, 245, 520, 711, 788, 1144, 938, 1152, 1169, 563, 1053, 1182, 331, 838, 112, 361, 1049, 717, 979, 956, 434, 534, 1083, 117, 280, 1104, 293, 1137, 592, 1019, 606, 526, 216, 924, 197, 137, 1041, 2, 910, 274, 1178, 267, 521, 626, 764, 691, 124, 446, 337, 676, 325, 288, 1120, 924, 512, 777, 1063, 979, 86, 677, 1183, 777, 418, 41, 852, 929, 712, 1132, 1030, 339, 943, 475, 851, 340, 894, 1091, 230, 654, 229, 945, 69, 182, 110, 255, 895, 645, 1061, 793, 1016, 807, 440, 330, 519, 73, 108, 660, 209, 1077, 1191, 938, 415, 1162, 579, 258, 14, 273, 561, 834, 371, 134, 1118, 1139, 1163, 667, 979, 483, 436, 42, 593, 139, 846, 875, 571, 808, 76, 713, 1198, 352, 299, 156, 793, 509, 245, 697, 1106, 236, 287, 236, 644, 683},
			[]int{487, 13, 943, 31, 661, 656, 690, 175, 1147, 760, 96, 290, 755, 504, 1111, 219, 187, 641, 380, 886, 781, 214, 714, 594, 41, 1154, 908, 977, 1183, 28, 464, 524, 895, 1177, 28, 225, 1180, 206, 387, 25, 166, 207, 394, 418, 771, 719, 153, 836, 970, 589, 114, 67, 1180, 1136, 863, 1144, 21, 9, 957, 861, 981, 849, 1031, 361, 541, 649, 220, 718, 263, 390, 24, 830, 103, 955, 233, 1174, 521, 580, 83, 869, 321, 712, 1011, 130, 297, 835, 406, 543, 849, 681, 337, 355, 867, 863, 552, 883, 155, 762, 982, 129, 605, 125, 1111, 442, 89, 1139, 1084, 870, 183, 1142, 78, 12, 893, 677, 817, 348, 1036, 482, 76, 619, 563, 435, 529, 311, 1148, 629, 786, 288, 112, 995, 854, 844, 1002, 948, 570, 736, 1088, 354, 618, 876, 926, 150, 1108, 412, 707, 233, 137, 775, 751, 1137, 481, 349, 150, 518, 1165, 191, 223, 754, 419, 1025, 817, 1001, 278, 692, 403, 1023, 106, 78, 124, 679, 598, 727, 1026, 966, 564, 726, 1148, 643, 806, 1182, 645, 300, 867, 613, 458, 844, 679, 907, 437, 1012, 902, 18, 843, 190, 43, 705, 818, 957, 146, 1175, 130, 744, 941, 975, 692, 1066, 541, 335, 20, 311, 606, 377, 558, 113, 545, 1115, 228, 29, 2, 1180, 331, 1072, 151, 692, 1156, 347, 293, 1135, 959, 865, 1031, 40, 425, 1191, 1178, 87, 989, 92, 1186, 1072, 105, 1058, 369, 401, 1117, 220, 484, 181, 901, 321, 923, 263, 72, 685, 820, 1123, 736, 942, 37, 419, 631, 545, 761, 227, 230, 25, 636, 1048, 834, 736, 899, 530, 217, 669, 278, 653, 657, 857, 724, 782, 146, 780, 615, 1156, 751, 463, 625, 707, 355, 92, 718, 855, 624, 504, 359, 744, 793, 868, 462, 985, 1087, 517, 886, 83, 221, 383, 601, 709, 683, 1097, 544, 411, 28, 1129, 781, 13, 752, 347, 1150, 1030, 269, 701, 658, 658, 1188, 222, 1160, 480, 953, 203, 132, 17, 723, 927, 911, 448, 977, 1126, 219, 118, 1033, 919, 1041, 712, 930, 963, 772, 264, 523, 479, 735, 919, 72, 1019, 131, 15, 628, 331, 408, 578, 37, 1123, 125, 527, 887, 54, 1043, 259, 654, 557, 872, 505, 700, 1077, 202, 368, 628, 29, 66, 199, 611, 730, 1108, 682, 705, 628, 1001, 705, 21, 266, 428, 1093, 800, 742, 701, 715, 845, 1151, 300, 460, 1134, 708, 443, 706, 819, 296, 199, 671, 452, 626, 386, 569, 999, 326, 1081, 202, 384, 783, 541, 811, 1058, 684, 146, 288, 1149, 480, 0, 658, 547, 773, 588, 758, 189, 489, 257, 436, 362, 417, 180, 671, 141, 657, 864, 808, 1026, 504, 827, 1125, 977, 1161, 699, 360, 689, 520, 796, 147, 746, 978, 833, 82, 102, 254, 736, 88, 525, 1037, 659, 1061, 333, 663, 520, 6, 439, 823, 1151, 395, 188, 32, 513, 34, 805, 1017, 1193, 157, 942, 87, 630, 915, 807, 215, 482, 1132, 23, 1008, 1175, 724, 1070, 339, 1139, 22, 455, 306, 369, 151, 1031, 1119, 846, 1195, 49, 1169, 773, 706, 45, 747, 875, 944, 1161, 1025, 258, 765, 1039, 459, 641, 1169, 211, 894, 556, 253, 831, 1115, 458, 351, 1138, 58, 1169, 1190, 743, 611, 392, 1015, 54, 831, 330, 1138, 681, 1012, 750, 967, 1179, 398, 564, 5, 2, 143, 787, 197, 590, 144, 589, 588, 493, 524, 748, 1123, 707, 585, 641, 293, 871, 331, 1173, 280, 218, 931, 11, 663, 1005, 1118, 555, 1000, 699, 720, 81, 527, 71, 458, 462, 1056, 447, 491, 1068, 1078, 842, 455, 497, 959, 745, 654, 1011, 939, 787, 430, 535, 594, 940, 1176, 656, 207, 815, 1133, 610, 1113, 596, 1018, 961, 840, 555, 582, 1187, 828, 161, 983, 686, 870, 800, 525, 34, 1013, 385, 870, 953, 59, 518, 521, 151, 633, 561, 848, 310, 712, 108, 1148, 446, 480, 983, 110, 442, 182, 119, 463, 909, 690, 48, 1040, 631, 17, 1027, 158, 353, 108, 524, 335, 1046, 514, 1027, 746, 917, 378, 437, 606, 829, 743, 462, 1013, 525, 290, 477, 749, 896, 12, 351, 13, 42, 819, 334, 912, 30, 1041, 815, 307, 1138, 168, 209, 1134, 100, 276, 292, 283, 1074, 123, 561, 857, 92, 879, 58, 706, 532, 75, 49, 385, 342, 887, 646, 301, 54, 483, 589, 1084, 1092, 178, 845, 243, 873, 611, 340, 712, 115, 157, 63, 773, 800, 844, 167, 384, 522, 877, 183, 368, 709, 501, 905, 512, 756, 246, 197, 463, 47, 232, 256, 37, 883, 1048, 774, 431, 943, 868, 111, 163, 336, 648, 313, 858, 536, 416, 680, 951, 320, 499, 199, 234, 482, 529, 676, 26, 1180, 721, 877, 586, 628, 1152, 835, 1145, 447, 763, 750, 188, 1169, 596, 1125, 310, 519, 1165, 488, 1063, 59, 292, 701, 1078, 1088, 663, 512, 172, 477, 187, 215, 1192, 22, 827, 279, 607, 286, 179, 744, 569, 500, 510, 1038, 570, 1159, 520, 1070, 759, 831, 906, 644, 753, 574, 257, 983, 1023, 227, 460, 710, 169, 595, 249, 111, 73, 991, 933, 448, 655, 559, 183, 244, 44, 644, 935, 466, 97, 605, 460, 800, 229, 977, 675, 236, 946, 73, 456, 623, 499, 423, 162, 342, 914, 386, 1082, 407, 954, 1081, 1099, 142, 539, 416, 791, 1195, 1099, 885, 965, 971, 796, 1198, 449, 768, 792, 541, 18, 476, 303, 137, 319, 1008, 343, 733, 128, 641, 709, 175, 691, 227, 307, 1177, 1198, 1075, 747, 944, 1038, 933, 643, 613, 1166, 1153, 120, 288, 1167, 246, 1103, 185, 85, 1008, 1060, 1051, 421, 150, 601, 376, 183, 1028, 146, 297, 515, 688, 886, 1090, 552, 969, 903, 1086, 931, 946, 1099, 546, 17, 851, 908, 1170, 418, 94, 61, 233, 1069, 510, 783, 302, 701, 564, 1195, 57, 1007, 994, 205, 1019, 694, 1020, 137, 1041, 803, 673, 1162, 14, 904, 418, 1076, 514, 79, 944, 242, 491, 867, 934, 40, 1059, 776, 817, 468, 516, 550, 906, 790, 459, 273, 924, 185, 1183, 337, 435, 699, 316, 768},
			126981,
		},
	}

	for _, td := range testData {
		result := DP.FindMaximizedCapital(td.k, td.w, td.profits, td.capital)
		assert.Equal(t, result, td.expected)
	}
}
