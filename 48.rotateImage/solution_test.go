package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=20003
 *
 * [48] 旋转图像
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rotate(matrix [][]int) {
	i, j := 0, len(matrix)-1
	for i < j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, test := range tests {
		rotate(test.matrix)
		for i := range test.matrix {
			for j := range test.matrix[i] {
				if test.matrix[i][j] != test.expected[i][j] {
					t.Errorf("expected %v, but got %v", test.expected, test.matrix)
				}
			}
		}
	}
}
