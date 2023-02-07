package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/bestTimeToBuyAndSellStock"
	"testing"
)

func TestBestTimeToBuyAndSell(t *testing.T) {
	var testData = []struct {
		input    []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 4, 1}, 2},
		{[]int{4, 7, 2, 1}, 3},
	}

	for _, tt := range testData {
		result := bestTimeToBuyAndSellStock.MaxProfit(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}
