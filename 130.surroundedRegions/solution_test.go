package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=20004
 *
 * [130] 被围绕的区域
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func solve(board [][]byte) {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'Z'
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}
	for i := 0; i < len(board); i++ {
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(0, j)
		dfs(len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Z' {
				board[i][j] = 'O'
			}
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected [][]byte
	}{
		{
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			board: [][]byte{
				{'X'},
			},
			expected: [][]byte{
				{'X'},
			},
		},
		{
			board: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
		},
		{
			board: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
		},
	}

	for _, test := range tests {
		solve(test.board)
		for i := range test.board {
			for j := range test.board[i] {
				if test.board[i][j] != test.expected[i][j] {
					t.Errorf("expected %v, but got %v", test.expected, test.board)
				}
			}
		}
	}
}
