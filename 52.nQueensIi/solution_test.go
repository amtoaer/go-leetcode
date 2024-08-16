package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	var (
		col            = make(map[int]struct{})
		diagonal_left  = make(map[int]struct{})
		diagonal_right = make(map[int]struct{})
		res            int
		backtrack      func(int)
	)
	backtrack = func(row int) {
		if row == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if _, ok := col[i]; ok {
				continue
			}
			if _, ok := diagonal_left[row+i]; ok {
				continue
			}
			if _, ok := diagonal_right[row-i]; ok {
				continue
			}
			col[i] = struct{}{}
			diagonal_left[row+i] = struct{}{}
			diagonal_right[row-i] = struct{}{}
			backtrack(row + 1)
			delete(col, i)
			delete(diagonal_left, row+i)
			delete(diagonal_right, row-i)
		}
	}
	backtrack(0)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output int
	}{
		{4, 2},
		{1, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, totalNQueens(tt.input))
	}
}
