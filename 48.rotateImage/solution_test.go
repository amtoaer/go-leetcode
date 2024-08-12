package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	left, right := 0, len(matrix[0])-1
	for left < right {
		for i := 0; i < len(matrix); i++ {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
		}
		left++
		right--
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			input:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			output: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tt := range tc {
		rotate(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
