package data_structure

type MyQueue struct {
	stack MyStack
}

func Constructor() MyQueue {
	return MyQueue{}
}

// Push pushes element x to the back of the queue.
func (this *MyQueue) Push(x int) {
	this.stack.Push(x)
}

// Pop removes the element from the front of the queue and returns it.
func (this *MyQueue) Pop() int {
	temp := make([]int, 0)
	for !this.stack.Empty() {
		temp = append(temp, this.stack.Pop())
	}

	for i := len(temp) - 2; i >= 0; i-- {
		this.stack.Push(temp[i])
	}

	return temp[len(temp)-1]
}

// Peek returns the element at the front of the queue.
func (this *MyQueue) Peek() int {
	temp := make([]int, 0)
	for !this.stack.Empty() {
		temp = append(temp, this.stack.Pop())
	}

	for i := len(temp) - 1; i >= 0; i-- {
		this.stack.Push(temp[i])
	}

	return temp[len(temp)-1]
}

// Empty returns true if the queue is empty, false otherwise.
func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
