package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Tree"
	"testing"
)

func TestIsSubTree(t *testing.T) {
	code := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		subRoot  *Tree.TreeNode
		expected bool
	}{
		{code.Deserialize("3,4,5,1,2"), code.Deserialize("4,1,2"), true},
		{code.Deserialize("3,4,5,1,2,null,null,null,null,0"), code.Deserialize("4,1,2"), true},
	}

	for _, td := range testData {
		result := Tree.IsSubtree(td.root, td.subRoot)
		assert.Equal(t, result, td.expected)
	}
}

func TestMaxPathSum(t *testing.T) {
	code := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		expected int
	}{
		{code.Deserialize("1,2,3"), 6},
		{code.Deserialize("-10,9,20,null,null,15,7"), 42},
		{code.Deserialize("-3"), -3},
	}

	for _, td := range testData {
		result := Tree.MaxPathSum(td.root)
		assert.Equal(t, result, td.expected)
	}
}

func TestSerialize(t *testing.T) {

	node1 := Tree.TreeNode{Val: 1}
	node2 := Tree.TreeNode{Val: 2}
	node3 := Tree.TreeNode{Val: 3}
	node4 := Tree.TreeNode{Val: 4}
	node5 := Tree.TreeNode{Val: 5}

	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4
	node3.Right = &node5
	testData := []struct {
		root     *Tree.TreeNode
		expected string
	}{
		{&node1, "1,2,3,null,null,4,5"},
	}
	codec := Tree.Codec{}
	for _, td := range testData {
		result := codec.Serialize(td.root)
		assert.Equal(t, result, td.expected)
	}
}

func TestDeserialize(t *testing.T) {
	testData := []struct {
		data     string
		expected *Tree.TreeNode
	}{
		{},
	}
	codec := Tree.Codec{}
	for _, td := range testData {
		result := codec.Deserialize(td.data)
		assert.Equal(t, result, td.expected)
	}
}
