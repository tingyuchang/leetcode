package traversal

func LongestCycle(edges []int) int {
	res := -1
	// -1: stop
	// 0: not visited
	// other: count
	visited := make(map[int]int)
	for i := 0; i < len(edges); i++ {

		if visited[edges[i]] == 0 && edges[i] != -1 {
			//fmt.Println("start", i)
			count := 0
			node := i
			isCycle := make(map[int]int)
			previous := -1
			for {
				//fmt.Println(previous, node)
				if visited[node] > 0 {
					//fmt.Println(isCycle, "pre: ", previous, "node: ", node)
					//fmt.Println("visited ", visited)
					//fmt.Println("count ", isCycle)
					count = isCycle[previous] - isCycle[node] + 1
					for k, _ := range isCycle {
						visited[k] = -1
					}
					break
				}

				count++
				isCycle[node] = count
				visited[node] = count
				previous = node
				node = edges[node]
				if node == -1 {
					//fmt.Println("end", previous)
					for k, _ := range isCycle {
						visited[k] = -1
					}
					count = 0
					break
				}
			}

			if count > res && count != 0 {
				res = count
			}
		}
	}

	return res
}
