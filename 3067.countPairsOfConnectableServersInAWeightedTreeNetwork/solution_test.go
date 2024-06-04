package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3067 lang=golang
 *
 * [3067] Count Pairs of Connectable Servers in a Weighted Tree Network
 */

// @lc code=start
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	graph := make([][][]int, n)
	for _, v := range edges {
		edgeL, edgeR, weight := v[0], v[1], v[2]
		graph[edgeL] = append(graph[edgeL], []int{edgeR, weight})
		graph[edgeR] = append(graph[edgeR], []int{edgeL, weight})
	}
	var dfs func(int, int, int) int
	dfs = func(start, end int, curr int) int {
		cnt := 0
		if curr == 0 {
			cnt += 1
		}
		for _, v := range graph[end] {
			next, weight := v[0], v[1]
			if next != start {
				cnt += dfs(end, next, (curr+weight)%signalSpeed)
			}
		}
		return cnt
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		pre := 0
		for _, v := range graph[i] {
			next, weight := v[0], v[1]
			cnt := dfs(i, next, weight%signalSpeed)
			res[i] += pre * cnt
			pre += cnt
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 4, 6, 6, 4, 0},
		countPairsOfConnectableServers(
			[][]int{{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2}},
			1,
		),
	)
	assert.Equal(
		t,
		[]int{2, 0, 0, 0, 0, 0, 2},
		countPairsOfConnectableServers(
			[][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}},
			3,
		),
	)
	assert.Equal(
		t,
		[]int{8, 28, 20, 0, 0, 0, 8, 0, 8, 0},
		countPairsOfConnectableServers(
			[][]int{
				{1, 0, 2},
				{2, 1, 4},
				{3, 2, 4},
				{4, 0, 3},
				{5, 1, 4},
				{6, 2, 2},
				{7, 6, 4},
				{8, 1, 2},
				{9, 8, 3},
			},
			1,
		),
	)
}
