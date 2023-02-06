package main

import (
	"fmt"
	_0230206 "leetcode/0_Daily_Prac/20230206"
	"leetcode/mergeKLists"
	"leetcode/pascalsTriangle"
	"log"
	"os"
	"path"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}
func main() {
	n := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println("Merge Sort: ", _0230206.MergeSort(n))
	fmt.Println("Heap Sort: ", _0230206.HeapSort(n))
	sorted := []int{3, 5, 7, 9, 11}
	fmt.Println(_0230206.BinarySearch(sorted, 7))
	fmt.Println(_0230206.BinarySearch(sorted, 11))
	fmt.Println(_0230206.BinarySearch(sorted, 6))
	fmt.Println("Pascal's triangle", pascalsTriangle.GenerateV2(5))
	fmt.Println("Pascal's triangle", pascalsTriangle.GetRow(5))
}

func generateList(firstNodes []int) *mergeKLists.ListNode {
	l1 := mergeKLists.ListNode{}
	tempNode1 := &l1

	for _, v := range firstNodes {
		tempNode1.Next = &mergeKLists.ListNode{Val: v}
		tempNode1 = tempNode1.Next
	}
	return l1.Next
}

func myLog(input interface{}) {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	LstFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstFlags)
	iLog.Println(input)
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
}
