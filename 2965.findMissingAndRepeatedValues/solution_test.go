package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2965 lang=golang
 *
 * [2965] Find Missing and Repeated Values
 */

// @lc code=start
func findMissingAndRepeatedValues(grid [][]int) []int {
	res := make([]int, 2)
	n := len(grid)
	flag := make([]bool, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if flag[grid[i][j]-1] {
				res[0] = grid[i][j]
			}
			flag[grid[i][j]-1] = true
		}
	}
	for i, v := range flag {
		if !v {
			res[1] = i + 1
			return res
		}
	}
	return nil
}

// @lc code=end

func Test(t *testing.T) {
	var testcases = []struct {
		grid [][]int
		want []int
	}{
		{
			[][]int{{1, 3}, {2, 2}},
			[]int{2, 4},
		},
		{
			[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}},
			[]int{9, 5},
		},
	}
	for _, tc := range testcases {
		got := findMissingAndRepeatedValues(tc.grid)
		assert.Equal(t, tc.want, got)
	}
}
