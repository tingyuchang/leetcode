package data_structure

type MyStack struct {
	list []int
}

func constructor() MyStack {
	return MyStack{}
}

// Push pushes element x to the top of the stack.
func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
}

// Pop removes the element on the top of the stack and returns it.
func (this *MyStack) Pop() int {
	n := len(this.list)
	x := this.list[n-1]
	this.list = this.list[:n-1]
	return x
}

// Top returns the element on the top of the stack.
func (this *MyStack) Top() int {
	return this.list[len(this.list)-1]
}

// Empty returns true if the stack is empty, false otherwise.
func (this *MyStack) Empty() bool {
	if len(this.list) > 0 {
		return false
	}
	return true
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
