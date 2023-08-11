package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	var sum int
	for i := 0; i < len(mat); i++ {
		sum += (mat[i][i] + mat[i][len(mat)-i-1])
	}
	if len(mat)&1 == 1 {
		sum -= mat[len(mat)>>1][len(mat)>>1]
	}
	return sum
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 25, diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	assert.Equal(t, 8, diagonalSum([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}))
	assert.Equal(t, 5, diagonalSum([][]int{{5}}))
}
