package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cur, cnt := points[0], 1
	for i := 1; i < len(points); i++ {
		if cur[1] < points[i][0] {
			cnt++
			cur = points[i]
		} else {
			cur[1] = min(cur[1], points[i][1])
		}
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, findMinArrowShots(tt.input))
	}
}
