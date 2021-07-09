package basicProblems

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAnagrams(t *testing.T) {
	testData := []struct{
		a string
		b string
		expected bool
	} {
		{"Whoa! Hi!", "Whoa! Hi!", true},
		{"hello", "llohe", true},
		{"hello", "lloheaa", false},
		{"Hello", "llohe!", true},
		{"One one one", "Two two two", false},
		{"A tree, a life, a bench", "A tree, a fence, a yard", false},

	}

	for _,tt := range testData {
		result := anagrams(tt.a, tt.b)
		assert.Equal(t, result, tt.expected)
	}
}