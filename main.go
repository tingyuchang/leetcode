package main

import (
	"fmt"
	_0230213 "leetcode/0_Daily_Prac/20230213"
	"leetcode/mergeKLists"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
)

var name = _0230213.Name{}
var mergeSort = _0230213.MergeSort
var heapSort = _0230213.HeapSort
var insertionSort = _0230213.InsertionSort
var quickSort = _0230213.QuickSort
var binarySearch = _0230213.BinarySearch

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}
func main() {
	dailyPrac()
}

func dailyPrac() {
	n := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	m := []int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}
	re := regexp.MustCompile(`\d{8}`)
	fmt.Printf("%q\n", re.Find([]byte(reflect.TypeOf(name).PkgPath())))
	fmt.Println("Merge Sort: ", mergeSort(n))
	n = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println("Heap Sort: ", heapSort(n))
	n = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	fmt.Println("Insertion Sort: ", insertionSort(n))
	fmt.Println("Quick Sort: ", quickSort(m, 0, len(m)-1))
	sorted := []int{3, 5, 7, 9, 11}
	target1 := 7
	target2 := 11
	target3 := 6
	fmt.Printf("Binary Saerch in %v: \n\tTarget1 %v result: %v\n\tTarget2 %v result: %v\n\tTarget3 %v result: %v  \n", sorted, target1, binarySearch(sorted, target1), target2, binarySearch(sorted, target2), target3, binarySearch(sorted, target3))

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
