package pascalsTriangle

func Generate(numRows int) [][]int {
	var result [][]int

	for i := 1; i <= numRows; i++ {
		if i == 1 {
			result = append(result, []int{1})
			continue
		}

		if i == 2 {
			result = append(result, []int{1, 1})
			continue
		}

		temp := make([]int, i)
		temp[0] = 1
		temp[i-1] = 1
		for j, v := range result[i-2] {
			if j == i-2 {
				break
			}
			temp[j+1] = v + result[i-2][j+1]
		}

		result = append(result, temp)

	}

	return result
}

func GenerateV2(numRows int) [][]int {
	var result [][]int
	var temp []int
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			result = append(result, []int{1})
			continue
		}

		if i == 2 {
			result = append(result, []int{1, 1})
			temp = []int{2}
			continue
		}

		temp = append([]int{1}, temp...)
		temp = append(temp, 1)
		result = append(result, temp)

		temp2 := make([]int, len(temp)-1)
		for i := 0; i < len(temp)-1; i++ {
			temp2[i] = temp[i] + temp[i+1]
		}
		temp = temp2
	}

	return result
}
