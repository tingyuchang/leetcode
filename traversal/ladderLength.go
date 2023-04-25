package traversal

func LadderLength(beginWord string, endWord string, wordList []string) int {
	return bfsLadderLength(beginWord, endWord, wordList)
}

/*
n => lenght of words
m => length of souce
T O(m*n)
S O(m)
*/

func bfsLadderLength(source, target string, words []string) int {
	if source == target {
		return 0
	}

	queue := make([]string, 0)
	queue = append(queue, source)
	backQueue := make([]string, 0)
	wordsMap := &map[string]struct{}{}

	for _, word := range words {
		(*wordsMap)[word] = struct{}{}
	}
	ans := 1

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		for word, _ := range *wordsMap {
			count := 0
			for i := range word {
				if word[i] != s[i] {
					count++
				}
				if count > 1 {
					break
				}
			}

			if count == 1 {
				delete(*wordsMap, word)
				if word == target {
					return ans + 1
				}

				backQueue = append(backQueue, word)
			}

		}

		if len(queue) == 0 {
			ans++
			queue, backQueue = backQueue, make([]string, 0)
		}

	}

	return ans
}

/*
n => lenght of words
m => length of souce
T O(m^2*n)
S O(m)
*/

func dfsLadderLength(s string, words []string, target string, step int, visited map[string]struct{}) int {
	ans := 21
	for _, word := range words {
		count := 0
		if _, ok := visited[word]; ok {
			continue
		}
		for i := range word {
			if word[i] != s[i] {
				count++
			}
			if count > 1 {
				break
			}
		}

		if count == 1 {
			if word == target {
				return step + 1
			}
			visited[word] = struct{}{}
			result := dfsLadderLength(word, words, target, step+1, visited)
			if result != -1 && result < ans {
				ans = result
			}
		}
	}

	if ans != 21 {
		return ans
	}

	return -1
}
