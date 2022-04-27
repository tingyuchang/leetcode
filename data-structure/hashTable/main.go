package main

import "fmt"

type HashTable interface {
	// Insert can be Update
	Insert(d Data)
	Search(key string) string
	Delete(key string)
}

type Data struct {
	Key   string
	VaLue string
}

func hashKey(key string) int {
	var sum int32 = 0
	for _, v := range key {
		sum += v
	}

	return int(sum % int32(SIZE))
}

func linerHashIndex(index int) int {

}

var SIZE = 10000
var dataSet []string

func main() {
	fmt.Println(hashKey("2Jelel"))
	fmt.Println(hashKey("dd"))
	fmt.Println(hashKey("123456789011233"))
}
