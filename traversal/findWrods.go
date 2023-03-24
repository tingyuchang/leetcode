package traversal

func FindWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	cache := make(map[byte][][]int)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cache[board[i][j]] = append(cache[board[i][j]], []int{i, j})
		}
	}

	for _, word := range words {
		for _, coordinate := range cache[word[0]] {

			visited := make([][]bool, len(board))
			for i, _ := range visited {
				visited[i] = make([]bool, len(board[0]))
			}

			if dfsFindWord(board, word, 0, coordinate[0], coordinate[1], visited) {
				res = append(res, word)
				break
			}
		}
	}

	return res
}

func dfsFindWord(board [][]byte, target string, currentIndex, x, y int, visited [][]bool) bool {

	if x < 0 || x >= len(board) ||
		y < 0 || y >= len(board[0]) ||
		currentIndex >= len(target) ||
		visited[x][y] ||
		board[x][y] != target[currentIndex] {
		return false
	}

	if currentIndex == len(target)-1 {
		return true
	}

	visited[x][y] = true
	if dfsFindWord(board, target, currentIndex+1, x+1, y, visited) {
		return true
	}
	if dfsFindWord(board, target, currentIndex+1, x-1, y, visited) {
		return true
	}
	if dfsFindWord(board, target, currentIndex+1, x, y+1, visited) {
		return true
	}
	if dfsFindWord(board, target, currentIndex+1, x, y-1, visited) {
		return true
	}
	visited[x][y] = false

	return false

}
