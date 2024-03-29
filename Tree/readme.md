# 114. Flatten Binary Tree to Linked List

Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.


Example 1:

![](https://assets.leetcode.com/uploads/2021/01/14/flaten.jpg)
```markdown
Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]
```
Example 2:

```
Input: root = []
Output: []
```
Example 3:
```
Input: root = [0]
Output: [0]
```

# 117. Populating Next Right Pointers in Each Node II

Given a binary tree
```
struct Node {
int val;
Node *left;
Node *right;
Node *next;
}
```
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.



Example 1:

```
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```
Example 2:
```
Input: root = []
Output: []
```

Constraints:

* The number of nodes in the tree is in the range [0, 6000].
* -100 <= Node.val <= 100


Follow-up:

* You may only use constant extra space.
* The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

# BST Successor Search

In a Binary Search Tree (BST), an Inorder Successor of a node is defined as the node with the smallest key greater than the key of the input node (see examples below). Given a node inputNode in a BST, you’re asked to write a function findInOrderSuccessor that returns the Inorder Successor of inputNode. If inputNode has no Inorder Successor, return null.

Explain your solution and analyze its time and space complexities.

![](https://www.pramp.com/img/content/img_02.png)

Example:
```
In the diagram above, for inputNode whose key = 11

Your function would return:

The Inorder Successor node whose key = 12
```
Constraints:

* [time limit] 5000ms
* [input] Node inputNode
* [output] Node

# 662. Maximum Width of Binary Tree

Given the root of a binary tree, return the maximum width of the given tree.

The maximum width of a tree is the maximum width among all levels.

The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes), where the null nodes between the end-nodes that would be present in a complete binary tree extending down to that level are also counted into the length calculation.

It is guaranteed that the answer will in the range of a 32-bit signed integer.


Example 1:

![](https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg)
```
Input: root = [1,3,2,5,3,null,9]
Output: 4
Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).
```
Example 2:

![](https://assets.leetcode.com/uploads/2022/03/14/maximum-width-of-binary-tree-v3.jpg)
```
Input: root = [1,3,2,5,null,null,9,6,null,7]
Output: 7
Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).
```
Example 3:

![](https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg)
```
Input: root = [1,3,2,5]
Output: 2
Explanation: The maximum width exists in the second level with length 2 (3,2).
```

Constraints:

The number of nodes in the tree is in the range [1, 3000].
-100 <= Node.val <= 100

# Unique Binary Search Trees

Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.


Example 1:

![](https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg)
```
Input: n = 3
Output: 5
```
Example 2:

```
Input: n = 1
Output: 1
```

Constraints:

1 <= n <= 19

# 1372. Longest ZigZag Path in a Binary Tree

You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.


Example 1:

```
Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
```
Example 2:

```
Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
```
Example 3:
```
Input: root = [1]
Output: 0
```

Constraints:

* The number of nodes in the tree is in the range [1, 5 * 104].
* 1 <= Node.val <= 100

# 173. Binary Search Tree Iterator

Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):

BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
int next() Moves the pointer to the right, then returns the number at the pointer.
Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.



Example 1:


Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // return 3
bSTIterator.next();    // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 20
bSTIterator.hasNext(); // return False


Constraints:

The number of nodes in the tree is in the range [1, 105].
0 <= Node.val <= 106
At most 105 calls will be made to hasNext, and next.


Follow up:

Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?

# 653. Two Sum IV - Input is a BST

Given the root of a binary search tree and an integer k, return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.

Example 1:

```
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true
```
Example 2:

```
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false
```

Constraints:

* The number of nodes in the tree is in the range [1, 104].
* -104 <= Node.val <= 104
* root is guaranteed to be a valid binary search tree.
* -105 <= k <= 105

# 701. Insert into a Binary Search Tree

You are given the root node of a binary search tree (BST) and a value to insert into the tree. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

Example 1:

```
Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:
```
Example 2:
```
Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]
```
Example 3:
```
Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]
```

Constraints:

The number of nodes in the tree will be in the range [0, 104].
-108 <= Node.val <= 108
All the values Node.val are unique.
-108 <= val <= 108
It's guaranteed that val does not exist in the original BST.


# 700. Search in a Binary Search Tree

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.


Example 1:

```
Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]
```
Example 2:

```
Input: root = [4,2,7,1,3], val = 5
Output: []
```

Constraints:

* The number of nodes in the tree is in the range [1, 5000].
* 1 <= Node.val <= 107
* root is a binary search tree.
* 1 <= val <= 107

# 429. N-ary Tree Level Order Traversal

Given an n-ary tree, return the level order traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).



