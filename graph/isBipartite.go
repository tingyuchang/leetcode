package graph

import "fmt"

func isBipartite(graph [][]int) bool {
	/*
		Approach 1 DFS
		Approach 2 BFS
		Approach 3 UnionFinder
	*/

	for i := 0; i < len(graph); i++ {
		if !bfs(graph, i) {
			return false
		}
	}

	return true
}

func bfs(graph [][]int, start int) bool {
	queue := graph[start]
	levelMap := make(map[int]int)
	levelMap[start] = 1
	currentLevel := 2
	visited := make([]bool, len(graph))
	visited[start] = true
	for len(queue) > 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			next := queue[i]
			visited[next] = true
			// fmt.Println(levelMap, next, currentLevel)
			if levelMap[next] != 0 && levelMap[next] != currentLevel {
				return false
			}
			levelMap[next] = currentLevel

			for _, v := range graph[next] {
				if !visited[v] {
					queue = append(queue, v)
					visited[v] = true
				} else {
					if levelMap[v] == currentLevel {
						return false
					}
				}
			}
		}
		queue = queue[queueSize:]
		// fmt.Println("break", queue)
		if currentLevel == 1 {
			currentLevel = 2
		} else {
			currentLevel = 1
		}
	}
	fmt.Println(levelMap)
	return true
}
