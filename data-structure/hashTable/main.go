package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type HashTable struct {
	size    int
	dataSet []Data
}

// Insert data to dataset and return number of index
// return -1 if not success
func (h *HashTable) Insert(key, value string) int {
	d := Data{key, value}
	index := h.linerHashIndex(h.hashKey(d.Key), d.Key)
	if index != -1 {
		h.dataSet[index] = d
	}
	return index
}

// Search
func (h *HashTable) Search(key string) Data {
	index := h.hashKey(key)

	if h.dataSet[index].Key == key {
		return h.dataSet[index]
	}

	for {
		index++
		if h.dataSet[index].isEmpty() {
			break
		}
		if h.dataSet[index].Key == key {
			return h.dataSet[index]
		}
	}

	return Data{}
}

// Delete
func (h *HashTable) Delete(key string) error {
	index := h.hashKey(key)

	if h.dataSet[index].Key == key {
		h.dataSet[index] = Data{}
		return nil
	}

	for {
		index++
		if h.dataSet[index].isEmpty() {
			break
		}
		if h.dataSet[index].Key == key {
			h.dataSet[index] = Data{}
			return nil
		}
	}

	return fmt.Errorf("Not found")
}

func (h *HashTable) hashKey(key string) int {
	var sum int32 = 0
	for _, v := range key {
		sum += v
	}

	return int(sum % int32(h.size))
}

// linerHashIndex return index in dataSet
// returns -1 if can't get valid number
func (h *HashTable) linerHashIndex(index int, key string) int {
	for index < h.size {
		if h.dataSet[index].isEmpty() || (h.dataSet[index].Key == key) {
			break
		}
		index++
	}

	if index == (h.size-1) && !(h.dataSet[index].isEmpty() || (h.dataSet[index].Key == key)) {
		index = -1
	}

	return index
}

func NewHashTable(size int) *HashTable {
	h := new(HashTable)
	h.size = size
	h.dataSet = make([]Data, h.size)
	return h
}

type Data struct {
	Key   string
	Value string
}

func (d *Data) isEmpty() bool {
	return *d == Data{}
}

var SIZE = 100

func main() {
	/*
		hashTable := NewHashTable(10000)
		fmt.Println("init dataSet")
		hashTable.Insert("1", "2")
		hashTable.Insert("Name", "Matt")
		hashTable.Insert("Surname", "Chang")
		fmt.Println(hashTable.Search("1"))
		fmt.Println(hashTable.Search("Name"))
		err := hashTable.Delete("Name")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(hashTable.Search("Name"))
	*/

	creatHashTables()
	createMaps()

}

func createRandomData() (keys, values []string) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer f.Close()
	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	rand.Seed(seed)

	keys = make([]string, SIZE)
	values = make([]string, SIZE)

	for i := 0; i < SIZE; i++ {
		keys[i] = strconv.Itoa(rand.Int())
		values[i] = strconv.Itoa(rand.Int())
	}
	return
}

func creatHashTables() {
	keys, values := createRandomData()
	hashTable := NewHashTable(SIZE * SIZE)
	for i := 0; i < SIZE; i++ {
		hashTable.Insert(keys[i], values[i])
	}
}

func createMaps() {
	keys, values := createRandomData()
	mapData := make(map[string]string)
	for i := 0; i < SIZE; i++ {
		mapData[keys[i]] = values[i]
	}
	fmt.Println(mapData)
}
