package main

import (
	"math"
	"testing"
)

/*
 * @lc app=leetcode.cn id=3240 lang=golang
 * @lcpr version=20003
 *
 * [3240] 最少翻转次数使二进制矩阵回文 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m/2; i++ {
		for j := 0; j < n/2; j++ {
			var cnt [2]int
			cnt[grid[i][j]]++
			cnt[grid[m-i-1][j]]++
			cnt[grid[i][n-j-1]]++
			cnt[grid[m-i-1][n-j-1]]++
			maxCnt := max(cnt[0], cnt[1])
			res += 4 - maxCnt
		}
	}
	var rowNum1Cnt, ColNum1Cnt int
	if m&1 == 1 {
		for j := 0; j < n/2; j++ {
			if grid[m/2][j] == grid[m/2][n-j-1] {
				if rowNum1Cnt != math.MinInt {
					rowNum1Cnt = (rowNum1Cnt + grid[m/2][j]*2) % 4
				}
			} else {
				res++
				// 这里根据修改方式可以选择出现 0 个或者 2 个 1，因此无论如何都有策略保证 1 的数量可以 % 4
				rowNum1Cnt = math.MinInt
			}
		}
	}
	if n&1 == 1 {
		for i := 0; i < m/2; i++ {
			if grid[i][n/2] == grid[m-i-1][n/2] {
				if ColNum1Cnt != math.MinInt {
					ColNum1Cnt = (ColNum1Cnt + grid[i][n/2]*2) % 4
				}
			} else {
				res++
				ColNum1Cnt = math.MinInt
			}
		}
	}
	if ColNum1Cnt != math.MinInt && rowNum1Cnt != math.MinInt && (ColNum1Cnt+rowNum1Cnt)%4 == 2 {
		res += 2
	}
	if m&1 == 1 && n&1 == 1 && grid[m/2][n/2] == 1 {
		res += 1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,0,0],[0,1,0],[0,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,1],[0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1],[1]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, 3},
		{[][]int{{0, 1}, {0, 1}, {0, 0}}, 2},
		{[][]int{{1}, {1}}, 2},
		{[][]int{{0}}, 0},
	}

	for _, test := range tests {
		result := minFlips(test.grid)
		if result != test.expected {
			t.Errorf("For grid %v, expected %d but got %d", test.grid, test.expected, result)
		}
	}
}
