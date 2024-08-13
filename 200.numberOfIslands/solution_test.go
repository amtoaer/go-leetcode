package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] Number of Islands
 */
// @lc code=start
var directions = [][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	cnt := 0
	var setVisited func(i, j int)
	setVisited = func(i, j int) {
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		for _, direction := range directions {
			ii, jj := i+direction[0], j+direction[1]
			if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) {
				setVisited(ii, jj)
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt += 1
				setVisited(i, j)
			}
		}
	}
	return cnt
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]byte
		output int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, numIslands(tt.input))
	}
}
