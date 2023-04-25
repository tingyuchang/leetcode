package SmallestInfiniteSet

type SmallestInfiniteSet struct {
	addBack       []int
	smallestIndex int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{smallestIndex: 1, addBack: make([]int, 0)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var ans int
	// addback will alway smaller than smallest index
	if len(this.addBack) != 0 {
		ans = this.addBack[0]
		this.addBack = this.addBack[1:]
	} else {
		ans = this.smallestIndex
		this.smallestIndex++
	}

	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.smallestIndex {
		if len(this.addBack) == 0 {
			this.addBack = append(this.addBack, num)
			return
		}
		// check exist in addBack or not
		// binary search to find element or insert index
		position := binarySearch(this.addBack, num)

		if position != -1 {
			temp := make([]int, position)
			copy(temp, this.addBack[:position])
			temp = append(temp, num)
			temp = append(temp, this.addBack[position:]...)
			this.addBack = temp
		}
	}
}

// return insert index
// -1 means not exist
func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if nums[mid] == target {
			return -1
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return l
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
