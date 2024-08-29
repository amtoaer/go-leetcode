package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3142 lang=golang
 *
 * [3142] Check if Grid Satisfies Conditions
 */

// @lc code=start
func satisfiesConditions(grid [][]int) bool {
	il := len(grid)
	jl := len(grid[0])
	for i := range grid {
		for j := range grid[0] {
			if i < il-1 && grid[i][j] != grid[i+1][j] {
				return false
			}
			if j < jl-1 && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		grid [][]int
		want bool
	}{
		{[][]int{{1, 0, 2}, {1, 0, 2}}, true},
	}
	for _, tt := range tc {
		assert.Equal(t, satisfiesConditions(tt.grid), tt.want)
	}
}
