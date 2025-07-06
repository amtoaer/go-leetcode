/*
 * @lc app=leetcode.cn id=1252 lang=golang
 * @lcpr version=30201
 *
 * [1252] 奇数值单元格的数目
 */

package main

import (
	"testing"
)

// @lc code=start
func oddCells(m int, n int, indices [][]int) int {
	var (
		row = make([]int, m)
		col = make([]int, n)
	)
	for _, i := range indices {
		row[i[0]]++
		col[i[1]]++
	}
	var res int
	for _, r := range row {
		for _, c := range col {
			res += (r + c) % 2
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		m       int
		n       int
		indices [][]int
		want    int
	}{
		{2, 3, [][]int{{0, 1}, {1, 1}}, 6},
		{2, 2, [][]int{{1, 1}, {0, 0}}, 0},
	}
	for _, tt := range tc {
		got := oddCells(tt.m, tt.n, tt.indices)
		if got != tt.want {
			t.Errorf("oddCells(%v, %v, %v) = %v, want %v", tt.m, tt.n, tt.indices, got, tt.want)
		}
	}
}

/*
// @lcpr case=start
// 2\n3\n[[0,1],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n2\n[[1,1],[0,0]]\n
// @lcpr case=end

*/
