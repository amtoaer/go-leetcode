package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] Equal Row and Column Pairs
 */

// @lc code=start
func equalPairs(grid [][]int) int {
	m := make(map[string]int)
	for _, row := range grid {
		m[fmt.Sprint(row)]++
	}
	tmp := make([]int, len(grid))
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			tmp[j] = grid[j][i]
		}
		if val, ok := m[fmt.Sprint(tmp)]; ok {
			res += val
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
		{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, equalPairs(tt.grid))
	}
}
