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