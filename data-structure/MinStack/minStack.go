package MinStack

type MinStack struct {
	minBuf []int
	data   []int
}

func Constructor() MinStack {
	return MinStack{
		minBuf: make([]int, 0),
		data:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.data = append(this.data, val)
		this.minBuf = append(this.minBuf, val)
	} else {
		this.data = append(this.data, val)

		if val < this.GetMin() {
			this.minBuf = append(this.minBuf, val)
		} else {
			this.minBuf = append(this.minBuf, this.GetMin())
		}
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minBuf = this.minBuf[:len(this.minBuf)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}

	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {

	if len(this.minBuf) == 0 {
		return 0
	}

	return this.minBuf[len(this.minBuf)-1]
}

/*
type MinStack struct {
	data []int
	heap []int
}

func Constructor() MinStack {
	return MinStack{data: []int{}, heap: []int{}}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.heap = append(this.heap, val)
	this.minHeapButtonUp()
}

func (this *MinStack) Pop() {
	removeEle := this.Top()
	this.removeFromHeap(removeEle)
	this.data = this.data[:len(this.data)-1]
	return
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.heap[0]
}

func (this *MinStack) minHeapButtonUp() {
	current := len(this.heap) - 1
	for current > 0 {
		parent := (current - 1) >> 1
		if this.heap[current] < this.heap[parent] {
			this.heap[current], this.heap[parent] = this.heap[parent], this.heap[current]
			current = parent
		} else {
			break
		}
	}
}

func (this *MinStack) removeFromHeap(n int) {
	idx := 0

	for i := 0; i < len(this.heap); i++ {
		if this.heap[i] == n {
			idx = i
			break
		}
	}

	this.heap = append(this.heap[:idx], this.heap[idx+1:]...)
	this.minHeapTopDown(idx)

}

func (this *MinStack) minHeapTopDown(n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var smallest int

	if l < len(this.heap) && this.heap[l] < this.heap[n] {
		smallest = l
	} else {
		smallest = n
	}

	if r < len(this.heap) && this.heap[r] < this.heap[smallest] {
		smallest = r
	}

	if smallest != n {
		this.heap[smallest], this.heap[n] = this.heap[n], this.heap[smallest]
		this.minHeapTopDown(smallest)
	}
}
*/
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
