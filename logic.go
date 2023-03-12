package main

import "errors"

func traverse(node *Node, visited []bool) ([]int, error) {
	if visited[node.num] {
		if !completedCourses[node.num] {
			return []int{}, errors.New("cyclic dependency")
		}
		return []int{}, nil
	}
	visited[node.num] = true

	out := []int{}
	for _, adjNode := range node.next {
		res, err := traverse(adjNode, visited)
		if err != nil {
			return []int{}, err
		}
		out = append(out, res...)
	}
	out = append(out, node.num)
	completedCourses[node.num] = true

	return out, nil
}
