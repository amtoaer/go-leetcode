package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=20003
 *
 * [54] 螺旋矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var (
		left, top     int
		right, bottom = len(matrix[0]) - 1, len(matrix) - 1
		res           []int
	)
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if bottom != top && left != right {
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			for i := bottom - 1; i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 4, 3},
		},
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			want:   []int{1, 2, 3, 6, 5, 4},
		},
	}

	for _, tt := range tests {
		got := spiralOrder(tt.matrix)
		if !equal(got, tt.want) {
			t.Errorf("spiralOrder(%v) = %v; want %v", tt.matrix, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
