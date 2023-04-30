package graph

import (
	"fmt"
	"sort"
)

func MaxNumEdgesToRemove(n int, edges [][]int) int {
	return maxNumEdgesToRemove(n, edges)
}

func maxNumEdgesToRemove(n int, edges [][]int) int {

	totalCount := 0

	ufA := NewUnionFindMaxGroup(n)
	ufB := NewUnionFindMaxGroup(n)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})

	currentA := 0
	currentB := 0
	for _, edge := range edges {
		switch edge[0] {
		// Both
		case 3:
			ufA.join(edge[1]-1, edge[2]-1)
			ufB.join(edge[1]-1, edge[2]-1)
		// Bob Only
		case 2:
			ufB.join(edge[1]-1, edge[2]-1)
		// Alice Only
		case 1:
			ufA.join(edge[1]-1, edge[2]-1)
		}
		fmt.Println("Alice Max: ", ufA.totalGroup, currentA)
		fmt.Println("Bob Max: ", ufB.totalGroup, currentB)
		if currentA != ufA.totalGroup || currentB != ufB.totalGroup {
			totalCount++
		}
		currentA = ufA.totalGroup
		currentB = ufB.totalGroup
		// check current group's status

	}

	if ufA.totalGroup != 1 || ufB.totalGroup != 1 {
		return -1
	}

	return len(edges) - totalCount
}
