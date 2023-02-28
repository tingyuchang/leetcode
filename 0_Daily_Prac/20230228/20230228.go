package _0230228

type Name struct{}

func MergeSort(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	listA := nums[:n/2]
	listB := nums[n/2:]

	listA = MergeSort(listA)
	listB = MergeSort(listB)

	return merge(listA, listB)
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
		nums[0], nums[i] = nums[i], nums[0]
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
		nums[largest], nums[n] = nums[n], nums[largest]
		maxHeap(nums, largest)
	}
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
		mid := (l + r) / 2

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

	l := 0
	r := n
	pivot := 0
	for r > l {
		mid := (r + l) / 2
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
	minVal, maxVal, maxProduct := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if v < 0 {
			maxVal, minVal = minVal, maxVal
		}

		minVal = min(minVal*v, v)
		maxVal = max(maxVal*v, v)

		if maxVal > maxProduct {
			maxProduct = maxVal
		}
	}

	return maxProduct
}

func LongestCharatersInReplacement(s string, k int) int {
	r, l, mostFreq, longest := 0, 0, 0, 0
	count := make([]int, 26)

	for r < len(s) {
		count[s[r]-'A']++
		mostFreq = max(mostFreq, count[s[r]-'A'])
		if (r-l+1)-mostFreq > k {
			count[s[l]-'A']--
			l++
		}

		if (r - l + 1) > longest {
			longest = r - l + 1
		}
		r++
	}

	return longest

}

func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 1

	for nums[i-1] >= nums[i] {
		i--

		if i == 0 {
			reverse(nums, 0)
			return
		}
	}
	j := len(nums) - 1

	for j > i && nums[j] <= nums[i-1] {
		j--
	}

	nums[i-1], nums[j] = nums[j], nums[i-1]
	reverse(nums, i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func reverse(nums []int, n int) {
	l := n
	r := len(nums) - 1

	for r > l {
		nums[l], nums[r] = nums[r], nums[l]
		r--
		l++
	}
}
