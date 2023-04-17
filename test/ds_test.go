package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/data-structure/MedianFinder"
	"leetcode/data-structure/Trie"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testData := []struct {
		strs []string
		exp  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a"}, "a"},
		{[]string{"", "b"}, ""},
		{[]string{"a", "ab"}, "a"},
		{[]string{"flower", "flower", "flower", "flower"}, "flower"},
	}

	for _, td := range testData {
		result := Trie.LongestCommonPrefix(td.strs)
		assert.Equal(t, result, td.exp)
	}
}

func TestMedianFinder(t *testing.T) {
	mf := MedianFinder.Constructor()
	mf.AddNum(6)
	assert.Equal(t, mf.FindMedian(), float64(6))
	mf.AddNum(10)
	assert.Equal(t, mf.FindMedian(), float64(8))
	mf.AddNum(2)
	assert.Equal(t, mf.FindMedian(), float64(6))
	mf.AddNum(6)
	assert.Equal(t, mf.FindMedian(), float64(6))
	mf.AddNum(5)
	assert.Equal(t, mf.FindMedian(), float64(6))
	mf.AddNum(0)
	assert.Equal(t, mf.FindMedian(), 5.5)
	mf.AddNum(6)
	assert.Equal(t, mf.FindMedian(), float64(6))
	mf.AddNum(3)
	assert.Equal(t, mf.FindMedian(), 5.5)
	mf.AddNum(1)
	assert.Equal(t, mf.FindMedian(), float64(5))
	mf.AddNum(0)
	assert.Equal(t, mf.FindMedian(), float64(4))
	mf.AddNum(0)
	assert.Equal(t, mf.FindMedian(), float64(3))

}