Example 1:

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)
```
Input: root = [1,null,3,2,4,null,5,6]
Output: [[1],[3,2,4],[5,6]]
```
Example 2:

![](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)
```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
```

Constraints:

* The height of the n-ary tree is less than or equal to 1000
* The total number of nodes is between [0, 104]

# 116. Populating Next Right Pointers in Each Node


You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.



Example 1:

![](https://assets.leetcode.com/uploads/2019/02/14/116_sample.png)
```
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```
Example 2:

```
Input: root = []
Output: []
```

Constraints:

* The number of nodes in the tree is in the range [0, 212 - 1].
* -1000 <= Node.val <= 1000


# 617. Merge Two Binary Trees

You are given two binary trees root1 and root2.

Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return the merged tree.

Note: The merging process must start from the root nodes of both trees.


Example 1:

![](https://assets.leetcode.com/uploads/2021/02/05/merge.jpg)
```
Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]
```
Example 2:
```
Input: root1 = [1], root2 = [1,2]
Output: [2,2]
```

Constraints:

* The number of nodes in both trees is in the range [0, 2000].
* -104 <= Node.val <= 104

# 589. N-ary Tree Preorder Traversal

Given the root of an n-ary tree, return the preorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)


Example 1:

![](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)


```
Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]
```
Example 2:


```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
```

Constraints:

* The number of nodes in the tree is in the range [0, 104].
* 0 <= Node.val <= 104
* The height of the n-ary tree is less than or equal to 1000.
* 
# 230. Kth Smallest Element in a BST

Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

Example 1:

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)
```
Input: root = [3,1,4,null,2], k = 1
Output: 1
```
Example 2:

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)
```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

Constraints:

* The number of nodes in the tree is n.
* 1 <= k <= n <= 104
* 0 <= Node.val <= 104


Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

# 236. Lowest Common Ancestor of a Binary Tree

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”



Example 1:

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)
```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
```
Example 2:

![](https://assets.leetcode.com/uploads/2018/12/14/binarytree.png)
```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
```

Example 3:

![]()
```
Input: root = [1,2], p = 1, q = 2
Output: 1
```

Constraints:

* The number of nodes in the tree is in the range [2, 105].
* -109 <= Node.val <= 109
* All Node.val are unique.
* p != q
* p and q will exist in the tree.


# 235. Lowest Common Ancestor of a Binary Search Tree
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”


Example 1:

![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.

```
Example 2:

![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

```
Example 3:

```
Input: root = [2,1], p = 2, q = 1
Output: 2
```

Constraints:

* The number of nodes in the tree is in the range [2, 105].
* -109 <= Node.val <= 109
* All Node.val are unique.
* p != q
* p and q will exist in the BST.


# 572. Subtree of Another Tree
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:

![](https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg)

```
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```

Example 2:

![](https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg)
```
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
```
Constraints:

* The number of nodes in the root tree is in the range [1, 2000].
* The number of nodes in the subRoot tree is in the range [1, 1000].
* -104 <= root.val <= 104
* -104 <= subRoot.val <= 104

# 124. Binary Tree Maximum Path Sum
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.


Example 1:

![](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)

```
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
```
Example 2:

![](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)
```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
```

Constraints:

* The number of nodes in the root tree is in the range [1, 2000].
* The number of nodes in the subRoot tree is in the range [1, 1000].
* -104 <= root.val <= 104
* -104 <= subRoot.val <= 104


# 297. Serialize and Deserialize Binary Tree
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

Example 1:

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```
Example 2:
```
Input: root = []
Output: []
```

Constraints:

* The number of nodes in the tree is in the range [0, 104].
* -1000 <= Node.val <= 1000
# 129. Sum Root to Leaf Numbers
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/num1tree.jpg)
```
Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
```
Example 2:

