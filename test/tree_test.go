package test

import (
	"github.com/magiconair/properties/assert"
	"leetcode/Tree"
	"testing"
)

func TestFindInOrderSuccessor(t *testing.T) {
	testData := []struct {
		inputNode *Tree.NodeP
		exp       int
	}{
		{Tree.GenerateNodeP("20,9,25,5,12,null,null,null,null,11,14", 11), 12},
		{Tree.GenerateNodeP("20,9,25,5,12,null,null,null,null,11,14", 9), 11},
		{Tree.GenerateNodeP("20,9,25,5,12,null,null,null,null,11,14", 14), 20},
	}

	for _, td := range testData {
		result := Tree.FindInOrderSuccessor(td.inputNode)
		assert.Equal(t, result.Value, td.exp)
	}
}

func TestWidthOfBinaryTree(t *testing.T) {
	codex := Tree.Codec{}
	testData := []struct {
		root *Tree.TreeNode
		exp  int
	}{
		{codex.Deserialize("1,3,2,5,3,null,9"), 4},
		{codex.Deserialize("1,3,2,5,null,null,9,6,null,7"), 7},
		{codex.Deserialize("1,3,2,5"), 2},
		{codex.Deserialize("1,1,1,1,1,1,1,null,null,null,1,null,null,null,null,2,2,2,2,2,2,2,null,2,null,null,2,null,2"), 8},
	}

	for _, td := range testData {
		result := Tree.WidthOfBinaryTree(td.root)
		assert.Equal(t, result, td.exp)
	}
}

func TestLongestZigZag(t *testing.T) {
	codex := Tree.Codec{}
	testData := []struct {
		root *Tree.TreeNode
		exp  int
	}{
		{
			codex.Deserialize("1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1"),
			3,
		},
		{
			codex.Deserialize("1,1,1,null,1,null,null,1,1,null,1"),
			4,
		},
		{
			codex.Deserialize("1"),
			0,
		},
	}

	for _, td := range testData {
		result := Tree.LongestZigZag(td.root)
		assert.Equal(t, result, td.exp)
	}
}

func TestKthSmallest(t *testing.T) {
	codex := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		k        int
		expected int
	}{
		{codex.Deserialize("3,1,4,null,2"), 1, 1},
		{codex.Deserialize("5,3,6,2,4,null,null,1"), 3, 3},
	}

	for _, td := range testData {
		result := Tree.KthSmallest(td.root, td.k)
		assert.Equal(t, result, td.expected)
	}
}
func TestLowestCommonAncestor(t *testing.T) {
	codec := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		p        *Tree.TreeNode
		q        *Tree.TreeNode
		expected *Tree.TreeNode
	}{
		{
			codec.Deserialize("3,5,1,6,2,0,8,null,null,7,4"),
			codec.Deserialize("5"),
			codec.Deserialize("1"),
			codec.Deserialize("3"),
		},
		{
			codec.Deserialize("3,5,1,6,2,0,8,null,null,7,4"),
			codec.Deserialize("5"),
			codec.Deserialize("4"),
			codec.Deserialize("5"),
		},
	}

	for _, td := range testData {
		result := Tree.LowestCommonAncestor(td.root, td.p, td.q)
		assert.Equal(t, result.Val, td.expected.Val)
	}
}

func TestLowestCommonAncestorOfBST(t *testing.T) {
	codec := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		p        *Tree.TreeNode
		q        *Tree.TreeNode
		expected *Tree.TreeNode
	}{
		{
			codec.Deserialize("6,2,8,0,4,7,9,null,null,3,5"),
			codec.Deserialize("2"),
			codec.Deserialize("8"),
			codec.Deserialize("6"),
		},
		{
			codec.Deserialize("6,2,8,0,4,7,9,null,null,3,5"),
			codec.Deserialize("2"),
			codec.Deserialize("4"),
			codec.Deserialize("2"),
		},
	}

	for _, td := range testData {
		result := Tree.LowestCommonAncestorOfBST(td.root, td.p, td.q)
		assert.Equal(t, result.Val, td.expected.Val)
	}
}

func TestIsSubTree(t *testing.T) {
	code := Tree.Codec{}
	testData := []struct {
		root     *Tree.TreeNode
		subRoot  *Tree.TreeNode
		expected bool
	}{
		{code.Deserialize("3,4,5,1,2"), code.Deserialize("4,1,2"), true},
		{code.Deserialize("3,4,5,1,2,null,null,null,null,0"), code.Deserialize("4,1,2"), false},
		{code.Deserialize("1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,2"), code.Deserialize("1,null,1,null,1,null,1,null,1,null,1,2"), true},
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
