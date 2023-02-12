package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Strings"
	"testing"
)

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
