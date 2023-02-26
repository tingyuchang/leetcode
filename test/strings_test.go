package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Strings"
	"testing"
)

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
