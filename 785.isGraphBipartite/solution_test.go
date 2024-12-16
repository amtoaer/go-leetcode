package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=785 lang=golang
 * @lcpr version=20004
 *
 * [785] 判断二分图
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isBipartite(graph [][]int) bool {
	var (
		aSet, bSet = make(map[int]struct{}), make(map[int]struct{})
		vertify    func(node int, ASet bool) bool
	)
	vertify = func(node int, isASet bool) bool {
		validSet, invalidSet := aSet, bSet
		if !isASet {
			validSet, invalidSet = invalidSet, validSet
		}
		if _, ok := invalidSet[node]; ok {
			return false
		}
		if _, ok := validSet[node]; ok {
			return true
		}
		validSet[node] = struct{}{}
		for _, target := range graph[node] {
			if !vertify(target, !isASet) {
				return false
			}
		}
		return true
	}
	for node := range graph {
		_, ok1 := aSet[node]
		_, ok2 := bSet[node]
		if ok1 || ok2 {
			continue
		}
		if !vertify(node, true) {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[0,2],[0,1,3],[0,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[0,2],[1,3],[0,2]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  bool
	}{
		{
			graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			want:  false,
		},
		{
			graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			want:  true,
		},
		{
			graph: [][]int{{1}, {0, 3}, {3}, {1, 2}},
			want:  true,
		},
		{
			graph: [][]int{{1, 2}, {0, 2}, {0, 1}},
			want:  false,
		},
		{
			graph: [][]int{{1}, {0, 2}, {1, 3}, {2}},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isBipartite(tt.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
