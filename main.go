package main

import (
	"fmt"
	"leetcode/basicProblems/tree"
	"leetcode/mergeKLists"
)

func main() {
	node1 := &tree.Node{
		Value: 20,
		Children: nil,
	}

	node2 := &tree.Node{
		Value: 0,
		Children: []*tree.Node{
			&tree.Node{
				Value:12,
			},
			&tree.Node{
				Value:-2,
			},
			&tree.Node{
				Value:1,
			},
		},
	}
	node3 := &tree.Node{
		Value: 40,
		Children: nil,
	}
	node4 := &tree.Node{
		Value: 15,
		Children: []*tree.Node{
			&tree.Node{
				Value:-2,
			},
		},
	}

	node1.Children = append(node1.Children, node2, node3, node4)

	tree1 := tree.Tree{
		Root: node1,
	}

	for _,v := range tree1.BreadthFirst() {
		fmt.Printf("%v ", v.Value)
	}

	fmt.Println("")

	for _,v := range tree1.DepthFirst() {
		fmt.Printf("%v ", v.Value)
	}



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