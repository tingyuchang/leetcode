package graph

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := make(map[int]int)

	for i := range nums {
		if _, ok := counter[nums[i]]; !ok {
			counter[nums[i]] = i
		}
	}

	uf := NewUnionFindMaxGroup(len(nums))

	for k, v := range counter {
		if _, ok := counter[k-1]; ok {
			uf.join(v, counter[k-1])
		}
	}

	return uf.maxGroup
}

type UnionFindMaxGroup struct {
	group      []int
	groupSize  []int
	rank       []int
	maxGroup   int
	totalGroup int
}

func NewUnionFindMaxGroup(n int) *UnionFindMaxGroup {
	group := make([]int, n)
	groupSize := make([]int, n)

	for i := range group {
		group[i] = i
		groupSize[i] = 1
	}

	return &UnionFindMaxGroup{
		group:      group,
		groupSize:  groupSize,
		maxGroup:   1,
		totalGroup: n,
		rank:       make([]int, n),
	}
}

func (uf *UnionFindMaxGroup) find(n int) int {
	if uf.group[n] != n {
		uf.group[n] = uf.find(uf.group[n])
	}

	return uf.group[n]
}

func (uf *UnionFindMaxGroup) join(n1, n2 int) {
	group1 := uf.find(n1)
	group2 := uf.find(n2)

	if group1 == group2 {
		return
	}

	if uf.rank[group1] > uf.rank[group2] {
		uf.group[group2] = group1
		uf.groupSize[group1] += uf.groupSize[group2]
		if uf.groupSize[group1] > uf.maxGroup {
			uf.maxGroup = uf.groupSize[group1]
		}
	} else {
		uf.group[group1] = group2
		uf.groupSize[group2] += uf.groupSize[group1]
		if uf.groupSize[group2] > uf.maxGroup {
			uf.maxGroup = uf.groupSize[group2]
		}

		if uf.rank[group1] == uf.rank[group2] {
			uf.rank[group2]++
		}
	}

	uf.totalGroup--
}

func (uf *UnionFindMaxGroup) areConnected(n1, n2 int) bool {
	group1 := uf.find(n1)
	group2 := uf.find(n2)

	return group1 == group2
}
