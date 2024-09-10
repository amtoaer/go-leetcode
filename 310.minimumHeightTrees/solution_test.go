package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	queue := []int{}
	degree := make([]int, n)
	adj := make([][]int, n)
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
		degree[v[0]]++
		degree[v[1]]++
	}
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	var res []int
	for len(queue) > 0 {
		l := len(queue)
		res = []int{}
		for i := 0; i < l; i++ {
			res = append(res, queue[i])
			for _, v := range adj[queue[i]] {
				degree[v]--
				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
		queue = queue[l:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{4, [][]int{{1, 0}, {1, 2}, {1, 3}}, []int{1}},
		{6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, []int{3, 4}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findMinHeightTrees(tt.n, tt.edges))
	}
}
