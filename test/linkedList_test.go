package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/LinkedList"
	"testing"
)

func TestPartitionLinkList(t *testing.T) {
	testData := []struct {
		head     *LinkedList.ListNode
		x        int
		expected *LinkedList.ListNode
	}{
		{
			LinkedList.GenerateNodeFromArray([]int{1, 4, 3, 2, 5, 2}),
			3,
			LinkedList.GenerateNodeFromArray([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{2, 1}),
			2,
			LinkedList.GenerateNodeFromArray([]int{1, 2}),
		},
	}

	for _, td := range testData {
		result := LinkedList.Partition(td.head, td.x)
		assert.Equal(t, result.String(), td.expected.String())
	}
}

func TestDeleteDuplicatesII(t *testing.T) {
	testData := []struct {
		head     *LinkedList.ListNode
		expected *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 3, 4, 4, 5}),
			LinkedList.GenerateNodeFromArray([]int{1, 2, 5})},
		{LinkedList.GenerateNodeFromArray([]int{1, 1, 1, 2, 3}),
			LinkedList.GenerateNodeFromArray([]int{2, 3})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 2}),
			LinkedList.GenerateNodeFromArray([]int{1})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 2, 2}),
			LinkedList.GenerateNodeFromArray([]int{1})},
	}

	for _, td := range testData {
		result := LinkedList.DeleteDuplicatesII(td.head)
		assert.Equal(t, td.expected.String(), result.String())
	}
}

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	var testData = []struct {
		input    *LinkedList.ListNode
		expected string
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 1, 2}), "12"},
		{LinkedList.GenerateNodeFromArray([]int{1, 1, 2, 3, 3}), "123"},
		{LinkedList.GenerateNodeFromArray([]int{0, 0, 0, 0, 0}), "0"},
	}

	for _, tt := range testData {
		result := LinkedList.DeleteDuplicates(tt.input)
		assert.Equal(t, result.Show(), tt.expected)
	}
}

func TestLinkedListIsPalindrome(t *testing.T) {
	testData := []struct {
		input    *LinkedList.ListNode
		expected bool
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5, 6}), false},
		{LinkedList.GenerateNodeFromArray([]int{2, 2, 1, 2, 1, 1}), false},
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 2, 1}), true},
		{LinkedList.GenerateNodeFromArray([]int{1, 2}), false},
		{LinkedList.GenerateNodeFromArray([]int{1}), true},
		{LinkedList.GenerateNodeFromArray([]int{1, 0, 1}), true},
	}

	for _, td := range testData {
		result := LinkedList.IsPalindrome(td.input)
		assert.Equal(t, result, td.expected)
	}
}

func TestRemoveElements(t *testing.T) {
	testData := []struct {
		input    *LinkedList.ListNode
		input2   int
		expected *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{7, 7, 7, 7}), 7, nil},
		{nil, 1, nil},
	}

	for _, td := range testData {
		result := LinkedList.RemoveElements(td.input, td.input2)
		assert.Equal(t, result, td.expected)
	}
}

func TestGetIntersectionNode(t *testing.T) {
	var testData = []struct {
		input1 []int
		input2 []int
		input3 int
		input4 int
	}{
		{[]int{1, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1},
		{[]int{2, 6, 4}, []int{1, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 1, 0},
		{[]int{2, 2, 4, 5, 5}, []int{2, 2, 4, 5, 4}, 2, 2},
		{[]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3},
		{[]int{1, 3, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, 4, 0},
	}
	for _, tt := range testData {
		headA, headB, intersection := generateTestData(tt.input1, tt.input2, tt.input3, tt.input4)
		actual := LinkedList.GetIntersectionNode(headA, headB)
		assert.Equal(t, actual, intersection)
	}
}

func generateTestData(a, b []int, skipA, skipB int) (headA, headB, intersection *LinkedList.ListNode) {
	headA, headB = &LinkedList.ListNode{}, &LinkedList.ListNode{}
	node1 := headA
	node2 := headB
	m, n := 0, 0
	for i := 0; i < skipA; i++ {
		node1.Val = a[i]
		m++
		if i < skipA-1 {
			node1.Next = &LinkedList.ListNode{}
			node1 = node1.Next
		}
	}

	for i := 0; i < skipB; i++ {
		node2.Val = b[i]
		n++
		if i < skipB-1 {
			node2.Next = &LinkedList.ListNode{}
			node2 = node2.Next
		}
	}

	if skipA == len(a) && skipB == len(b) {
		fmt.Println("HeadA: ", headA.Show())
		fmt.Println("HeadB: ", headB.Show())
		return
	}
	intersection = &LinkedList.ListNode{}
	if skipA > 0 {
		node1.Next = intersection
	} else {
		headA = intersection
	}

	if skipB > 0 {
		node2.Next = intersection
	} else {
		headB = intersection
	}

	node3 := intersection
	for i, v := range a[skipA:] {
		node3.Val = v
		if i < len(a)-skipA-1 {
			node3.Next = &LinkedList.ListNode{}
			node3 = node3.Next
		}
	}

	fmt.Println("HeadA: ", headA.Show())
	fmt.Println("HeadB: ", headB.Show())
	return
}
