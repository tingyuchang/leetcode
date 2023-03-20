package Math

func CanFinish(numCourses int, prerequisites [][]int) bool {
	cache := make(map[int][]int)
	for i, _ := range prerequisites {
		cache[prerequisites[i][1]] = append(cache[prerequisites[i][1]], prerequisites[i][0])
	}

	memo := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if !isCycle(&cache, &memo, i) {
			return false
		}
	}

	return true
}

func isCycle(cache *map[int][]int, memo *[]int, index int) bool {
	if (*memo)[index] == 2 {
		return false
	}

	if (*memo)[index] == 1 {
		return true
	}

	(*memo)[index] = 2
	for _, v := range (*cache)[index] {
		if !isCycle(cache, memo, v) {
			return false
		}
	}

	(*memo)[index] = 1

	return true
}

func CanFinish2(numCourses int, prerequisites [][]int) bool {
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
