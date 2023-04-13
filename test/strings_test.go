package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Strings"
	"testing"
)

func TestBullsCowsGetHint(t *testing.T) {
	testData := []struct {
		secret   string
		guess    string
		expected string
	}{
		//{
		//	"1807", "7810", "1A3B",
		//},
		//{
		//	"1123", "0111", "1A1B",
		//},
		{
			"1122", "1222", "3A0B",
		},
	}

	for _, td := range testData {
		result := Strings.BullsCowsGetHint(td.secret, td.guess)
		assert.Equal(t, result, td.expected)
	}
}

func TestRemoveStars(t *testing.T) {
	testData := []struct {
		s        string
		expected string
	}{
		{
			"leet**cod*e", "lecoe",
		},
		{"erase*****", ""},
	}

	for _, td := range testData {
		result := Strings.RemoveStars(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestCheckInclusion(t *testing.T) {
	testData := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"ab",
			"eidbaooo",
			true},
		{"ab",
			"eidboaoo",
			false},
		{"adc",
			"dcda",
			true},
	}

	for _, td := range testData {
		result := Strings.CheckInclusion(td.s1, td.s2)
		assert.Equal(t, result, td.expected)
	}
}

func TestFirstUniqChar(t *testing.T) {
	testData := []struct {
		s        string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aaaa", -1},
	}

	for _, td := range testData {
		result := Strings.FirstUniqChar(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	testData := []struct {
		s        string
		expected int
	}{
		{"abccccdd", 7},
		{"a", 1},
	}

	for _, td := range testData {
		result := Strings.LongestPalindrome(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestReverseWords(t *testing.T) {
	testData := []struct {
		s        string
		expected string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}

	for _, td := range testData {
		result := Strings.ReverseWords(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestIsSubSequence(t *testing.T) {
	testData := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"foo", "bar", false},
		{"ab", "baab", true},
		{"aaaaaa", "bbaaaa", false},
	}

	for _, td := range testData {
		result := Strings.IsSubsequence(td.s, td.t)
		assert.Equal(t, result, td.expected)
	}
}

func TestNumDecoding(t *testing.T) {
	testData := []struct {
		s        string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}

	for _, td := range testData {
		result := Strings.NumDecodings(td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestRepeatedSubstringPattern(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{
			"abab",
			true,
		},
		{"aba", false},
		{"abcabcabcabc", true},
		{"abaabaa", false},
		{"abaababaab", true},
	}

	for _, td := range testData {
		result := Strings.RepeatedSubstringPattern(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestIndexOfSubString(t *testing.T) {
	testData := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			"sadbutsad", "sad", 0,
		},
		{
			"leetcode", "leeto", -1,
		},
		{"a", "a", 0},
		{"abc", "c", 2},
		{"aaa", "aaaa", -1},
		{"mississippi", "issip", 4},
	}

	for _, td := range testData {
		result := Strings.IndexOfSubString(td.haystack, td.needle)
		assert.Equal(t, result, td.expected)
	}
}

func TestReverseVowels(t *testing.T) {
	testData := []struct {
		input    string
		expected string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{" ", " "},
		{" le ol ", " lo el "},
		{".", "."},
		{"aA", "Aa"},
	}

	for _, td := range testData {
		result := Strings.ReverseVowels(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestReverseString(t *testing.T) {
	testData := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, td := range testData {
		Strings.ReverseString(td.input)
		assert.Equal(t, td.input, td.expected)
	}
}

func TestCompress(t *testing.T) {
	testData := []struct {
		input    []byte
		expected int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
		{[]byte{'a', 'a', 'a', 'a', 'b', 'a'}, 4},
		{[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'a', 'b', 'c'}, 12},
	}
	for _, td := range testData {
		result := Strings.Compress(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestMinWindow(t *testing.T) {
	testData := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, td := range testData {
		result := Strings.MinWindow(td.s, td.t)
		assert.Equal(t, result, td.expected)
	}
}

func TestWordPattern(t *testing.T) {
	testData := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{
			"abba",
			"dog cat cat dog",
			true,
		},
		{
			"abba",
			"dog cat cat fish",
			false,
		},
		{
			"aaaa",
			"dog cat cat dog",
			false,
		},
		{
			"abba",
			"dog dog dog dog",
			false,
		},
		{
			"aaa",
			"aa aa aa aa",
			false,
		},
	}

	for _, td := range testData {
		result := Strings.WordPattern(td.pattern, td.s)
		assert.Equal(t, result, td.expected)
	}
}

func TestEditDistance(t *testing.T) {
	testData := []struct {
		word1    string
		word2    string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}

	for _, td := range testData {
		result := Strings.MinDistance(td.word1, td.word2)
		assert.Equal(t, result, td.expected)
	}
}

func TestCountSubStrings(t *testing.T) {
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
		result := Strings.CountSubstrings(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestPalindromePartition(t *testing.T) {
	testData := []struct {
		input    string
		expected [][]string
	}{
		{"aab", [][]string{
			{"a", "a", "b"},
			{"aa", "b"},
		}},
		{"a", [][]string{
			{"a"},
		}},
		{"aaab", [][]string{
			{"a", "a", "a", "b"},
			{"a", "aa", "b"},
			{"aa", "a", "b"},
			{"aaa", "b"},
		}},
	}

	for _, td := range testData {
		result := Strings.PalindromePartition(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestCharacterReplacement(t *testing.T) {
	testData := []struct {
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

	for _, td := range testData {
		result := Strings.CharacterReplacement(td.input, td.replacement)
		assert.Equal(t, result, td.expected)
	}
}

func TestIsAnagram(t *testing.T) {
	testData := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"car", "rat", false},
	}

	for _, td := range testData {
		result := Strings.IsAnagram(td.input1, td.input2)
		assert.Equal(t, result, td.expected)
	}
}

func TestIsIsomorphic(t *testing.T) {
	testData := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"egg", "add", true},
		{"paper", "title", true},
		{"foo", "bar", false},
	}

	for _, td := range testData {
		result := Strings.IsIsomorphic(td.input1, td.input2)
		assert.Equal(t, result, td.expected)
	}
}

func TestConvertToTitle(t *testing.T) {
	testData := []struct {
		input    int
		expected string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, td := range testData {
		result := Strings.ConvertToTitle(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestTitleToNumber(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, td := range testData {
		result := Strings.TitleToNumber(td.input)
		assert.Equal(t, result, td.expected)
	}
}
