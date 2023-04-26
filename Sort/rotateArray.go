package Sort

import "fmt"

func Rotate(nums []int, k int) {
	rotate(nums, k)
}

/*

approach 1  using queue/stack to push dada from end of array

1,2,3,4,5,6,7 k = 3

queue [7, 6, 5] -> pop element to the front of array

TO (k+n)
SO (k+n)

approach 2 move elements in place

k = k%len(nums)
n = len(nums)
ans[(i+k)%n] = nums[i]

TO (n)
SO (n)

approach 3 reverse element

reverse  array

and reverse first k element and rest elements seperated


*/

func rotate(nums []int, k int) {

	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse := func(l, r int) {
		for r > l {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	reverse(0, len(nums)-1)
	fmt.Println(nums)
	reverse(0, k-1)
	fmt.Println(nums)
	reverse(k, len(nums)-1)
	fmt.Println(nums)
}
