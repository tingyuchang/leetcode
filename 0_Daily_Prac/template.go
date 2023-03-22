package __Daily_Prac

type Name struct {
}

type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

func Constructor() MedianFinder {
	min := make([]int, 0)
	max := make([]int, 0)
	return MedianFinder{min, max}
}

func (mf *MedianFinder) AddNum(n int) {

}
func (mf *MedianFinder) FindMedian() int {
	return -1
}

func (mf *MedianFinder) minHeapUp() {

}

func (mf *MedianFinder) minHeapDown() {

}

func (mf *MedianFinder) maxHeapUp() {

}

func (mf *MedianFinder) maxHeapDown() {

}

func MergeSort(nums []int) []int {

	return nil
}

func merge(a, b []int) []int {

	return nil
}

func HeapSort(nums []int) []int {

	return nil
}

func buildHeap(nums []int) {

}

func maxHeap(nums []int, n int) {

}

func InsertionSort(nums []int) []int {

	return nums
}

func QuickSort(nums []int, p, r int) []int {

	return nums
}

func partition(nums []int, p, r int) int {

	return 0
}

func BinarySearch(nums []int, target int) int {

	return -1
}

func BinarySearchInRotatedArray(nums []int, target int) int {

	return -1
}

func MaxProduct(nums []int) int {

	return -1
}

func CoinChange(coins []int, amount int) int {

	return -1
}

func LongestCommonSubsequence(text1 string, text2 string) int {

	return 0
}

func LongestCharatersInReplacement(s string, k int) int {

	return -1
}

func MinWindow(s, t string) string {

	return ""
}

func NextPermutation(nums []int) {

}

func Combination(nums []int, target int) [][]int {

	return nil
}

func combination(nums []int, target, currentIndex, currentSum int, res *[][]int, currentRes *[]int) {

}
