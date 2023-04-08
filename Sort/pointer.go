package Sort

func KClosest(points [][]int, k int) [][]int {
	return kClosest(points, k)
}

func kClosest(points [][]int, k int) [][]int {
	res := make([][]int, 0)
	n := len(points)

	for i := n / 2; i >= 0; i-- {
		minPoint(points, i)
	}

	for k > 0 {
		points[0], points[len(points)-1] = points[len(points)-1], points[0]
		res = append(res, points[len(points)-1])
		points = points[:len(points)-1]
		minPoint(points, 0)
		k--
		// if need to consider same priority?
		//for distance(temp[0]) == distance(points[0]) {
		//	points[0], points[len(points)-1] = points[len(points)-1], points[0]
		//	temp = append(temp, points[len(points)-1])
		//	points = points[:len(points)-1]
		//	k--
		//}

	}

	return res
}

func minPoint(points [][]int, n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1
	var smallest int

	if l < len(points) && smallerDistance(points[l], points[n]) {
		smallest = l
	} else {
		smallest = n
	}

	if r < len(points) && smallerDistance(points[r], points[smallest]) {
		smallest = r
	}

	if smallest != n {
		points[n], points[smallest] = points[smallest], points[n]
		minPoint(points, smallest)
	}
}

func smallerDistance(a, b []int) bool {
	if distance(a) > distance(b) {
		return false
	}
	return true
}

func distance(a []int) int {
	return a[0]*a[0] + a[1]*a[1]
}
