package traversal

func LongestCycle(edges []int) int {
	res := -1

	for i := 0; i < len(edges); i++ {

		if edges[i] != -1 {
			//fmt.Println("start", i)
			count := 0
			node := i
			isCycle := make(map[int]int)
			previous := -1
			for {
				//fmt.Println(previous, node)
				if isCycle[node] > 0 {
					// fmt.Println("count ", isCycle)
					count = isCycle[previous] - isCycle[node] + 1
					break
				}

				if edges[node] == -1 {
					count = 0
					break
				}

				count++
				isCycle[node] = count
				previous = node
				node = edges[node]
				edges[previous] = -1
			}

			if count > res && count != 0 {
				res = count
			}
		}
	}
	return res
}
