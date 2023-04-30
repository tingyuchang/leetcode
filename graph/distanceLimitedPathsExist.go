package graph

import (
	"sort"
)

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


approach 2 Disjoint Set Union(DSU) or Union-Find
*/

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	uf := NewUnionFind(n)
	ans := make([]bool, len(queries))
	queriesWithIndex := make([][]int, len(queries))
	edgesIndex := 0

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	for i := range queries {
		queriesWithIndex[i] = append(queries[i], i)
	}
	sort.Slice(queriesWithIndex, func(i, j int) bool {
		return queriesWithIndex[i][2] < queriesWithIndex[j][2]
	})

	for i := 0; i < len(queriesWithIndex); i++ {
		p := queriesWithIndex[i][0]
		q := queriesWithIndex[i][1]
		limit := queriesWithIndex[i][2]
		orgIndex := queriesWithIndex[i][3]

		for edgesIndex < len(edgeList) && edgeList[edgesIndex][2] < limit {
			node1 := edgeList[edgesIndex][0]
			node2 := edgeList[edgesIndex][1]
			uf.join(node1, node2)
			edgesIndex++
		}

		ans[orgIndex] = uf.areConnected(p, q)
	}

	return ans

}

func DFSDistanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
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
