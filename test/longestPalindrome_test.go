package test

import (
	"leetcode/longestPalindrome"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	var testData = []struct{
		input 		string
		expected 	string
	} {
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"acbca", "acbca"},
		{"abcda", "a"},
		{"", ""},
		{"babadada", "adada"},
		{"1baccab2", "baccab"},
	}


	for _, tt := range testData {
		actual := longestPalindrome.LongestPalindrome(tt.input)
		assert.Equal(t, actual, tt.expected)
	}

}