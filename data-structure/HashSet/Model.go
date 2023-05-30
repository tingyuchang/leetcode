package HashSet

type MyHashSet struct {
	keys []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		keys: []int{},
	}
}

func (this *MyHashSet) Add(key int) {
	if binarySearchInHashSet(this.keys, key) == -1 {
		l := 0
		r := len(this.keys) - 1

		for r >= l {
			mid := int(uint(l+r) >> 1)

			if this.keys[mid] < key {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		this.keys = append(this.keys, 0)

		for i := len(this.keys) - 1; i > l; i-- {
			this.keys[i] = this.keys[i-1]
		}

		this.keys[l] = key
	}
}

func (this *MyHashSet) Remove(key int) {
	index := binarySearchInHashSet(this.keys, key)
	if index != -1 {
		for i := index; i < len(this.keys)-1; i++ {
			this.keys[i] = this.keys[i+1]
		}

		this.keys = this.keys[:len(this.keys)-1]
	}
}

func (this *MyHashSet) Contains(key int) bool {
	result := binarySearchInHashSet(this.keys, key)

	return result != -1
}

// return -1 if cannot find
func binarySearchInHashSet(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := int(uint(l+r) >> 1)

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
