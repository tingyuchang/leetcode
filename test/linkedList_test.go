package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/LinkedList"
	"testing"
)

func TestAddTwoNumbersII(t *testing.T) {
	testData := []struct {
		l1  *LinkedList.ListNode
		l2  *LinkedList.ListNode
		exp *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{2, 3, 4}),
			LinkedList.GenerateNodeFromArray([]int{5, 6, 4}),
			LinkedList.GenerateNodeFromArray([]int{7, 9, 8}),
		},
		{LinkedList.GenerateNodeFromArray([]int{7, 2, 4, 3}),
			LinkedList.GenerateNodeFromArray([]int{5, 6, 4}),
			LinkedList.GenerateNodeFromArray([]int{7, 8, 0, 7}),
		},
	}

	for _, td := range testData {
		result := LinkedList.AddTwoNumbersII(td.l1, td.l2)
		assert.Equal(t, result.String(), td.exp.String())
	}
}

func TestAddTwoNumbers(t *testing.T) {

	var testData = []struct {
		l1  *LinkedList.ListNode
		l2  *LinkedList.ListNode
		exp *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{2, 3, 4}), LinkedList.GenerateNodeFromArray([]int{5, 6, 4}), LinkedList.GenerateNodeFromArray([]int{7, 9, 8})},
		{LinkedList.GenerateNodeFromArray([]int{2, 9, 2}), LinkedList.GenerateNodeFromArray([]int{1, 3, 8}), LinkedList.GenerateNodeFromArray([]int{3, 2, 1, 1})},
		{LinkedList.GenerateNodeFromArray([]int{9, 1, 6}), LinkedList.GenerateNodeFromArray([]int{0}), LinkedList.GenerateNodeFromArray([]int{9, 1, 6})},
		{LinkedList.GenerateNodeFromArray([]int{1, 8}), LinkedList.GenerateNodeFromArray([]int{0}), LinkedList.GenerateNodeFromArray([]int{1, 8})},
		{LinkedList.GenerateNodeFromArray([]int{9, 9, 9, 9, 9, 9, 9}), LinkedList.GenerateNodeFromArray([]int{9, 9, 9, 9}), LinkedList.GenerateNodeFromArray([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}

	for _, tt := range testData {
		result := LinkedList.AddTwoNumbers(tt.l1, tt.l2)
		assert.Equal(t, result.String(), tt.exp.String())
	}
}

func TestSwapPairs2(t *testing.T) {
	testData := []struct {
		root     *LinkedList.ListNode
		expected *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4}), LinkedList.GenerateNodeFromArray([]int{2, 1, 4, 3})},
		{
			LinkedList.GenerateNodeFromArray([]int{
				28, 34, 48, 74, 42, 49, 37, 59, 97, 96, 73, 44, 39, 50, 80, 30, 95, 26, 94, 88, 87, 84, 57, 47, 100, 56, 69, 27, 58, 2, 64, 86, 8, 43, 41, 32, 67, 51, 53, 101, 7, 76, 92, 45, 83, 6, 46, 40, 16, 66, 91, 1, 63, 89, 90, 4, 52, 65, 3, 70, 62, 29, 71, 15, 22, 93, 24, 25, 61, 82, 54, 60, 81, 14, 33, 85, 13, 17, 20, 31, 18, 79, 68, 10, 38, 11, 78, 72, 55, 36, 19, 99, 77, 5, 12, 35, 23, 21, 98,
			}),
			LinkedList.GenerateNodeFromArray([]int{
				34, 28, 74, 48, 49, 42, 59, 37, 96, 97, 44, 73, 50, 39, 30, 80, 26, 95, 88, 94, 84, 87, 47, 57, 56, 100, 27, 69, 2, 58, 86, 64, 43, 8, 32, 41, 51, 67, 101, 53, 76, 7, 45, 92, 6, 83, 40, 46, 66, 16, 1, 91, 89, 63, 4, 90, 65, 52, 70, 3, 29, 62, 15, 71, 93, 22, 25, 24, 82, 61, 60, 54, 14, 81, 85, 33, 17, 13, 31, 20, 79, 18, 10, 68, 11, 38, 72, 78, 36, 55, 99, 19, 5, 77, 35, 12, 21, 23, 98,
			}),
		},
	}

	for _, td := range testData {
		result := LinkedList.SwapPairs(td.root)
		assert.Equal(t, result.String(), td.expected.String())
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	testData := []struct {
		input    *LinkedList.ListNode
		target   int
		expected *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}), 2, LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 5})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2}), 1, LinkedList.GenerateNodeFromArray([]int{1})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2}), 2, LinkedList.GenerateNodeFromArray([]int{2})},
	}

	for _, td := range testData {
		result := LinkedList.RemoveNthFromEnd(td.input, td.target)
		for result.Next != nil {
			assert.Equal(t, result.Val, td.expected.Val)
			result = result.Next
			td.expected = td.expected.Next
		}
	}

}

func TestMiddleNode(t *testing.T) {
	testData := []struct {
		head     *LinkedList.ListNode
		expected *LinkedList.ListNode
	}{
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}),
			LinkedList.GenerateNodeFromArray([]int{3, 4, 5})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5, 6}),
			LinkedList.GenerateNodeFromArray([]int{4, 5, 6})},
		{LinkedList.GenerateNodeFromArray([]int{1, 2}),
			LinkedList.GenerateNodeFromArray([]int{2})},
		{LinkedList.GenerateNodeFromArray([]int{1}),
			LinkedList.GenerateNodeFromArray([]int{1})},
	}

	for _, td := range testData {
		result := LinkedList.MiddleNode(td.head)
		assert.Equal(t, result.String(), td.expected.String())
	}
}

func TestReverseListII(t *testing.T) {
	testData := []struct {
		head     *LinkedList.ListNode
		left     int
		right    int
		expected *LinkedList.ListNode
	}{
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}),
			2,
			4,
			LinkedList.GenerateNodeFromArray([]int{1, 4, 3, 2, 5}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{5}),
			1,
			1,
			LinkedList.GenerateNodeFromArray([]int{5}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{3, 5}),
			1,
			2,
			LinkedList.GenerateNodeFromArray([]int{5, 3}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3}),
			2,
			3,
			LinkedList.GenerateNodeFromArray([]int{1, 3, 2}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{3, 5}),
			1,
			1,
			LinkedList.GenerateNodeFromArray([]int{3, 5}),
		},
		{
			LinkedList.GenerateNodeFromArray([]int{1, 2, 3, 4, 5}),
			1,
			5,
			LinkedList.GenerateNodeFromArray([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, td := range testData {
		result := LinkedList.ReverseListII(td.head, td.left, td.right)
		assert.Equal(t, result.String(), td.expected.String())
	}
}

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
