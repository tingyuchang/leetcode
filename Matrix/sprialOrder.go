package Matrix

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}

	sum := n * n
	i := 0
	x, y := 0, -1

	layer := 0
	a, b := n, n-1
	currentLayer := a
	currentIndex := 0
	currentDirection := 0

	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for i < sum {
		x, y = x+directions[currentDirection][0], y+directions[currentDirection][1]
		res[x][y] = i + 1
		currentIndex++

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

	}

	return res
}

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

func spiralOrder2(matrix [][]int) []int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	ans := make([]int, 0)
	numOfRow := len(matrix)
	numOfCol := len(matrix[0])

	steps := []int{numOfCol, numOfRow - 1}

	indexOfDirection := 0
	x, y := 0, -1 // start position

	for steps[indexOfDirection%2] != 0 {
		for i := 0; i < steps[indexOfDirection%2]; i++ {
			x += directions[indexOfDirection][0]
			y += directions[indexOfDirection][1]
			ans = append(ans, matrix[x][y])
		}

		steps[indexOfDirection%2]--
		indexOfDirection = (indexOfDirection + 1) % 4

	}

	return ans
}