![](https://assets.leetcode.com/uploads/2021/02/19/num2tree.jpg)
```
Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
```
# 958. Check Completeness of a Binary Tree

Given the root of a binary tree, determine if it is a complete binary tree.

In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.



Example 1:

![](https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-1.png)
```
Input: root = [1,2,3,4,5,6]
Output: true
Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
```
Example 2:

![](https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-2.png)
```
Input: root = [1,2,3,4,5,null,7]
Output: false
Explanation: The node with value 7 isn't as far left as possible.
>>>>>>> ed82c1f (ADD: 958 Check Completeness of a Binary Tree):Tree/README.md
```

Constraints:
* The number of nodes in the tree is in the range [1, 1000].
* 0 <= Node.val <= 9
* The depth of the tree will not exceed 10.

# 106. Construct Binary Tree from Inorder and Postorder Traversal

Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)
```
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
```
Example 2:
```
Input: inorder = [-1], postorder = [-1]
Output: [-1]
```

Constraints:

* 1 <= inorder.length <= 3000
* postorder.length == inorder.length
* -3000 <= inorder[i], postorder[i] <= 3000
* inorder and postorder consist of unique values.
* Each value of postorder also appears in inorder.
* inorder is guaranteed to be the inorder traversal of the tree.
* postorder is guaranteed to be the postorder traversal of the tree.

# 105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Example 1:
![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)
```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```
Example 2:
```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

Constraints:

* 1 <= preorder.length <= 3000
* inorder.length == preorder.length
* -3000 <= preorder[i], inorder[i] <= 3000
* preorder and inorder consist of unique values.
* Each value of inorder also appears in preorder.
* preorder is guaranteed to be the preorder traversal of the tree.
* inorder is guaranteed to be the inorder traversal of the tree.
# 98. Validate Binary Search Tree
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree
of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Example 1:

![](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)
```
Input: root = [2,1,3]
Output: true
```

Example 2:

![](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)
```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

Constraints:

* The number of nodes in the tree is in the range [1, 104].
8 -231 <= Node.val <= 231 - 1

# 652. Find Duplicate Subtrees
Given the root of a binary tree, return all duplicate subtrees.

For each kind of duplicate subtrees, you only need to return the root node of any one of them.

Two trees are duplicate if they have the same structure with the same node values.



Example 1:

