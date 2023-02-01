package main

import (
	"fmt"
	"leetcode/insertionSort"
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
	fmt.Println(insertionSort.AddBinary([]int{1, 1, 1, 1}, []int{1, 0, 0, 1}))
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
