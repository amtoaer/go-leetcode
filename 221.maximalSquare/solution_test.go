package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	var res int
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				res = max(res, dp[i][j]*dp[i][j])
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]byte
		output int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{[][]byte{{'0', '1'}, {'1', '0'}}, 1},
		{[][]byte{{'0'}}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, maximalSquare(tt.input))
	}
}
