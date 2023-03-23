package traversal

import "fmt"

func MakeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	cache := make(map[int][]int)

	for _, v := range connections {
		cache[v[0]] = append(cache[v[0]], v[1])
		cache[v[1]] = append(cache[v[1]], v[0])
	}
	res := 0
	visited := make([]bool, n)
	cluster := 0
	for k, _ := range cache {
		if !visited[k] {
			dfsMakeConnections(cache, k, visited)
			cluster++
		}
	}

	for _, v := range visited {
		if !v {
			res++
		}
	}
	fmt.Println(visited, "res: ", res, "cluster: ", cluster)
	return res + cluster - 1
}

func dfsMakeConnections(cache map[int][]int, current int, visited []bool) {
	if visited[current] == true {
		return
	}

	visited[current] = true

	for _, v := range cache[current] {
		dfsMakeConnections(cache, v, visited)
	}
}
