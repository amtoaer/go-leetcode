package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3111 lang=golang
 *
 * [3111] Minimum Rectangles to Cover Points
 */

// @lc code=start
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cur, cnt := points[0][0], 1
	for _, p := range points {
		if cur+w >= p[0] {
			continue
		}
		cnt += 1
		cur = p[0]
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		points [][]int
		w      int
		want   int
	}{
		{[][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, minRectanglesToCoverPoints(tt.points, tt.w))
	}
}
