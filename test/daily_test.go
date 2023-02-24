package test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	_0230224 "leetcode/0_Daily_Prac/20230224"
	"reflect"
	"regexp"
	"testing"
)

var name = _0230224.Name{}
var mergeSort = _0230224.MergeSort
var heapSort = _0230224.HeapSort
var insertionSort = _0230224.InsertionSort
var quickSort = _0230224.QuickSort
var binarySearch = _0230224.BinarySearch
var binarySearchRotated = _0230224.BinarySearchInRotatedArray
var maxProduct = _0230224.MaxProduct
var longestChar = _0230224.LongestCharatersInReplacement
var nextPermutation = _0230224.NextPermutation

func TestDaily(t *testing.T) {
	re := regexp.MustCompile(`\d{8}`)
	fmt.Printf("%q\n", re.Find([]byte(reflect.TypeOf(name).PkgPath())))

	mergeSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	mergeSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	mergeSortAns := mergeSort(mergeSortInput)
	fmt.Printf("MergeSort")
	assert.Equal(t, mergeSortAns, mergeSortExpected)

	heapSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heapSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	heapSortAns := heapSort(heapSortInput)
	fmt.Printf("\nHeapSort")
	assert.Equal(t, heapSortAns, heapSortExpected)

	insertionSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	insertionSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	insertionSortAns := insertionSort(insertionSortInput)
	fmt.Printf("\nInsertionSort")
	assert.Equal(t, insertionSortAns, insertionSortExpected)

	quickSortInput := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	quickSortExpected := []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}
	quickSortAns := quickSort(quickSortInput, 0, len(quickSortInput)-1)
	fmt.Printf("\nQuickSort")
	assert.Equal(t, quickSortAns, quickSortExpected)

	binarySearchInput := []int{3, 5, 7, 9, 11}
	binarySearchTarget1 := 7
	binarySearchTarget2 := 11
	binarySearchTarget3 := 6
	binarySearchExpected1 := 2
	binarySearchExpected2 := 4
	binarySearchExpected3 := -1

	binarySearchAns1 := binarySearch(binarySearchInput, binarySearchTarget1)
	binarySearchAns2 := binarySearch(binarySearchInput, binarySearchTarget2)
	binarySearchAns3 := binarySearch(binarySearchInput, binarySearchTarget3)
	fmt.Printf("\nBinarySearch")
	assert.Equal(t, binarySearchAns1, binarySearchExpected1)
	assert.Equal(t, binarySearchAns2, binarySearchExpected2)
	assert.Equal(t, binarySearchAns3, binarySearchExpected3)

	binarySearchRotatedTestData := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 9, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 2, -1},
		{[]int{3, 1}, 3, 0},
	}
	fmt.Printf("\nBinarySearch Rotated Array")
	for _, td := range binarySearchRotatedTestData {
		result := binarySearchRotated(td.input, td.target)
		assert.Equal(t, result, td.expected)
	}

	maxProductTestData := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2}, -2},
		{[]int{-3, -1, -1}, 3},
		{[]int{-4, -3}, 12},
		{[]int{0, 2}, 2},
		{[]int{2, -5, -2, -4, 3}, 24},
		{[]int{7, -2, -4}, 56},
		{[]int{-1, -2, -9, -6}, 108},
		{[]int{1, 0, -1, 2, 3, -5, -2}, 60},
		{[]int{0, -2, -3}, 6},
		{[]int{-2, 3, 1, 3}, 9},
		{[]int{0, -2, 0}, 0},
		{[]int{3, -2, -3, 3, -1, 0, 1}, 54},
	}
	fmt.Printf("\nMax Prodcut")
	for _, td := range maxProductTestData {
		result := maxProduct(td.input)
		assert.Equal(t, result, td.expected)
	}

	LongestCharatersTestData := []struct {
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
	fmt.Printf("\nLongest Charaters")
	for _, td := range LongestCharatersTestData {
		result := longestChar(td.input, td.replacement)
		assert.Equal(t, result, td.expected)
	}

	nextPermutationTestData := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 6, 5}},
		{[]int{1, 2, 3, 4, 6, 5}, []int{1, 2, 3, 5, 4, 6}},
		{[]int{1, 2, 3, 5, 4, 6}, []int{1, 2, 3, 5, 6, 4}},
	}

	fmt.Printf("\nnext Permutation")
	for _, td := range nextPermutationTestData {
		nextPermutation(td.input)
		// in place exchange
		assert.Equal(t, td.input, td.expected)
	}

}
