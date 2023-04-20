package RandomizedSet

import (
	"math/rand"
)

type RandomizedSet struct {
	data map[int]int
	keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{data: make(map[int]int), keys: make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.data[val]

	if ok {
		return false
	}
	this.keys = append(this.keys, val)

	this.data[val] = len(this.keys) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.data[val]

	if !ok {
		return false
	}
	// swap with last item

	swapKey := this.keys[len(this.keys)-1]
	this.data[swapKey] = idx
	this.keys[idx] = swapKey
	this.keys = this.keys[:len(this.keys)-1]

	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	num := rand.Intn(len(this.keys))

	return this.keys[num]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
