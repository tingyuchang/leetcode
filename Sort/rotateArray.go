package Sort

func Rotate(nums []int, k int) {
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	k = k % len(nums)

	if k == 0 {
		return
	}

	for i := 0; i < len(nums); i++ {

	}

}

//func rotate(nums []int, k int) {
//
//	k = k % len(nums)
//
//	if k == 0 {
//		return
//	}
//
//	temp := append([]int{}, nums[:len(nums)-k]...)
//
//	for i := 0; i < len(nums); i++ {
//		if i < k {
//			nums[i] = nums[len(nums)-k+i]
//		} else {
//			nums[i] = temp[i-k]
//		}
//	}
//
//}
