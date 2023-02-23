package traversal

type status int

const (
	unvisited = iota
	visited
	partiallyVistited
)

var Fuel = 0

func MinimumFuelCost(roads [][]int, seats int) int64 {
	Fuel = 0
	adjList := make(map[int][]int)
	for _, edge := range roads {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visitedList := make(map[int]status)

	for i := 0; i < len(adjList); i++ {
		dfs(i, adjList, &visitedList, seats)
	}

	return int64(Fuel)
}

func dfs(node int, adjList map[int][]int, visitedList *map[int]status, seats int) int {
	neighbourArr, ok := adjList[node]
	if !ok {
		return 0
	}

	if (*visitedList)[node] == partiallyVistited {
		return 0
	}

	if (*visitedList)[node] == visited {
		return 0
	}

	(*visitedList)[node] = partiallyVistited

	var people int
	if node != 0 {
		people = 1
	}
	for _, neighbour := range neighbourArr {
		people += dfs(neighbour, adjList, visitedList, seats)
	}

	(*visitedList)[node] = visited
	if node == 0 {
		return 0
	}

	var cost int
	if people > seats {
		if people%seats == 0 {
			cost = people / seats
		} else {
			cost = people/seats + 1
		}
	} else {
		cost = 1
	}

	Fuel += cost
	return people
}
