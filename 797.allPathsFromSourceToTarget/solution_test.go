package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=797 lang=golang
 * @lcpr version=20004
 *
 * [797] 所有可能的路径
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res     [][]int
		visited = make([]bool, len(graph))
	)
	var dfs func(i int, route []int)
	dfs = func(i int, route []int) {
		if visited[i] {
			return
		}
		visited[i] = true
		route = append(route, i)
		if i == len(graph)-1 {
			res = append(res, append([]int{}, route...))
		} else {
			for _, edge := range graph[i] {
				dfs(edge, route)
			}
		}
		visited[i] = false
	}
	dfs(0, []int{})
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[3],[3],[]]\n
// @lcpr case=end

// @lcpr case=start
// [[4,3,1],[3,2,4],[3],[4],[]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  [][]int
	}{
		{
			graph: [][]int{{1, 2}, {3}, {3}, {}},
			want:  [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			want:  [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			graph: [][]int{{1}, {}},
			want:  [][]int{{0, 1}},
		},
		{
			graph: [][]int{{1, 2, 3}, {2}, {3}, {}},
			want:  [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := allPathsSourceTarget(tt.graph)
			if !equal(got, tt.want) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
