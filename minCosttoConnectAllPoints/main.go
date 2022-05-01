package main

import "fmt"

/*
You are given an array points representing integer coordinates of some points on a 2D-plane,
where points[i] = [xi, yi].
The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|,
where |val| denotes the absolute value of val.
Return the minimum cost to make all points connected.
All points are connected if there is exactly one simple path between any two points.
*/

func main() {
	input1 := [][]int{[]int{0, 0}, []int{2, 2}, []int{3, 10}, []int{5, 2}, []int{7, 0}}
	input2 := [][]int{[]int{3, 12}, []int{-2, 5}, []int{-4, 1}}
	fmt.Println("[[0,0],[2,2],[3,10],[5,2],[7,0]] ", minCostConnectPoints(input1))
	fmt.Println("[[3,12],[-2,5],[-4,1]] ", minCostConnectPoints(input2))
}

func minCostConnectPoints(points [][]int) int {
	for _, v := range points {
		fmt.Println(v)
	}
	return 0
}
