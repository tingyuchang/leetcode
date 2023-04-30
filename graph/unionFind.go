package graph

type UnionFind struct {
	group []int
	rank  []int
}

func NewUnionFind(n int) *UnionFind {
	group := make([]int, n)

	for i := range group {
		group[i] = i
	}

	return &UnionFind{
		group: group,
		rank:  make([]int, n),
	}
}

// find return node group
func (uf *UnionFind) find(node int) int {
	if uf.group[node] != node {
		uf.group[node] = uf.find(uf.group[node])
	}
	return uf.group[node]
}

// join
func (uf *UnionFind) join(node1, node2 int) {
	group1 := uf.find(node1)
	group2 := uf.find(node2)

	if group1 == group2 {
		return
	}

	if uf.rank[group1] > uf.rank[group2] {
		uf.group[group2] = group1
	} else if uf.rank[group1] < uf.rank[group2] {
		uf.group[group1] = group2
	} else {
		uf.group[group1] = group2
		uf.rank[group2] += 1
	}
}

func (uf *UnionFind) areConnected(node1, node2 int) bool {
	group1 := uf.find(node1)
	group2 := uf.find(node2)

	return group1 == group2
}
