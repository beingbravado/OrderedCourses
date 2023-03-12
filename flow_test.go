package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	t.Run("cyclic dependency", func(t *testing.T) {
		resp, err := orderOfCourses(6, [][]int{
			{0, 1},
			{1, 2},
			{2, 3},
			{3, 4},
			{4, 0},
			{1, 5},
		})

		assert.Equal(t, []int{}, resp)
		assert.Equal(t, "cannot find the sequence: cyclic dependency", err.Error())
	})
	t.Run("Example 1", func(t *testing.T) {
		resp, err := orderOfCourses(2, [][]int{{1, 0}})

		assert.Equal(t, []int{0, 1}, resp)
		assert.Equal(t, nil, err)
	})
	t.Run("Example 2", func(t *testing.T) {
		resp, err := orderOfCourses(4, [][]int{
			{1, 0},
			{2, 0},
			{3, 1},
			{3, 2},
		})

		assert.Equal(t, []int{0, 2, 1, 3}, resp)
		assert.Equal(t, nil, err)
	})
	t.Run("Example 2", func(t *testing.T) {
		resp, err := orderOfCourses(1, [][]int{})

		assert.Equal(t, []int{0}, resp)
		assert.Equal(t, nil, err)
	})
	t.Run("hard test", func(t *testing.T) {
		resp, err := orderOfCourses(10, [][]int{
			{1, 0},
			{5, 1},
			{2, 0},
			{4, 2},
			{3, 2},
			{2, 6},
			{7, 6},
			{8, 7},
			{9, 7},
		})

		assert.Equal(t, []int{6, 7, 9, 8, 0, 2, 3, 4, 1, 5}, resp)
		assert.Equal(t, nil, err)
	})
}
