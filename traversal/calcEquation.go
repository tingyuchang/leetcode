package traversal

import "fmt"

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	return calcEquation(equations, values, queries)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ans := make([]float64, len(queries))
	uf := NewUnionFindStr()
	equationMap := make(map[string][]int)
	for i, equation := range equations {
		x := equation[0]
		y := equation[1]
		uf.add(x)
		uf.add(y)
		uf.join(x, y)
		if _, ok := equationMap[x]; ok {
			equationMap[x] = append(equationMap[x], i)
		} else {
			equationMap[x] = []int{i}
		}

		if _, ok := equationMap[y]; ok {
			equationMap[y] = append(equationMap[y], i)
		} else {
			equationMap[y] = []int{i}
		}
	}
	uf.tidy()
	// calculate ration by each uf group
	valueMap := make(map[string]float64)
	for _, group := range uf.uniqueGroup() {
		// bfs
		subEquations := make([]int, 0)
		subEquations = append(subEquations, equationMap[group]...)
		queue := make([]int, 0)
		queue = append(queue, equationMap[group]...)
		visited := make(map[int]struct{})

		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			visited[n] = struct{}{}

			// a, b
			x := equations[n][0]
			y := equations[n][1]

			for _, v := range equationMap[x] {
				if _, ok := visited[v]; !ok {
					queue = append(queue, v)
					subEquations = append(subEquations, v)
				}
			}

			for _, v := range equationMap[y] {
				if _, ok := visited[v]; !ok {
					queue = append(queue, v)
					subEquations = append(subEquations, v)
				}
			}
		}

		valueMap[group] = 1.0
		for _, index := range subEquations {
			equation := equations[index]
			value := values[index]
			x := equation[0]
			y := equation[1]
			//fmt.Println(index, equation, value)
			if x == group || y == group {
				if x == group {
					valueMap[y] = 1.0 / value
				} else {
					valueMap[x] = value
				}
			} else {
				if v, ok := valueMap[x]; ok {
					valueMap[y] = v / value
				} else if v2, ok2 := valueMap[y]; ok2 {
					valueMap[x] = v2 * value
				} else {
					fmt.Printf("%v, %v all not exist\n", x, y)
				}

			}
		}
	}

	fmt.Println(valueMap)

	for i, query := range queries {
		x := query[0]
		y := query[1]
		if !uf.areConnected(x, y) {
			ans[i] = -1.0
		} else {
			if x == y {
				ans[i] = 1
				continue
			}
			ans[i] = valueMap[x] / valueMap[y]
		}
	}

	return ans
}

type UnionFindStr struct {
	group map[string]string
	rank  map[string]int
}

func NewUnionFindStr() *UnionFindStr {
	return &UnionFindStr{
		group: make(map[string]string),
		rank:  make(map[string]int),
	}
}

func (uf *UnionFindStr) add(s string) {
	if _, ok := uf.group[s]; !ok {
		uf.group[s] = s
	}
}

func (uf *UnionFindStr) tidy() {
	for i := range uf.group {
		uf.find(i)
	}
}

// find return node group
func (uf *UnionFindStr) exist(s string) bool {
	_, ok := uf.group[s]
	return ok
}

func (uf *UnionFindStr) find(s string) string {
	// check exist ot not
	if !uf.exist(s) {
		return ""
	}
	if uf.group[s] != s {
		uf.group[s] = uf.find(uf.group[s])
	}
	return uf.group[s]
}

// join
func (uf *UnionFindStr) join(s1, s2 string) {
	group1 := uf.find(s1)
	group2 := uf.find(s2)

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

func (uf *UnionFindStr) areConnected(s1, s2 string) bool {
	if !uf.exist(s1) || !uf.exist(s2) {
		return false
	}
	group1 := uf.find(s1)
	group2 := uf.find(s2)

	return group1 == group2
}

func (uf *UnionFindStr) uniqueGroup() []string {
	groupMap := make(map[string]struct{})
	result := make([]string, 0)
	for _, v := range uf.group {
		if _, ok := groupMap[v]; !ok {
			result = append(result, v)
			groupMap[v] = struct{}{}
		}
	}

	return result
}

/*

approach 1 flatten equations & value to hashMap

like {
	a: 6
	b: 3
	c: 1
}


approach2 UnionFinder



*/
