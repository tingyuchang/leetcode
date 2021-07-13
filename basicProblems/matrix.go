package basicProblems

/*
// Write a function that accepts an integer N
// and returns a NxN spiral matrix.
// --- Examples
//   matrix(2)
//     [[1, 2],
//     [4, 3]]
//   matrix(3)
//     [[1, 2, 3],
//     [8, 9, 4],
//     [7, 6, 5]]
//  matrix(4)
//     [[1,   2,  3, 4],
//     [12, 13, 14, 5],
//     [11, 16, 15, 6],
//     [10,  9,  8, 7]]
 */

func matrix(n int) [][]int {
	// working direction
	cases := []string{"right", "down", "left", "up"}
	// return value
	var output = make([][]int, n)
 	// make empty []int
	for i:=0;i<n;i++{
		output[i] = make([]int, n)
	}
	// row: current row index
	// column: current column index
	row := 0
	column :=0

	direction := 0
	// set [0][0] = 1
	output[0][0] = 1
	for i:=0;i<n*n-1; i++ {
		switch cases[direction] {
		case "right":
			column++
		case "down":
			row++
		case"left":
			column--
		case"up":
			row--
		}
		output[row][column] = i+2
		// check changing direction or not
		nextRow := row
		nextColumn := column

		switch cases[direction] {
		case "right":
			nextColumn++
		case "down":
			nextRow++
		case"left":
			nextColumn--
		case"up":
			nextRow--
		}

		// check need changing direction or not
		// touch edge limit for slice or exist value in slice
		if nextColumn > n-1 || nextRow > n-1 || nextColumn < 0 || nextRow < 0 || output[nextRow][nextColumn] != 0 {
			if direction == 3 {
				direction = 0
			} else {
				direction++
			}
		}
	}

	return output
	
}