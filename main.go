package main

import "github.com/pingcap/errors"

var completedCourses map[int]bool

func orderOfCourses(n int, prereqs [][]int) ([]int, error) {
	g := NewGraph()

	var res []int
	var visited []bool
	completedCourses = map[int]bool{}

	// creating graph for the courses and initializing the values
	for i := 0; i < n; i++ {
		visited = append(visited, false)
		g.AddNode(i)
	}
	for _, edge := range prereqs {
		g.AddEdge(edge[1], edge[0])
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			path, err := traverse(g.nodes[i], visited)
			if err != nil {
				return []int{}, errors.Wrap(err, "cannot find the sequence")
			}

			res = append(res, path...)
		}
	}

	return reverseArray(res), nil
}

func reverseArray(array []int) []int {
	n := len(array)
	var resp []int
	for i := range array {
		resp = append(resp, array[n-1-i])
	}
	return resp
}
