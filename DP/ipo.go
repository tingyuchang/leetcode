package DP

import (
	"sort"
)

type Project struct {
	Capital int
	Profit  int
}

type ProjectList []Project

func (p ProjectList) Len() int {
	return len(p)
}

func (p ProjectList) Less(i, j int) bool {
	return p[i].Capital < p[j].Capital
}

func (p ProjectList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(capital)
	projects := make([]Project, 0)
	for i := 0; i < len(capital); i++ {
		project := Project{
			capital[i],
			profits[i],
		}

		projects = append(projects, project)
	}

	sort.Sort(ProjectList(projects))
	ptr := 0
	q := make([]int, 0)
	for i := 0; i < k; i++ {
		for ptr < n && projects[ptr].Capital <= w {
			index := insert(q, projects[ptr].Profit)
			if index >= len(q) {
				q = append(q, projects[ptr].Profit)
			} else {
				q = append(q[:index+1], q[index:]...)
				q[index] = projects[ptr].Profit
			}
			ptr++
		}

		if len(q) == 0 {
			break
		}

		w += q[len(q)-1]
		q = q[:len(q)-1]
	}
	return w
}

func insert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for r >= l {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return l
}

// TLE -> can be optimized

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	quickSort(profits, capital, 0, len(capital)-1)
	for k > 0 {
		// find capital[x] <= w
		// which x has maximum net profit

		start, end := -1, -1

		for i := 0; i < len(capital); i++ {
			v := capital[i]
			if v <= w {
				if start == -1 {
					start = i
				}
				end = i
			}

			if v > w {
				end = i - 1
				break
			}
		}

		if start == -1 && end == -1 {
			break
		}

		max := -1000000000
		project := -1
		for i := start; i <= end; i++ {
			if profits[i] > max {
				max = profits[i]
				project = i
			}
		}

		if project == -1 {
			break

		}

		profit := profits[project]
		profits[project] = -10000000000
		w += profit
		k--

	}
	return w
}

func quickSort(profits, capital []int, p, r int) {
	if r > p {
		q := partition(profits, capital, p, r)
		quickSort(profits, capital, p, q-1)
		quickSort(profits, capital, q+1, r)
	}
}

func partition(profits, capital []int, p, r int) int {
	x := capital[r]
	i := p - 1

	for j := p; j < r; j++ {
		if capital[j] < x {
			i++
			capital[i], capital[j] = capital[j], capital[i]
			profits[i], profits[j] = profits[j], profits[i]
		}
	}

	i++
	capital[i], capital[r] = capital[r], capital[i]
	profits[i], profits[r] = profits[r], profits[i]
	return i
}
