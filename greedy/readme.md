# 605. Can Place Flowers
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.


Example 1:
```
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
```
Example 2:
```
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false
```

Constraints:

* 1 <= flowerbed.length <= 2 * 104
* flowerbed[i] is 0 or 1.
* There are no two adjacent flowers in flowerbed.
* 0 <= n <= flowerbed.length
# 435. Non-overlapping Intervals

Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Example 1:
```
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
```
Example 2:
```
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
```
Example 3:
```
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

Constraints:

* 1 <= intervals.length <= 105
* intervals[i].length == 2
* -5 * 104 <= starti < endi <= 5 * 104