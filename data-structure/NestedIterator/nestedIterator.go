package NestedIterator

import "math"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	data interface{}
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	_, ok := this.data.(int)
	return ok
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check

func (this NestedInteger) GetInteger() int {
	v, ok := this.data.(int)

	if ok {
		return v
	}

	return math.MaxInt
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.data = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	v, ok := this.data.([]*NestedInteger)
	if ok {
		v = append(v, &elem)
		this.data = v
	}
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	v, ok := this.data.([]*NestedInteger)

	if ok {
		return v
	}

	return nil
}

type NestedIterator struct {
	nestedList []*NestedInteger
	nums       []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	n := &NestedIterator{nestedList: nestedList}
	n.nums = flatten(n.nestedList)
	return n
}

func flatten(nestedList []*NestedInteger) []int {
	res := make([]int, 0)
	for _, nestInt := range nestedList {
		if nestInt.IsInteger() {
			res = append(res, nestInt.GetInteger())
		} else {
			res = append(res, flatten(nestInt.GetList())...)
		}
	}

	return res
}

func (this *NestedIterator) Next() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	return len(this.nums) > 0
}
