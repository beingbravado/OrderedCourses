package main

import (
	"github.com/pingcap/errors"
)

// orderOfCourses function uses topological sort
func orderOfCourses(n int, prereqs [][]int) ([]int, error) {
	graph := createGraph(n, prereqs)

	// maintaining a track of courses which have been visited
	// and completed. Useful in edge cases of graph
	var trackCourses []*Track

	// initializing the values
	for i := 0; i < n; i++ {
		trackCourses = append(trackCourses, &Track{
			visited:   false,
			completed: false,
		})
		graph.AddNode(i)
	}

	var courseOrder []int
	for i := 0; i < n; i++ {
		if !trackCourses[i].visited {
			path, err := traverse(graph.nodes[i], trackCourses)
			if err != nil {
				return []int{}, errors.Wrap(err, "cannot find the sequence")
			}

			courseOrder = append(courseOrder, path...)
		}
	}

	return reverseArray(courseOrder), nil
}

// reverseArray functions reverses the array in O(len(array))
func reverseArray(array []int) []int {
	n := len(array)
	var revArray []int
	for i := range array {
		revArray = append(revArray, array[n-i-1])
	}
	return revArray
}

// traverse is a recursive function which traverses through all the edges in graph
// Time Complexity O(nodes+edges)
func traverse(node *Node, trackCourses []*Track) ([]int, error) {
	if trackCourses[node.num].visited {
		if !trackCourses[node.num].completed {
			return []int{}, errors.New("cyclic dependency")
		}
		return []int{}, nil
	}
	trackCourses[node.num].visited = true

	out := []int{}
	// for each of the adjacent nodes, call the function recursively
	// if it hasn't yet been trackCourses
	for i := range node.next {
		res, err := traverse(node.next[i], trackCourses)
		if err != nil {
			return []int{}, err
		}
		out = append(out, res...)
	}
	out = append(out, node.num)
	trackCourses[node.num].completed = true

	return out, nil
}
