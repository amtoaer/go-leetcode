package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var res []int
	left, up, right, down := 0, 0, len(matrix[0])-1, len(matrix)-1
	for left <= right && up <= down {
		for j := left; j <= right; j++ {
			res = append(res, matrix[up][j])
		}
		for i := up + 1; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		for j := right - 1; j >= left && up != down; j-- {
			res = append(res, matrix[down][j])
		}
		for i := down - 1; i > up && left != right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		up++
		right--
		down--
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, spiralOrder(tt.input))
	}
}
