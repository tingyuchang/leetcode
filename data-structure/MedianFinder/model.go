package MedianFinder

type MedianFinder struct {
	nums []float64
}

func Constructor() MedianFinder {
	return MedianFinder{make([]float64, 0)}
}

func (this *MedianFinder) AddNum(num int) {
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

func (this *MedianFinder) FindMedian() float64 {
	mid := int(uint(0+len(this.nums)-1) >> 1)

	if len(this.nums)%2 == 0 {
		return (this.nums[mid] + this.nums[mid+1]) / 2
	}

	return this.nums[mid]
}
