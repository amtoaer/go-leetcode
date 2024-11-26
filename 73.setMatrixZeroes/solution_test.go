package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=73 lang=golang
 * @lcpr version=20003
 *
 * [73] 矩阵置零
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func setZeroes(matrix [][]int) {
	var firstRowEmpty, firstColEmpty bool
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColEmpty = true
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowEmpty = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if firstRowEmpty {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if firstColEmpty {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix:   [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix:   [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, test := range tests {
		setZeroes(test.matrix)
		for i := range test.matrix {
			for j := range test.matrix[i] {
				if test.matrix[i][j] != test.expected[i][j] {
					t.Errorf("expected %v, but got %v", test.expected, test.matrix)
				}
			}
		}
	}
}
