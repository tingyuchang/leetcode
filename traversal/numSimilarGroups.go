package traversal

func NumSimilarGroups(strs []string) int {
	return numSimilarGroups(strs)
}

func numSimilarGroups(strs []string) int {
	ans := 0

	strsMap := make(map[string]bool)

	for i := range strs {
		strsMap[strs[i]] = true
	}

	for str, available := range strsMap {
		if available {
			queue := []string{str}
			backQueue := []string{}

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				strsMap[current] = false

				for str2, available2 := range strsMap {
					if available2 {
						count := 0
						// compare str2 & current's difference
						for i := range str2 {
							if str2[i] != current[i] {
								count++
							}

							if count > 2 {
								break
							}
						}
						if count == 2 {
							backQueue = append(backQueue, str2)
						}
					}
				}

				if len(queue) == 0 {
					queue, backQueue = backQueue, []string{}
				}
			}
			ans++
		}
	}

	return ans
}

/*

approach 1 bfs to find similar string and add it into group

n: size of strs
m: string length
TO(n^2M)
SO(3n)
*/
