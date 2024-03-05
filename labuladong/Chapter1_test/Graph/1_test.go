package Tree

import (
	"fmt"
	"testing"
)

//课程表
//https://leetcode.cn/problems/course-schedule/description/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0, numCourses)
	}
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}

	onPath := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	path := make([]int, 0, numCourses)
	valid := false

	for i := 0; i < numCourses; i++ {
		canFinishHelper(graph, &onPath, &visited, &path, i, &valid)
	}

	return !valid
}

func canFinishHelper(graph [][]int, onPath, visited *[]bool, path *[]int, s int, valid *bool) {
	if (*onPath)[s] {
		*valid = true
		return
	}
	if (*visited)[s] || *valid {
		return
	}
	(*visited)[s] = true
	(*onPath)[s] = true
	*path = append(*path, s)
	for _, b := range graph[s] {
		canFinishHelper(graph, onPath, visited, path, b, valid)
	}
	(*onPath)[s] = false
}

//课程表 II
//https://leetcode.cn/problems/course-schedule-ii/description/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := range graph {
		graph[i] = make([]int, 0, numCourses)
	}
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}

	onPath := make([]bool, numCourses)
	visited := make([]bool, numCourses)
	path := make([]int, 0, numCourses)
	valid := false

	for i := 0; i < numCourses; i++ {
		canFinishHelper(graph, &onPath, &visited, &path, i, &valid)
	}

	if valid {
		return []int{}
	}

	return path
}

func Test_1(t *testing.T) {
	//fmt.Println(canFinish(2, [][]int{{1, 0}}))
	//fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	//fmt.Println(findOrder(2, [][]int{{0, 1}}))
	//fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	//fmt.Println(findOrder(1, [][]int{}))
}
