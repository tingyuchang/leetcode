package main

import (
	"fmt"
	"leetcode/divide"
	"leetcode/mergeKLists"
)

func main() {
	fmt.Println("result:", divide.Divide(720, 2))
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