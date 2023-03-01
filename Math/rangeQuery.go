package Math

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{make([]int, len(nums)+1)}
	copy(na.nums[1:], nums)

	for i := 0; i < len(nums); i++ {
		na.nums[i+1] += na.nums[i]
	}

	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}
