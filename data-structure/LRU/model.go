package LRU

import "container/list"

// not thread safe

type Key int
type entry struct {
	key   Key
	value int
}

type LRUCache struct {
	Capacity int
	ll       *list.List
	cache    map[Key]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.New(),
		make(map[Key]*list.Element),
	}
}

func (this *LRUCache) Get(key Key) int {
	if this.cache == nil {
		return -1
	}

	if ele, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key Key, value int) {
	if this.cache == nil {
		this.cache = make(map[Key]*list.Element)
		this.ll = list.New()
	}

	if ee, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}

	ele := this.ll.PushFront(&entry{key: key, value: value})
	this.cache[key] = ele
	if this.Capacity != 0 && this.ll.Len() > this.Capacity {
		this.RemoveOldest()
	}
}

func (this *LRUCache) Remove(key Key) {
	if this.cache == nil {
		return
	}

	if ele, ok := this.cache[key]; ok {
		this.removeElement(ele)
	}
}

func (this *LRUCache) RemoveOldest() {
	if this.cache == nil {
		return
	}

	ele := this.ll.Back()

	if ele != nil {
		this.removeElement(ele)
	}
}

func (this *LRUCache) Len() int {
	if this.cache == nil {
		return 0
	}
	return this.ll.Len()
}

func (this *LRUCache) Clear() {
	this.ll = nil
	this.cache = nil
}

func (this LRUCache) removeElement(e *list.Element) {
	this.ll.Remove(e)

	kv := e.Value.(*entry)

	delete(this.cache, kv.key)
}
