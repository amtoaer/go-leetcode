package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1761 lang=golang
 *
 * [1761] Minimum Degree of a Connected Trio in a Graph
 */

// @lc code=start

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minTrioDegree(n int, edges [][]int) int {
	var (
		connected = make([][]bool, n)
		outEdges  = make([][]int, n)
		degrees   = make([]int, n)
	)
	for idx := range connected {
		connected[idx] = make([]bool, n)
	}
	for _, edge := range edges {
		i, j := edge[0]-1, edge[1]-1
		connected[i][j], connected[j][i] = true, true
		degrees[i]++
		degrees[j]++
	}
	for _, edge := range edges {
		i, j := edge[0]-1, edge[1]-1
		if degrees[i] < degrees[j] || (degrees[i] == degrees[j] && i < j) {
			outEdges[i] = append(outEdges[i], j)
		} else {
			outEdges[j] = append(outEdges[j], i)
		}
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		for _, j := range outEdges[i] {
			for _, k := range outEdges[j] {
				if connected[i][k] {
					res = min(res, degrees[i]+degrees[j]+degrees[k]-6)
				}
			}
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}))
	assert.Equal(
		t,
		0,
		minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}),
	)
}