![](https://assets.leetcode.com/uploads/2020/08/16/e1.jpg)
```
Input: root = [1,2,3,4,null,2,4,null,null,4]
Output: [[2,4],[4]]
```
Example 2:

![](https://assets.leetcode.com/uploads/2020/08/16/e2.jpg)
```markdown
Input: root = [2,1,1]
Output: [[1]]
```

Example 3:

![](https://assets.leetcode.com/uploads/2020/08/16/e33.jpg)
```markdown
Input: root = [2,2,2,3,null,3,null]
Output: [[2,3],[3]]
```
# 427. Construct Quad Tree
Given a n * n matrix grid of 0's and 1's only. We want to represent the grid with a Quad-Tree.

Return the root of the Quad-Tree representing the grid.

Notice that you can assign the value of a node to True or False when isLeaf is False, and both are accepted in the answer.

A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:

* val: True if the node represents a grid of 1's or False if the node represents a grid of 0's.
* isLeaf: True if the node is leaf node on the tree or False if the node has the four children.
```
class Node {
public boolean val;
public boolean isLeaf;
public Node topLeft;
public Node topRight;
public Node bottomLeft;
public Node bottomRight;
}
```
We can construct a Quad-Tree from a two-dimensional area using the following steps:

If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
Recurse for each of the children with the proper sub-grid.

![](https://assets.leetcode.com/uploads/2020/02/11/new_top.png)

Example 1:

![](https://assets.leetcode.com/uploads/2020/02/11/grid1.png)
```markdown
Input: grid = [[0,1],[1,0]]
Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
Explanation: The explanation of this example is shown below:
Notice that 0 represnts False and 1 represents True in the photo representing the Quad-Tree.
```
![](https://assets.leetcode.com/uploads/2020/02/12/e1tree.png)

Example 2:

![](https://assets.leetcode.com/uploads/2020/02/12/e2mat.png)
![](https://assets.leetcode.com/uploads/2020/02/12/e2tree.png)
```markdown
Input: grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
Output: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
Explanation: All values in the grid are not the same. We divide the grid into four sub-grids.
The topLeft, bottomLeft and bottomRight each has the same value.
The topRight have different values so we divide it into 4 sub-grids where each has the same value.
Explanation is shown in the photo below:
```
Constraints:

n == grid.length == grid[i].length
n == 2x where 0 <= x <= 6

# 257. Binary Tree Paths
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

Example 1:

![](https://assets.leetcode.com/uploads/2021/03/12/paths-tree.jpg)
```
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
```
Example 2:
```
Input: root = [1]
Output: ["1"]
```

Constraints:

- The number of nodes in the tree is in the range [1, 100].
- -100 <= Node.val <= 100

# 109. Convert Sorted List to Binary Search Tree
Given the head of a singly linked list where elements are sorted in ascending order, convert it to a
height-balanced
binary search tree.


Example 1:

![](https://assets.leetcode.com/uploads/2020/08/17/linked.jpg)
```
Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
```
Example 2:
```
Input: head = []
Output: []
```

Constraints:

- The number of nodes in head is in the range [0, 2 * 104].
- -105 <= Node.val <= 105

# 107. Binary Tree Level Order Traversal II

Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to
right, level by level from leaf to root).

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]
```

Example 2:

```
Input: root = [1]
Output: [[1]]
```

Example 3:

```
Input: root = []
Output: []
```

Constraints:

- The number of nodes in the tree is in the range [0, 2000].
- -1000 <= Node.val <= 1000

# 103. Binary Tree Zigzag Level Order Traversal

Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to
right, then right to left for the next level and alternate between).

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
```

Example 2:

```
Input: root = [1]
Output: [[1]]
```

Example 3:

```
Input: root = []
Output: []
```

Constraints:

- The number of nodes in the tree is in the range [0, 2000].
- -100 <= Node.val <= 100

# 110. Balanced Binary Tree

Given a binary tree, determine if it is height-balanced
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by
more than one.

Example 1:

![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

Example 2:

![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

Example 3:

```
Input: root = []
Output: true
```

Constraints:

- The number of nodes in the tree is in the range [0, 5000].
- -104 <= Node.val <= 104

# 108. Convert Sorted Array to Binary Search Tree

Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced
binary search tree.

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg)
![](https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg)

```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

```

Example 2:

![](https://assets.leetcode.com/uploads/2021/02/18/btree.jpg)

```
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
```

Constraints:

- 1 <= nums.length <= 104
- -104 <= nums[i] <= 104
- nums is sorted in a strictly increasing order.

# 101. Symmetric Tree

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)

```
Input: root = [1,2,2,3,4,4,3]
Output: true
```

Example 2:
![](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)

```markdown
Input: root = [1,2,2,null,3,null,3]
Output: false
```

Constraints:

- The number of nodes in the tree is in the range [1, 1000].
- -100 <= Node.val <= 100

Follow up: Could you solve it both recursively and iteratively?

# 226. Invert Binary Tree

Given the root of a binary tree, invert the tree, and return its root.

Example 1:
![]()

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

Example 2:
![]()

```
Input: root = [2,1,3]
Output: [2,3,1]
```

Example 3:

```
Input: root = []
Output: []
```

Constraints:

- The number of nodes in the tree is in the range [0, 100].
- -100 <= Node.val <= 100

# 94. Binary Tree Inorder Traversal

Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
![](https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)

Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

# 783. Minimum Distance Between BST Nodes

# 530. Minimum Absolute Difference in BST

Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two
different nodes in the tree.

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg)

```
Input: root = [4,2,6,1,3]
Output: 1
```

Example 2:

![](https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg)

```
Input: root = [1,0,48,null,null,12,49]
Output: 1
```

Constraints:

- The number of nodes in the tree is in the range [2, 104].
- 0 <= Node.val <= 105