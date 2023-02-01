package main

import (
	"fmt"
	"leetcode/DivideAndConquer"
	"leetcode/mergeKLists"
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
	a := []int{3, 5, 1, 6, 7, 2, 4, 9, 10, 8}
	fmt.Println(DivideAndConquer.MergeSort(a))
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
