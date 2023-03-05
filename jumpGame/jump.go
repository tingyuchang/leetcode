package jumpGame

func CanJump(nums []int) bool {
	longest := 0
	for i, v := range nums {
		if longest >= i && i+v > longest {
			longest = i + v
		}

		if longest > len(nums)-1 {
			return true
		}
	}

	return longest >= len(nums)-1
	//n := len(nums)
	//if n == 1 {
	//	return true
	//}
	//
	//visited := make(map[int]bool)
	//q := []int{0}
	//
	//for len(q) > 0 {
	//	index := q[0]
	//	q = q[1:]
	//	if nums[index]+index >= n-1 {
	//		return true
	//	}
	//	// add possible
	//	for i := index + 1; i <= index+nums[index]; i++ {
	//		if !visited[i] {
	//			visited[i] = true
	//			q = append(q, i)
	//		}
	//	}
	//}
	//
	//return false
}

func JumpsIV(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	if arr[0] == arr[n-1] {
		return 1
	}

	counter := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := counter[arr[i]]; ok {
			counter[arr[i]] = append(counter[arr[i]], i)
		} else {
			counter[arr[i]] = []int{i}
		}
	}

	visited := make(map[int]bool)
	q := []struct {
		index int
		step  int
	}{
		{0, 0},
	}

	for len(q) != 0 {
		index, step := q[0].index, q[0].step
		q = q[1:]
		if index == n-1 {
			return step
		}

		for _, v := range counter[arr[index]] {
			if !visited[v] {
				visited[v] = true
				q = append(q, struct{ index, step int }{v, step + 1})
			}
		}
		// clean neighbor
		counter[arr[index]] = nil

		// add adjacent

		if index-1 >= 0 && !visited[index-1] {
			visited[index-1] = true
			q = append(q, struct{ index, step int }{index - 1, step + 1})
		}

		if index+1 < n && !visited[index+1] {
			visited[index+1] = true
			q = append(q, struct{ index, step int }{index + 1, step + 1})
		}
	}

	return len(arr)
}

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > len(nums)-1 {
		return 1
	}
	step := 1
	index := 0
	temp := nums[index+1 : index+1+nums[index]]

	maxIndex := 0
	for index+nums[index] < len(nums)-1 {
		max := 0
		for i, v := range temp {
			if i+v >= max {
				max = i + v
				maxIndex = i
			}
		}
		index = index + maxIndex + 1

		if index+nums[index] >= len(nums)-1 {
			step++
			break
		}
		temp = nums[index+1 : index+1+nums[index]]
		step++
	}
	return step
}
