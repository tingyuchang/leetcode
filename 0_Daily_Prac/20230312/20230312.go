package _0230312

type Name struct {
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
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > key {
				nums[j+1] = nums[j]
				if j == 0 {
					nums[0] = key
				}
			} else {
				nums[j+1] = key
				break
			}
		}
	}
	return nums
}

func QuickSort(nums []int, p, r int) []int {
	return nil
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
