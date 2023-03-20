package Math

func CanFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	cache := make(map[int][]int)
	memo := make(map[int]bool)
	for _, v := range prerequisites {

		k, ok := cache[v[0]]
		if !ok {
			cache[v[0]] = []int{v[1]}
		} else {
			cache[v[0]] = append(k, v[1])
		}

	}

	for i := 0; i < numCourses; i++ {
		requires, ok := cache[i]
		if ok {
			for _, v := range requires {
				visited := make(map[int]bool)
				if !dfs(&cache, &visited, &memo, v) {
					return false
				}
			}
		}

		memo[i] = true

	}

	return true
}

func dfs(cache *map[int][]int, visited *map[int]bool, memo *map[int]bool, course int) bool {
	next, ok := (*cache)[course]

	if ok {
		if (*visited)[course] {
			return false
		}
		(*visited)[course] = true
		for _, v := range next {
			if (*memo)[v] {
				continue
			}
			subVisited := make(map[int]bool)
			for k, v := range *visited {
				subVisited[k] = v
			}
			if !dfs(cache, &subVisited, memo, v) {
				return false
			}
		}
		return true
	}
	return true
}
