package Tree

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	result := make([]string, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		for _, node := range queue {
			if node == nil {
				result = append(result, "null")
			} else {
				result = append(result, strconv.Itoa(node.Val))
			}

			if node != nil {
				temp = append(temp, node.Left)
				temp = append(temp, node.Right)
			}
		}
		queue = temp
	}

	// remove unnecessary null

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "null" {
			break
		}
		result = result[:i]
	}

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")

	if len(vals) == 0 || vals[0] == "null" {
		return nil
	}

	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   rootVal,
	}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if vals[i] != "null" {
			value, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{
				nil, nil, value,
			}
			queue = append(queue, node.Left)
		}
		i++
		if vals[i] != "null" {
			value, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{
				nil, nil, value,
			}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func levelOrder(root *TreeNode) string {
	res := ""
	queue := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	queue = append(queue, root)

	var current *[]*TreeNode
	var background *[]*TreeNode
	current = &queue
	background = &queue2

	for len(*current) > 0 {
		node := (*current)[0]
		*current = (*current)[1:]

		if node == nil {
			res += "null,"
			continue
		} else {
			res += strconv.Itoa(node.Val) + ","
		}

		if node.Left != nil {

			*background = append(*background, node.Left)
		} else {
			*background = append(*background, nil)
		}
		if node.Right != nil {

			*background = append(*background, node.Right)
		} else {
			*background = append(*background, nil)
		}

		if len(*current) == 0 {
			current, background = background, current
		}

	}

	// remove last ","
	res = res[:len(res)-1]

	return res
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
