package Matrix

func SpiralOrder(matrix [][]int) []int {
	sum := len(matrix) * len(matrix[0])

	res := make([]int, sum)

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	i := 0
	x, y := 0, -1

	layer := 0
	a, b := len(matrix[0]), len(matrix)-1
	currentLayer := a
	currentIndex := 0
	currentDirection := 0
	for i < sum {

		x, y = x+directions[currentDirection][0], y+directions[currentDirection][1]
		res[i] = matrix[x][y]

		currentIndex++
		//fmt.Printf("a: %d, b: %d currentIndex: %d, currentLayer: %d \n", a, b, currentIndex, currentLayer)
		if currentIndex == currentLayer {
			currentDirection++
			if currentDirection == len(directions) {
				currentDirection = 0
			}

			layer++
			if layer%2 == 1 {
				currentLayer = b
				a--
			} else {
				currentLayer = a
				b--
			}
			currentIndex = 0
		}

		i++
		//fmt.Println(res)
	}

	return res
}
