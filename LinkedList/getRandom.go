package LinkedList

import "math/rand"

type Solution struct {
	lists []*ListNode
}

func Constructor(head *ListNode) Solution {
	solution := Solution{lists: make([]*ListNode, 0)}
	node := head
	for node != nil {
		solution.lists = append(solution.lists, node)
		node = node.Next
	}
	return solution
}

func (this *Solution) GetRandom() int {
	return this.lists[rand.Intn(this.Len())-1].Val
}

func (this *Solution) Len() int {
	return len(this.lists)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
