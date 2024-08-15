package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3148 lang=golang
 *
 * [3148] Maximum Difference Score in a Grid
 */

// @lc code=start
func maxScore(grid [][]int) int {
	preRow, preCol, dp := make([][]int, len(grid)), make([][]int, len(grid)), make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		preRow[i] = make([]int, len(grid[0]))
		preCol[i] = make([]int, len(grid[0]))
		dp[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			dp[i][j] = math.MinInt
		}
	}
	res := math.MinInt
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 {
				dp[i][j] = max(dp[i][j], preRow[i-1][j]+grid[i][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], preCol[i][j-1]+grid[i][j])
			}
			res = max(res, dp[i][j])
			preRow[i][j] = max(dp[i][j], 0) - grid[i][j]
			preCol[i][j] = preRow[i][j]
			if i > 0 {
				preRow[i][j] = max(preRow[i][j], preRow[i-1][j])
			}
			if j > 0 {
				preCol[i][j] = max(preCol[i][j], preCol[i][j-1])
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}}, 9},
		{[][]int{{4, 3, 2}, {3, 2, 1}}, -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxScore(tt.grid))
	}
}
