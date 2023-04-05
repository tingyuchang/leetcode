# 2439. Minimize Maximum of Array

You are given a 0-indexed array nums comprising of n non-negative integers.

In one operation, you must:

* Choose an integer i such that 1 <= i < n and nums[i] > 0.
* Decrease nums[i] by 1.
* Increase nums[i - 1] by 1.
Return the minimum possible value of the maximum integer of nums after performing any number of operations.



Example 1:

```
Input: nums = [3,7,1,6]
Output: 5
Explanation:
One set of optimal operations is as follows:
1. Choose i = 1, and nums becomes [4,6,1,6].
2. Choose i = 3, and nums becomes [4,6,2,5].
3. Choose i = 1, and nums becomes [5,5,2,5].
   The maximum integer of nums is 5. It can be shown that the maximum number cannot be less than 5.
   Therefore, we return 5.
```
Example 2:

```
Input: nums = [10,1]
Output: 10
Explanation:
It is optimal to leave nums as is, and since 10 is the maximum value, we return 10.
```

Constraints:

* n == nums.length
* 2 <= n <= 105
* 0 <= nums[i] <= 109

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