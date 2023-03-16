package _0230317

import (
	__Daily_Prac "leetcode/0_Daily_Prac"
	"math"
	"sort"
)

type Name struct {
}

func MergeSort(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	a := nums[:n/2]
	b := nums[n/2:]

	a = MergeSort(a)
	b = MergeSort(b)

	return merge(a, b)
}

func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	l, m, n := 0, 0, 0

	for m < len(a) && n < len(b) {
		if a[m] > b[n] {
			c[l] = b[n]
			n++
		} else {
			c[l] = a[m]
			m++
		}
		l++
	}

	for l < len(c) {
		if m < len(a) {
			c[l] = a[m]
			m++
		} else {
			c[l] = b[n]
			n++
		}
		l++
	}

	return c
}

func HeapSort(nums []int) []int {
	n := len(nums)
	buildHeap(nums)
	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		nums = nums[:i]
		maxHeap(nums, 0)
	}
	return nums[:n]
}

func buildHeap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		maxHeap(nums, i)
	}
}

func maxHeap(nums []int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var largest int

	if l < len(nums) && nums[l] > nums[n] {
		largest = l
	} else {
		largest = n
	}

	if r < len(nums) && nums[r] > nums[largest] {
		largest = r
	}

	if largest != n {
		nums[n], nums[largest] = nums[largest], nums[n]
		maxHeap(nums, largest)
	}
}

func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > x {
				nums[j+1] = nums[j]
				if j == 0 {
					nums[0] = x
				}
			} else {
				nums[j+1] = x
				break
			}
		}
	}
	return nums
}

func QuickSort(nums []int, p, r int) []int {
	if r > p {
		q := partition(nums, p, r)
		QuickSort(nums, p, q-1)
		QuickSort(nums, q+1, r)
	}
	return nums
}

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j < r; j++ {
		if nums[j] < x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchInRotatedArray(nums []int, target int) int {
	n := len(nums) - 1
	if nums[0] < nums[n] {
		return BinarySearch(nums, target)
	}
	l, r, pivot := 0, n, 0

	for r > l {
		mid := int(uint(l+r) >> 1)
		if nums[mid] > nums[mid+1] {
			pivot = mid + 1
			break
		}

		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
	}

	if target >= nums[pivot] && target <= nums[n] {
		result := BinarySearch(nums[pivot:], target)
		if result == -1 {
			return -1
		}
		return result + pivot
	} else {
		return BinarySearch(nums[:pivot], target)
	}
}

func MaxProduct(nums []int) int {
	minVal, maxVal, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if v < 0 {
			minVal, maxVal = maxVal, minVal
		}

		minVal = __Daily_Prac.Min(minVal*v, v)
		maxVal = __Daily_Prac.Max(maxVal*v, v)

		if maxVal > res {
			res = maxVal
		}
	}
	return res
}

func CoinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if c > i {
				break
			}

			if dp[i-c] != math.MaxInt {
				dp[i] = __Daily_Prac.Min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] != math.MaxInt {
		return dp[amount]
	}
	return -1
}

func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = __Daily_Prac.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func LongestCharatersInReplacement(s string, k int) int {
	l, r, mostFreq, res := 0, 0, 0, 0
	counter := make([]int, 26)

	for r < len(s) {
		counter[s[r]-'A']++
		mostFreq = __Daily_Prac.Max(mostFreq, counter[s[r]-'A'])

		if (r-l+1)-mostFreq > k {
			counter[s[l]-'A']--
			l++
		}

		if (r - l + 1) > res {
			res = r - l + 1
		}
		r++
	}
	return res
}

func MinWindow(s, t string) string {
	l, r, start, distance, lengthOfT := 0, 0, 0, math.MaxInt, len(t)
	counter := make(map[byte]int)
	for i := 0; i < lengthOfT; i++ {
		counter[t[i]]++
	}

	for r < len(s) {
		if counter[s[r]] > 0 {
			lengthOfT--
		}
		counter[s[r]]--

		for lengthOfT == 0 {
			if (r - l + 1) < distance {
				distance = r - l + 1
				start = l
			}

			if counter[s[l]] == 0 {
				lengthOfT++
			}

			counter[s[l]]++
			l++
		}
		r++
	}

	if distance != math.MaxInt {
		return s[start : start+distance]
	}
	return ""
}

func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1

	for nums[i-1] >= nums[i] {
		i--
		if i == 0 {
			__Daily_Prac.Reverse(nums, 0)
			return
		}
	}

	j := len(nums) - 1

	for j > i && nums[j] <= nums[i-1] {
		j--
	}

	nums[j], nums[i-1] = nums[i-1], nums[j]

	__Daily_Prac.Reverse(nums, i)
}

func Combination(nums []int, target int) [][]int {
	res := make([][]int, 0)
	currentRes := make([]int, 0)
	combination(nums, target, 0, 0, &res, &currentRes)
	return res
}

func combination(nums []int, target, currentIndex, currentSum int, res *[][]int, currentRes *[]int) {
	if currentSum > target {
		return
	}

	if currentSum == target {
		temp := make([]int, len(*currentRes))
		copy(temp[:], (*currentRes)[:])
		*res = append(*res, temp)
	}

	for i := currentIndex; i < len(nums); i++ {
		currentSum += nums[i]
		*currentRes = append(*currentRes, nums[i])
		combination(nums, target, i, currentSum, res, currentRes)
		currentSum -= nums[i]
		*currentRes = (*currentRes)[:len(*currentRes)-1]
	}
}
