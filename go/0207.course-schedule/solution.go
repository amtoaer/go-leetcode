// Created by amtoaer at 2025/12/18 20:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/course-schedule/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		in[pre[0]]++
	}
	visited := make([]bool, numCourses)
	var dfs func(idx int)
	dfs = func(idx int) {
		if visited[idx] {
			return
		}
		if in[idx] > 0 {
			in[idx]--
			if in[idx] != 0 {
				return
			}
		}
		visited[idx] = true
		for _, edge := range edges[idx] {
			dfs(edge)
		}
	}
	for i := range numCourses {
		if visited[i] || in[i] != 0 {
			continue
		}
		dfs(i)
	}
	for _, isVisited := range visited {
		if !isVisited {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numCourses := Deserialize[int](ReadLine(stdin))
	prerequisites := Deserialize[[][]int](ReadLine(stdin))
	ans := canFinish(numCourses, prerequisites)

	fmt.Println("\noutput:", Serialize(ans))
}
