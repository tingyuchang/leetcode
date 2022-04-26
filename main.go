package main

import (
	"fmt"
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
	ints := []int{1, 2,3,4,5}
	strs := []string{"A", "B", "X"}
	PrintSlice(ints)
	PrintSlice(strs)
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

func myLog(input interface{}) {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	LstFlags := log.Ldate | log.Lshortfile
	iLog :=log.New(f, "LNum ", LstFlags)
	iLog.Println(input)
	iLog.SetFlags(log.Lshortfile|log.LstdFlags)
}