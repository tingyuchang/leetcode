package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/data-structure/MedianFinder"
	"testing"
)

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
