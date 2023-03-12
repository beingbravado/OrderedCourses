package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number of courses: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Enter the size of prereqs array: ")
	var m int
	fmt.Scanln(&m)

	fmt.Println("Enter prereqs array: ")
	fmt.Printf("Enter enter the course and its dependent course: ex. 2 4")
	var prereqs = [][]int{}
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		prereqs = append(prereqs, []int{x, y})
	}

	fmt.Printf("Here is the sequence in which you should complete your course")
	sol, err := orderOfCourses(n, prereqs)
	if err != nil {
		print(err.Error())
	} else {
		for _, value := range sol {
			fmt.Printf("- %d\n", value)
		}
	}
}
