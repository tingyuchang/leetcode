package graph

func DistanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	return distanceLimitedPathsExist(n, edgeList, queries)
}

/*
approach 1: brute force
using dfs to traverse all paths

m := len(queries)
n edges
T O(m*n^n)

k := len(edgList)
S O(k)

*/

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	res := make([]bool, len(queries))
	pathsMap := make(map[int]map[int][]int)

	for _, v := range edgeList {
		if _, ok := pathsMap[v[0]]; ok {
			pathsMap[v[0]][v[1]] = append(pathsMap[v[0]][v[1]], v[2])
		} else {
			pathsMap[v[0]] = make(map[int][]int)
			pathsMap[v[0]][v[1]] = []int{v[2]}
		}

		if _, ok := pathsMap[v[1]]; ok {
			pathsMap[v[1]][v[0]] = append(pathsMap[v[1]][v[0]], v[2])
		} else {
			pathsMap[v[1]] = make(map[int][]int)
			pathsMap[v[1]][v[0]] = []int{v[2]}
		}

	}

	//fmt.Println(pathsMap)

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		//fmt.Println("start")
		if dfsDistanceLimitedPathsExist(pathsMap, query[0], query[1], query[2], make(map[int]bool)) {
			res[i] = true
		}
		//fmt.Println("end")
	}
	return res
}

func dfsDistanceLimitedPathsExist(pathsMap map[int]map[int][]int, start, target, limitation int, visited map[int]bool) bool {
	//fmt.Printf("%d->", start)
	if start == target {
		return true
	}

	visited[start] = true

	for k, costs := range pathsMap[start] {

		for _, cost := range costs {
			if cost < limitation && !visited[k] {
				if dfsDistanceLimitedPathsExist(pathsMap, k, target, limitation, visited) {
					return true
				}
			}
		}

	}
	//fmt.Printf("(%d false)", start)
	return false
}
