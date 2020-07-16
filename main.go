package main

import (
	"fmt"
	"leetcode/mergeKLists"
)

func main() {
	fmt.Println("Finish: ", mergeKLists.MergeKLists([]*mergeKLists.ListNode{
		generateList([]int{1,4,5}),
		generateList([]int{1,3,4}),
		generateList([]int{2,6}),
	}))
}

func generateList(firstNodes[]int) (*mergeKLists.ListNode) {
	l1 := mergeKLists.ListNode{}
	tempNode1 := &l1

	for _,v := range firstNodes {
		tempNode1.Next = &mergeKLists.ListNode{Val: v}
		tempNode1 = tempNode1.Next
	}


	return l1.Next

}