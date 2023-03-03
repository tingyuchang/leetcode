package LFU

import (
	"container/list"
)

type entry struct {
	key   int
	value int
	usage int
}

type LFUCache struct {
	Capacity    int
	Usage       int
	least       int
	frequencies map[int]*list.List
	cache       map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity,
		0,
		1,
		make(map[int]*list.List),
		make(map[int]*list.Element),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}

	if ele, ok := this.cache[key]; ok {

		ele.Value.(*entry).usage++
		usage := ele.Value.(*entry).usage
		this.frequencies[usage-1].Remove(ele)
		if this.frequencies[usage] == nil {
			this.frequencies[usage] = list.New()
		}

		newEle := this.frequencies[usage].PushFront(ele.Value)
		this.cache[key] = newEle

		if usage-1 == this.least && this.frequencies[usage-1].Len() == 0 {
			this.least++
		}
		return ele.Value.(*entry).value
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cache == nil {
		return
	}

	if ele, ok := this.cache[key]; ok {
		ele.Value.(*entry).value = value
		ele.Value.(*entry).usage++
		usage := ele.Value.(*entry).usage
		this.frequencies[usage-1].Remove(ele)
		if this.frequencies[usage] == nil {
			this.frequencies[usage] = list.New()
		}

		newEle := this.frequencies[usage].PushFront(ele.Value)
		this.cache[key] = newEle

		if usage-1 == this.least && this.frequencies[usage-1].Len() == 0 {
			this.least++
		}
		return
	}

	if this.frequencies[1] == nil {
		this.frequencies[1] = list.New()
	}

	ele := this.frequencies[1].PushFront(&entry{key: key, value: value, usage: 1})
	this.cache[key] = ele
	if this.Capacity != 0 && this.Usage >= this.Capacity {
		this.RemoveLeast()
	}
	this.least = 1
	this.Usage++
}

func (this *LFUCache) Remove(key int) {
	if this.cache == nil {
		return
	}

	if ele, ok := this.cache[key]; ok {
		this.removeElement(ele)
	}
}

func (this *LFUCache) RemoveLeast() {
	if this.cache == nil {
		return
	}
	ele := this.frequencies[this.least].Back()
	if ele != nil {
		this.removeElement(ele)
	}
}

func (this LFUCache) removeElement(e *list.Element) {
	kv := e.Value.(*entry)
	this.frequencies[kv.usage].Remove(e)
	delete(this.cache, kv.key)
	this.Usage--
}
