package Math

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums {
		if cache[v] == 0 {
			res = append(res, v)
		}
		cache[v]++
	}

	buildHeapTopK(res, &cache)
	n := len(res)
	fmt.Println(res)
	for i := n - 1; i > 0; i-- {
		res[i], res[0] = res[0], res[i]
		res = res[:i]
		maxHeapTopK(res, &cache, 0)
	}
	res = res[:n]
	return res[len(res)-k:]
}

func buildHeapTopK(nums []int, cache *map[int]int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		maxHeapTopK(nums, cache, i)
	}
}

func maxHeapTopK(nums []int, cache *map[int]int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var largest int

	if l < len(nums) {
		if (*cache)[nums[l]] > (*cache)[nums[n]] {
			largest = l
		} else if (*cache)[nums[l]] == (*cache)[nums[n]] {
			if nums[l] < nums[n] {
				largest = l
			} else {
				largest = n
			}
		} else {
			largest = n
		}
	} else {
		largest = n
	}

	if r < len(nums) {
		if (*cache)[nums[r]] > (*cache)[nums[largest]] {
			largest = r
		} else if (*cache)[nums[r]] == (*cache)[nums[largest]] {
			if nums[r] < nums[largest] {
				largest = r
			}
		}
	}

	if largest != n {
		nums[n], nums[largest] = nums[largest], nums[n]
		maxHeapTopK(nums, cache, largest)
	}

}

func TopKFrequentNotFinished(nums []int, k int) []int {
	cache := make(map[int]int)
	cacheFreq := make(map[int]int)
	res := make([]int, 0)
	freq := make([]int, 0)
	for _, v := range nums {
		if count, ok := cacheFreq[v]; ok {

			// got current position
			count++
			position := binarySearch(freq, count)
			current := cache[v]

			if position == len(freq) || freq[position] > count {
				position--
			}
			//fmt.Printf("freq: %v, cache: %v value: %d count: %d current: %d new: %d\n", freq, cache, v, count, current, position)

			for i := current; i <= position; i++ {
				if i == position {
					freq[i] = count
					res[i] = v
					break
				}
				freq[i] = freq[i+1]
				res[i] = res[i+1]

			}
			cache[v] = position
			cacheFreq[v] = count

		} else {
			for k, v := range cacheFreq {
				if v > 1 {
					cache[k]++
					cacheFreq[k]++
				}
			}
			cache[v] = 0
			cacheFreq[v] = 1
			res = append([]int{v}, res...)
			freq = append([]int{1}, freq...)
		}
		//fmt.Println(res)
	}

	return res[len(res)-k:]
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			for mid+1 < len(nums) && nums[mid+1] == target {
				mid += 1
			}
			return mid
		}
	}
	return l

}
