# 19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```
Example 2:

```
Input: head = [1], n = 1
Output: []
```
Example 3:

```
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:

* The number of nodes in the list is sz.
* 1 <= sz <= 30
* 0 <= Node.val <= 100
* 1 <= n <= sz


Follow up: Could you do this in one pass?

# 1367. Linked List in Binary Tree

Given a binary tree root and a linked list with head as the first node.

Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary tree otherwise return False.

In this context downward path means a path that starts at some node and goes downwards.


Example 1:

![]()
```
Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: true
Explanation: Nodes in blue form a subpath in the binary Tree.  
```
Example 2:

![]()
```
Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: true
```
Example 3:

```
Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
Output: false
Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.
```

Constraints:

* The number of nodes in the tree will be in the range [1, 2500].
* The number of nodes in the list will be in the range [1, 100].
* 1 <= Node.val <= 100 for each node in the linked list and binary tree.

# 876. Middle of the Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.


Example 1:

![](https://assets.leetcode.com/uploads/2021/07/23/lc-midlist1.jpg)
```
Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
```
Example 2:

![](https://assets.leetcode.com/uploads/2021/07/23/lc-midlist2.jpg)
```
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
```

Constraints:

* The number of nodes in the list is in the range [1, 100].
* 1 <= Node.val <= 100
* 
# 92. Reverse Linked List II

Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.


Example 1:

```
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
```
Example 2:
```
Input: head = [5], left = 1, right = 1
Output: [5]
```

Constraints:

* The number of nodes in the list is n.
* 1 <= n <= 500
* -500 <= Node.val <= 500
* 1 <= left <= right <= n


# 86. Partition List

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example 1:

!()[https://assets.leetcode.com/uploads/2021/01/04/partition.jpg]

```
Input: head = [1,4,3,2,5,2], x = 3
Output: [1,2,2,4,3,5]
```
Example 2:
```
Input: head = [2,1], x = 2
Output: [1,2]
```

Constraints:

* The number of nodes in the list is in the range [0, 200].
* -100 <= Node.val <= 100
* -200 <= x <= 200

# 82. Remove Duplicates from Sorted List II

Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:

![](https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg)
```
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]
```
Example 2:

![](https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg)
```
Input: head = [1,1,1,2,3]
Output: [2,3]
```

Constraints:

* The number of nodes in the list is in the range [0, 300].
* -100 <= Node.val <= 100
* The list is guaranteed to be sorted in ascending order.
* 
# 83. Remove Duplicates from Sorted List
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:

Input: head = [1,1,2]
Output: [1,2]

Example 2:

Input: head = [1,1,2,3,3]
Output: [1,2,3]

Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.

# 382. Linked List Random Node
Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.

Implement the Solution class:

Solution(ListNode head) Initializes the object with the head of the singly-linked list head.
int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of the list should be equally likely to be chosen.


Example 1:

![](https://assets.leetcode.com/uploads/2021/03/16/getrand-linked-list.jpg)
```
Input
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 3, 2, 2, 3]

Explanation
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // return 1
solution.getRandom(); // return 3
solution.getRandom(); // return 2
solution.getRandom(); // return 2
solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
```

Constraints:

* The number of nodes in the linked list will be in the range [1, 104].
* -104 <= Node.val <= 104
* At most 104 calls will be made to getRandom.
# 142. Linked List Cycle II
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.



Example 1:
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```
Example 2:

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)
```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```
Example 3:

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```
Constraints:

* The number of the nodes in the list is in the range [0, 104].
* -105 <= Node.val <= 105
* pos is -1 or a valid index in the linked-list.


# 143. Reorder List

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.


Example 1:
![](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)
```
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```
Example 2:
![](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)
```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

Constraints:

* The number of nodes in the list is in the range [1, 5 * 104].
* 1 <= Node.val <= 1000

# 234. Palindrome Linked List
Given the head of a singly linked list, return true if it is a
palindrome
or false otherwise.


Example 1:
![](https://assets.leetcode.com/uploads/2021/03/03/pal1linked-list.jpg)
```markdown
Input: head = [1,2,2,1]
Output: true
```
Example 2:

![](https://assets.leetcode.com/uploads/2021/03/03/pal2linked-list.jpg)
```
Input: head = [1,2]
Output: false
```

# 206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)
```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Input: head = []
Output: []
```

# 203. Remove Linked List Elements
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Example 1:

![](https://assets.leetcode.com/uploads/2021/03/06/removelinked-list.jpg)
```
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
```
Example 2:
```
Input: head = [], val = 1
Output: []
```
Example 3:
```
Input: head = [7,7,7,7], val = 7
Output: []
```
# 160. Intersection of Two Linked Lists

Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.

For example, the following two linked lists begin to intersect at node c1:
![](https://assets.leetcode.com/uploads/2021/03/05/160_statement.png)

The test cases are generated such that there are no cycles anywhere in the entire linked structure.

Note that the linked lists must retain their original structure after the function returns.

Custom Judge:

The inputs to the judge are given as follows (your program is not given these inputs):

intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
listA - The first linked list.
listB - The second linked list.
skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program. If you correctly return the intersected node, then your solution will be accepted.

Example 1:

![](https://assets.leetcode.com/uploads/2021/03/05/160_example_1_1.png)

```js
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
- Note that the intersected node's value is not 1 because the nodes with value 1 in A and B (2nd node in A and 3rd node in B) are different node references. In other words, they point to two different locations in memory, while the nodes with value 8 in A and B (3rd node in A and 4th node in B) point to the same location in memory.
```

Example 2:

![](https://assets.leetcode.com/uploads/2021/03/05/160_example_2.png)
```js
Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
```

Example 3:

![](https://assets.leetcode.com/uploads/2021/03/05/160_example_3.png)

```js
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
```

Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?

# 141. Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.



Example 1:

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)
```
nput: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```
Example 2:

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

```js
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
```

Example 3:

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

```js
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

```

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.


# 21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:

![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)
```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```
Example 2:
```
Input: list1 = [], list2 = []
Output: []
```
Example 3:
```
Input: list1 = [], list2 = [0]
Output: [0]
```

Constraints:

* The number of nodes in both lists is in the range [0, 50].
* -100 <= Node.val <= 100
* Both list1 and list2 are sorted in non-decreasing order.