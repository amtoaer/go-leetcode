package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=20004
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func numIslands(grid [][]byte) int {
	var dfs func(int, int)
	var res int
	dfs = func(i, j int) {
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = 2
		for _, direction := range directions {
			targetI, targetJ := i+direction[0], j+direction[1]
			if targetI >= 0 && targetI < len(grid) && targetJ >= 0 && targetJ < len(grid[0]) {
				dfs(targetI, targetJ)
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			grid: [][]byte{
				{'1', '0', '1', '0', '1'},
				{'0', '1', '0', '1', '0'},
				{'1', '0', '1', '0', '1'},
			},
			expected: 8,
		},
		{
			grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := numIslands(test.grid)
		if result != test.expected {
			t.Errorf("For grid %v, expected %d but got %d", test.grid, test.expected, result)
		}
	}
}
