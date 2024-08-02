package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3128 lang=golang
 *
 * [3128] Right Triangles
 */

// @lc code=start
func numberOfRightTriangles(grid [][]int) int64 {
	cols := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(cols); j++ {
			cols[j] += grid[i][j]
		}
	}
	var res int64
	for i := 0; i < len(grid); i++ {
		row := 0
		for j := 0; j < len(cols); j++ {
			row += grid[i][j]
		}
		for j := 0; j < len(cols); j++ {
			if grid[i][j] == 1 {
				res += int64(row-1) * int64(cols[j]-1)
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		grid [][]int
		want int64
	}{
		{[][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}, 2},
		{[][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}}, 2},
		{[][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, numberOfRightTriangles(tt.grid))
	}
}
