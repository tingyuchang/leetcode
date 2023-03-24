package traversal

type Trie struct {
	word     string
	children map[byte]*Trie
}

func FindWords(board [][]byte, words []string) []string {
	root := &Trie{
		children: make(map[byte]*Trie),
	}
	// build Trie

	for _, word := range words {
		current := root
		for i := 0; i < len(word); i++ {
			node, ok := current.children[word[i]]

			if !ok {
				node = &Trie{
					children: make(map[byte]*Trie),
				}
				current.children[word[i]] = node
			}
			current = node

		}
		current.word = word
	}

	res := make([]string, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if node, ok := root.children[board[i][j]]; ok {
				backtracking(i, j, root, node, board, &res)
			}
		}
	}

	return res
}

func backtracking(x, y int, parent, node *Trie, board [][]byte, res *[]string) {
	char := board[x][y]

	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	board[x][y] = '#'

	next := [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, 1},
		[]int{0, -1},
	}

	for _, v := range next {
		newX, newY := x+v[0], y+v[1]
		if newX >= 0 && newX < len(board) &&
			newY >= 0 && newY < len(board[0]) {
			if nextNode, ok := node.children[board[newX][newY]]; ok {
				backtracking(newX, newY, node, nextNode, board, res)
			}
		}
	}

	board[x][y] = char
	if len(node.children) == 0 {
		delete(parent.children, char)
	}

}

func findWords(board [][]byte, words []string) []string {
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
