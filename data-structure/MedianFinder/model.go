package MedianFinder

import "fmt"

type MedianFinder struct {
	maxHeap []int
	minHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	lengthOfMin := len(this.minHeap)
	lengthOfMax := len(this.maxHeap)
	if lengthOfMin == 0 {
		this.minHeap = append(this.minHeap, num)
		return
	}

	if lengthOfMin < lengthOfMax {
		// add to min
		if num < this.maxHeap[0] {
			this.minHeap = append(this.minHeap, this.maxHeap[0])
			this.minHeapUp()
			this.maxHeap[0] = num
			this.maxHeapDown()
		} else {
			this.minHeap = append(this.minHeap, num)
			this.minHeapUp()
		}

	} else {
		// add to max
		if num > this.minHeap[0] {
			this.maxHeap = append(this.maxHeap, this.minHeap[0])
			this.maxHeapUp()
			this.minHeap[0] = num
			this.minHeapDown()
		} else {
			this.maxHeap = append(this.maxHeap, num)
			this.maxHeapUp()
		}
	}
	fmt.Println("min: ", this.minHeap, " max: ", this.maxHeap)
}

// from last element, if last element is bigger than others, make ti up
func (this *MedianFinder) minHeapUp() {
	n := len(this.minHeap) - 1
	m := (n - 1) / 2

	for n > 0 && this.minHeap[m] > this.minHeap[n] {
		this.minHeap[n], this.minHeap[m] = this.minHeap[m], this.minHeap[n]
		n = m
		m = (m - 1) / 2
	}
}

// from index 0, if index 0 is smaller than others, make it down
func (this *MedianFinder) minHeapDown() {
	n := 0
	l := n<<1 + 1
	r := l + 1

	for l < len(this.minHeap) {
		min := l

		if r < len(this.minHeap) && this.minHeap[l] > this.minHeap[r] {
			min = r
		}
		if this.minHeap[min] > this.minHeap[n] {
			min = n
		}

		if min == n {
			break
		}
		this.minHeap[min], this.minHeap[n] = this.minHeap[n], this.minHeap[min]
		n = min
		l = n<<1 + 1
		r = l + 1
	}
}

func (this *MedianFinder) maxHeapUp() {
	n := len(this.maxHeap) - 1
	m := (n - 1) / 2

	for n > 0 && this.maxHeap[m] < this.maxHeap[n] {
		this.maxHeap[n], this.maxHeap[m] = this.maxHeap[m], this.maxHeap[n]
		n = m
		m = (m - 1) / 2
	}
}

func (this *MedianFinder) maxHeapDown() {
	n := 0
	l := n<<1 + 1
	r := l + 1

	for l < len(this.maxHeap) {
		max := l

		if r < len(this.maxHeap) && this.maxHeap[l] < this.maxHeap[r] {
			max = r
		}
		if this.maxHeap[max] < this.maxHeap[n] {
			max = n
		}

		if max == n {
			break
		}
		this.maxHeap[max], this.maxHeap[n] = this.maxHeap[n], this.maxHeap[max]
		n = max
		l = n<<1 + 1
		r = l + 1
	}
}

func (this *MedianFinder) FindMedian() float64 {
	lengthOfMin := len(this.minHeap)
	lengthOfMax := len(this.maxHeap)
	if lengthOfMin == 0 && lengthOfMax == 0 {
		return -1
	}

	if lengthOfMax == lengthOfMin {
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2
	} else if lengthOfMax > lengthOfMin {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}

type MedianFinder2 struct {
	nums []float64
}

func Constructor2() MedianFinder2 {
	return MedianFinder2{make([]float64, 0)}
}

func (this *MedianFinder2) AddNum(num int) {
	l := 0
	r := len(this.nums) - 1
	insert := -1
	for r >= l {
		mid := int(uint(l+r) >> 1)
		if this.nums[mid] > float64(num) {
			r = mid - 1
		} else if this.nums[mid] < float64(num) {
			l = mid + 1
		} else {
			insert = mid
			break
		}
	}

	if insert == -1 {
		insert = l
	}

	this.nums = append(this.nums, 0)

	for i := len(this.nums) - 1; i > insert; i-- {
		this.nums[i] = this.nums[i-1]
	}
	this.nums[insert] = float64(num)
}

func (this *MedianFinder2) FindMedian() float64 {
	mid := int(uint(0+len(this.nums)-1) >> 1)

	if len(this.nums)%2 == 0 {
		return (this.nums[mid] + this.nums[mid+1]) / 2
	}

	return this.nums[mid]
}
